package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"
	"strings"
	"time"

	"github.com/google/uuid"
)

// UploadService menangani penyimpanan file yang diunggah
type UploadServiceServer struct {
	pb.UnimplementedUploadDataSekolahServiceServer
	uploadDir          string
	repoSiswa          repositories.GenericRepository[models.PesertaDidik]
	repoSiswaPelengkap repositories.GenericRepository[models.PesertaDidikPelengkap]
	repoKelas          repositories.GenericRepository[models.RombonganBelajar]
	repoKelasAnggota   repositories.GenericRepository[models.RombelAnggota]
	repoGuru           repositories.GenericRepository[models.TabelPTK]
	repoGuruTerdaftar  repositories.GenericRepository[models.PTKTerdaftar]
	repoGuruPelengkap  repositories.GenericRepository[models.PtkPelengkap]
	// uploadParameter    ParamTemplate
}

func NewUploadServiceServer() *UploadServiceServer {
	repoSiswa := repositories.NewSiswaRepository(config.DB)
	repoSiswaPelengkap := repositories.NewSiswaPelengkapRepository(config.DB)
	repoKelas := repositories.NewrombonganBelajarRepository(config.DB)
	repoKelasAnggota := repositories.NewRombelAnggotaRepository(config.DB)
	repoGuru := repositories.NewPTKRepository(config.DB)
	repoGuruTerdaftar := repositories.NewPTKTerdaftarRepository(config.DB)
	repoGuruPelengkap := repositories.NewPTKPelengkapRepository(config.DB)
	// if repoSiswa == nil {
	// 	log.Fatal("❌ ERROR: Gagal menginisialisasi repoSiswa") // Debugging
	// }
	return &UploadServiceServer{
		uploadDir:          "uploads",
		repoSiswa:          *repoSiswa,
		repoSiswaPelengkap: *repoSiswaPelengkap,
		repoKelas:          *repoKelas,
		repoKelasAnggota:   *repoKelasAnggota,
		repoGuru:           *repoGuru,
		repoGuruTerdaftar:  *repoGuruTerdaftar,
		repoGuruPelengkap:  *repoGuruPelengkap,
	}
}

// UploadFile menangani upload melalui gRPC
func (s *UploadServiceServer) UploadDataSekolah(ctx context.Context, req *pb.UploadDataSekolahRequest) (*pb.UploadDataSekolahResponse, error) {
	// Tentukan lokasi penyimpanan file
	filePath := filepath.Join(s.uploadDir, req.Filename)

	// Simpan file yang dikirim dalam bytes
	err := os.WriteFile(filePath, req.File, 0644)
	if err != nil {
		return nil, fmt.Errorf("gagal menyimpan file: %w", err)
	}

	// Kembalikan URL file yang diunggah
	return &pb.UploadDataSekolahResponse{
		Message: "File berhasil diunggah",
		FileUrl: "/storage/uploads/" + req.Filename,
	}, nil
}

// UploadFileHTTP menangani upload file melalui REST API dengan multipart/form-data
func (s *UploadServiceServer) UploadFileHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Gagal mem-parsing form data", http.StatusBadRequest)
		return
	}

	// Ambil upload_type dari form data, bukan query
	paramValue := r.FormValue("upload_type")
	if paramValue == "" {
		http.Error(w, "upload_type tidak boleh kosong", http.StatusBadRequest)
		return
	}

	fileHeader := r.MultipartForm.File["file"]
	if len(fileHeader) == 0 {
		http.Error(w, "File tidak ditemukan", http.StatusBadRequest)
		return
	}

	file, err := fileHeader[0].Open()
	if err != nil {
		http.Error(w, "Gagal membuka file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	param := ParamTemplate{
		templateType: paramValue,
		schemaname:   r.FormValue("schemaname"),
		// semesterId:   r.FormValue("semester_id"),
	}

	// s.uploadParameter
	if len(fileHeader) == 0 {
		http.Error(w, "File tidak ditemukan", http.StatusBadRequest)
		return
	}

	// Validasi file Excel
	fileName := fileHeader[0].Filename
	if !strings.HasSuffix(fileName, ".xlsx") {
		http.Error(w, "Hanya file Excel (.xlsx) yang diizinkan", http.StatusBadRequest)
		return
	}

	// Simpan file sementara
	tempFile, err := os.CreateTemp("", "upload-*.xlsx")
	if err != nil {
		http.Error(w, "Gagal membuat file sementara", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Gagal menyimpan file sementara", http.StatusInternalServerError)
		return
	}
	param.filePath = tempFile.Name()
	// Gunakan map untuk memilih fungsi yang sesuai berdasarkan uploadType
	uploadHandlers := map[string]func() error{
		"siswa": func() error {
			return s.processUploadSiswa(
				context.Background(), param,
				BacaDataExcel,
			)
		},
		"guru": func() error {
			return s.processUploadGuru(
				context.Background(), param,
				BacaDataExcel,
			)
		},
		// Tambahkan tipe lain di sini...
	}

	// Periksa apakah uploadType valid
	if handler, exists := uploadHandlers[param.templateType]; exists {
		if err := handler(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Upload sukses!")
	} else {
		http.Error(w, "Tipe upload tidak dikenali", http.StatusBadRequest)
	}
	// var data []interface{}
	// if uploadType == "siswa" {
	// 	data, err = UploadDataSekolah[models.PesertaDidik](filePath, uploadType)
	// 	if err != nil {
	// 		http.Error(w, fmt.Sprintf("Gagal memproses file: %v", err), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	defer cancel()
	// 	for i := range data {
	// 		utils.HandleNilPointers(&data[i])
	// 		// pesertaDidikPointers = append(pesertaDidikPointers, &data[i])
	// 		if s == nil {
	// 			log.Fatal("❌ ERROR: s is nil! Periksa inisialisasi UploadServiceServer")
	// 		}
	// 		err = s.repoSiswa.Save(ctx, &data[i], schemaname)
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 		}
	// 		fmt.Println("sukses")
	// 	}
	// }

	// Berikan respon
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "File berhasil diproses",
		// "data":    data,
	})
}

// GetTemplate menyediakan template Excel berdasarkan jenis data
func (s *UploadServiceServer) DownloadDataSekolah(ctx context.Context, req *pb.DownloadDataSekolahRequest) (*pb.DownloadDataSekolahResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"TemplateType"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	templateType := req.GetDownloadType()
	templatePath := fmt.Sprintf("templates/template_%s.xlsx", templateType)
	var param = ParamTemplate{
		schemaname:   "",
		filePath:     "",
		semesterId:   "",
		templateType: templateType,
	}
	// Buat file template jika belum ada
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		err := GenerateTemplate(param, config.DB)
		if err != nil {
			return nil, fmt.Errorf("gagal membuat template %s: %w", templateType, err)
		}
	}

	// Baca file template
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca template %s: %w", templateType, err)
	}

	return &pb.DownloadDataSekolahResponse{
		Filename: fmt.Sprintf("template_%s.xlsx", templateType),
		File:     data,
	}, nil
}

// HandleDownloadTemplate adalah handler untuk mengunduh file template .xlsx.
func (h *UploadServiceServer) DownloadTemplateHTTP(w http.ResponseWriter, r *http.Request) {
	// Ambil nama file dari query parameter
	templateType := r.URL.Query().Get("template_type")
	if templateType == "" {
		http.Error(w, "template-type is required", http.StatusBadRequest)
		return
	}
	// Lokasi direktori template
	var param = ParamTemplate{
		schemaname:   r.FormValue("Schemaname"),
		semesterId:   r.FormValue("semesterId"),
		templateType: templateType,
	}
	templatePath := fmt.Sprintf("templates/template_%s_%s_%s.xlsx", templateType, param.schemaname, param.semesterId)
	param.filePath = templatePath
	// Buat file template jika belum ada
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		err := GenerateTemplate(param, config.DB)
		if err != nil {
			http.Error(w, fmt.Sprintf("Gagal membuat template: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// Buka file template
	file, err := os.Open(templatePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal membuka file template: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Mendapatkan informasi file untuk header
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal mendapatkan informasi file: %v", err), http.StatusInternalServerError)
		return
	}

	// w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"; filename*=UTF-8''%s", fileInfo.Name(), url.QueryEscape(fileInfo.Name())))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name()))

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet") // MIME type untuk Excel
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")

	// Kirim file
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, fmt.Sprintf("Gagal mengirim file: %v", err), http.StatusInternalServerError)
		return
	}

}

// Definisikan fungsi untuk menangani upload berdasarkan tipe data
// func processUpload[T any](
// 	ctx context.Context,
// 	param ParamTemplate,
// 	uploadFunc func(ParamTemplate) ([]T, error),
// 	saveFunc func(context.Context, *T, string) error,
// ) error {
// 	// Ambil data dari file
// 	data, err := uploadFunc(param)
// 	if err != nil {
// 		return fmt.Errorf("gagal memproses file: %v", err)
// 	}

// 	// Iterasi data dan simpan ke database
// 	for i := range data {
// 		utils.HandleNilPointers(&data[i]) // Hindari pointer nil

//			// Simpan ke database
//			if err := saveFunc(ctx, &data[i], param.schemaname); err != nil {
//				return fmt.Errorf("gagal menyimpan data: %v", err)
//			}
//		}
//		return nil
//	}
func (s *UploadServiceServer) processUploadSiswa(
	ctx context.Context,
	param ParamTemplate,
	uploadFunc func(*ParamTemplate) ([][]string, error),
) error {
	// Ambil data dari file
	data, err := uploadFunc(&param)
	if err != nil {
		return fmt.Errorf("gagal memproses file: %v", err)
	}

	for i, row := range data {
		// Skip header atau baris kosong (misalnya baris pertama sebagai judul)
		if i == 0 || len(row) < 16 {
			continue
		}

		// Buat UUID untuk entitas terkait
		pesertaDidikId := uuid.New()
		pelengkapSiswaId := uuid.New()
		anggotaRombelId := uuid.New()

		// Parse tanggal lahir jika tersedia
		var tanggalLahir *time.Time
		tanggalLahirStr := row[4] // tanggal_lahir di index 4
		if tanggalLahirStr != "" {
			parsedDate, err := time.Parse("2006-01-02", tanggalLahirStr)
			if err == nil {
				tanggalLahir = &parsedDate
			}
		}

		// Simpan ke tabel peserta didik (siswa)
		err := s.repoSiswa.Save(ctx, &models.PesertaDidik{
			PesertaDidikId:  pesertaDidikId.String(),
			NmSiswa:         row[2],       // nm_siswa
			Nis:             row[0],       // nis
			Nisn:            row[1],       // nisn
			JenisKelamin:    row[5],       // jenis_kelamin
			TempatLahir:     row[3],       // tempat_lahir
			TanggalLahir:    tanggalLahir, // tanggal_lahir
			Agama:           row[6],       // agama
			AlamatSiswa:     &row[7],      // alamat_siswa
			TeleponSiswa:    row[8],       // telepon_siswa
			DiterimaTanggal: nil,          // tidak ada di contoh data
			NmAyah:          row[10],      // nm_ayah
			NmIbu:           row[11],      // nm_ibu
			PekerjaanAyah:   row[12],      // pekerjaan_ayah
			PekerjaanIbu:    row[13],      // pekerjaan_ibu
			NmWali:          &row[14],     // nm_wali
			PekerjaanWali:   &row[15],     // pekerjaan_wali
		}, param.schemaname)

		if err != nil {
			return fmt.Errorf("gagal menyimpan peserta didik: %v", err)
		}

		// Simpan ke tabel pelengkap siswa (opsional)
		err = s.repoSiswaPelengkap.Save(ctx, &models.PesertaDidikPelengkap{
			PelengkapSiswaId: pelengkapSiswaId.String(),
			PesertaDidikId:   func(s string) *string { return &s }(pesertaDidikId.String()),
			SekolahAsal:      "",
			AnakKe:           nil,
		}, param.schemaname)

		if err != nil {
			return fmt.Errorf("gagal menyimpan data pelengkap siswa: %v", err)
		}

		// Jika kelas tersedia, simpan ke rombel anggota
		namaKelas := row[15] // Misalkan pekerjaan_wali digunakan untuk kelas? Cek lagi formatmu
		if namaKelas != "" {
			err = s.repoKelasAnggota.Save(ctx, &models.RombelAnggota{
				AnggotaRombelId: utils.StringToUUID(anggotaRombelId.String()),
				PesertaDidikId:  utils.StringToUUID(pesertaDidikId.String()),
				RombonganBelajar: models.RombonganBelajar{
					NmKelas: namaKelas,
				},
			}, param.schemaname)

			if err != nil {
				return fmt.Errorf("gagal menyimpan anggota rombel: %v", err)
			}
		}
	}

	return nil
}

func (s *UploadServiceServer) processUploadGuru(
	ctx context.Context,
	param ParamTemplate,
	uploadFunc func(*ParamTemplate) ([][]string, error),
) error {
	// Ambil data dari file
	data, err := uploadFunc(&param)
	if err != nil {
		return fmt.Errorf("gagal memproses file: %v", err)
	}

	// Iterasi data dan simpan ke database
	for i := range data {
		ptkId := uuid.New()
		ptkTerdaftarId := uuid.New()

		var tanggalLahir time.Time
		tanggalLahirStr := data[i][4]
		if tanggalLahirStr != "" {
			tanggalLahir, err = time.Parse("2006-01-02", tanggalLahirStr)
			if err == nil {
				// tanggalLahir = &tanggalLahir
			}
		}
		err := s.repoGuru.Save(ctx, &models.TabelPTK{
			PtkID:        ptkId,
			Nama:         data[i][1],
			JenisKelamin: &data[i][2],
			TempatLahir:  &data[i][3],
			TanggalLahir: &tanggalLahir,
			Agama:        &data[i][5],
		}, param.schemaname)
		if err != nil {
			return err
		}
		err = s.repoGuruTerdaftar.Save(ctx, &models.PTKTerdaftar{
			PtkTerdaftarId: ptkTerdaftarId,
			PtkID:          ptkId,
			TahunAjaranId:  param.semesterId[:4],
			// FotoSiswa: ,
		}, param.schemaname)

		if err != nil {
			return err
		}
		err = s.repoGuruPelengkap.Save(ctx, &models.PtkPelengkap{
			PtkId:      ptkId,
			GelarDepan: &data[i][6],
			// GelarBelakang: ,
		}, param.schemaname)
		if err != nil {
			return err
		}
	}
	return nil
}
