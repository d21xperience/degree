package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// UploadService menangani penyimpanan file yang diunggah
type UploadServiceServer struct {
	pb.UnimplementedUploadDataSekolahServiceServer
	uploadDir              string
	repoSiswa              repositories.GenericRepository[models.PesertaDidik]
	repoSiswaPelengkap     repositories.GenericRepository[models.PesertaDidikPelengkap]
	repoKelas              repositories.GenericRepository[models.RombonganBelajar]
	repoKelasAnggota       repositories.GenericRepository[models.RombelAnggota]
	repoGuru               repositories.GenericRepository[models.TabelPTK]
	repoGuruTerdaftar      repositories.GenericRepository[models.PTKTerdaftar]
	repoGuruPelengkap      repositories.GenericRepository[models.PtkPelengkap]
	repoKategoriSekolah    repositories.GenericRepository[models.KategoriSekolah]
	repoKategoriSekolahLog repositories.GenericRepository[models.KategoriSekolahLog]
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
	repoKategoriSekolah := repositories.NewKategoriSekolahRepository(config.DB)
	repoKategoriSekolahLog := repositories.NewKategoriSekolahLogRepository(config.DB)

	return &UploadServiceServer{
		uploadDir:              "uploads",
		repoSiswa:              *repoSiswa,
		repoSiswaPelengkap:     *repoSiswaPelengkap,
		repoKelas:              *repoKelas,
		repoKelasAnggota:       *repoKelasAnggota,
		repoGuru:               *repoGuru,
		repoGuruTerdaftar:      *repoGuruTerdaftar,
		repoGuruPelengkap:      *repoGuruPelengkap,
		repoKategoriSekolah:    *repoKategoriSekolah,
		repoKategoriSekolahLog: *repoKategoriSekolahLog,
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
		"kelas": func() error {
			return s.processUploadKelas(
				context.Background(), param,
				BacaDataExcel,
			)
		},
		"transkrip": func() error {
			return s.processUploadTranskrip(
				context.Background(), param,
				BacaDataExcel,
			)
		},
		"nilai": func() error {
			return s.processUploadTranskrip(
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
	// Berikan respon
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "File berhasil diproses",
		// "data":    data,
	})
}

// ===============versi 3
func (h *UploadServiceServer) DownloadTemplateHTTP(w http.ResponseWriter, r *http.Request) {

	// Cek method
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON body
	var req struct {
		RombonganBelajarId  string `json:"rombelId"`
		KurikulumId         string `json:"kurikulumId"`
		TahunAjaranId       string `json:"tahunAjaranId"`
		TingkatPendidikanId string `json:"tingkatPendidikanId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	templateType := r.URL.Query().Get("template_type")
	if templateType == "" || r.URL.Query().Get("schemaname") == "" {
		http.Error(w, "template-type or schemaname is required", http.StatusBadRequest)
		return
	}

	// if templateType == "nilai" {
	// 	if r.URL.Query().Get("kurikulum_id") == "" {
	// 		http.Error(w, "kurikulum_id is required", http.StatusBadRequest)
	// 		return
	// 	}
	// }

	param := ParamTemplate{
		schemaname:          r.FormValue("schemaname"),
		semesterId:          r.FormValue("semesterId"),
		templateType:        templateType,
		RombonganBelajarId:  req.RombonganBelajarId,
		KurikulumId:         req.KurikulumId,
		tahunAjaranId:       req.TahunAjaranId,
		TingkatPendidikanId: req.TingkatPendidikanId,
	}

	// Generate template langsung ke memori (tidak simpan ke file)
	f, err := GenerateTemplateV2(param, config.DB)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal membuat template: %v", err), http.StatusInternalServerError)
		return
	}

	// Tulis file ke buffer memory
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		http.Error(w, fmt.Sprintf("Gagal menulis template ke buffer: %v", err), http.StatusInternalServerError)
		return
	}

	// Header respons
	filename := fmt.Sprintf("template_%s_%s_%s.xlsx", templateType, param.schemaname, param.semesterId)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", buf.Len()))
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")

	// Kirim file langsung dari buffer
	if _, err := io.Copy(w, &buf); err != nil {
		http.Error(w, fmt.Sprintf("Gagal mengirim file: %v", err), http.StatusInternalServerError)
		return
	}
}

func (s *UploadServiceServer) processUploadSiswa(
	ctx context.Context,
	param ParamTemplate,
	uploadFunc func(*ParamTemplate) ([][]string, error),
) error {
	var err error
	// Ambil data dari file
	data, err := uploadFunc(&param)
	if err != nil {
		return fmt.Errorf("gagal memproses file: %v", err)
	}
	// log.Printf("pembacaan di luar loop %s", data[0][4])
	for i, row := range data {
		pesertaDidikId := uuid.New()

		if err != nil {
			return err
		}

		var tglLahir, tglDiterima *time.Time
		if len(row[7]) != 0 {
			cek, err := utils.StringToTime(data[i][7], "02/01/2006")
			if err != nil {
				log.Printf("gagal parsing tanggal lahir %s", row[7])
			}
			tglLahir = utils.TimeToPointer(cek.Format("2006-01-02"))
		}
		if len(row[11]) != 0 {
			cek, err := utils.StringToTime(data[i][11], "02/01/2006")
			if err != nil {
				log.Printf("gagal parsing tanggal mendaftar %s", row[11])
				return nil
			}
			tglDiterima = utils.TimeToPointer(cek.Format("2006-01-02"))
		}

		// Simpan ke tabel peserta didik (siswa)
		err = s.repoSiswa.Save(ctx, &models.PesertaDidik{
			PesertaDidikId:  pesertaDidikId,
			Nis:             safeGet(row, 2),
			Nisn:            safeGet(row, 3),
			NmSiswa:         row[4],
			JenisKelamin:    safeGet(row, 5),
			TempatLahir:     safeGet(row, 6),
			TanggalLahir:    tglLahir,
			Agama:           safeGet(row, 8),
			AlamatSiswa:     safeGet(row, 9),
			TeleponSiswa:    safeGet(row, 10),
			DiterimaTanggal: tglDiterima,
			NmAyah:          safeGet(row, 12),
			NmIbu:           safeGet(row, 13),
			PekerjaanAyah:   safeGet(row, 14),
			PekerjaanIbu:    safeGet(row, 15),
			NmWali:          safeGet(row, 16),
			PekerjaanWali:   safeGet(row, 17),
			Nik:             safeGet(row, 18),
		}, param.schemaname)

		if err != nil {
			// Jika duplikat NISN, ambil pesertaDidikId dari database
			if strings.Contains(err.Error(), "failed to save record") {
				log.Printf("Duplikat NISN ditemukan baris %d: %s — mengambil ID dari DB", i+1, row[3])

				// Ambil ID dari database berdasarkan NISN
				pd, getErr := s.repoSiswa.FindByID(ctx, row[3], param.schemaname, "nisn")
				if getErr != nil {
					return fmt.Errorf("duplikat NISN, tetapi gagal ambil ID: %v", getErr)
				}
				pesertaDidikId = pd.PesertaDidikId
			} else {
				return fmt.Errorf("gagal menyimpan peserta didik: %v", err)
			}
		} else {
			// Simpan ke tabel pelengkap siswa (opsional)
			err = s.repoSiswaPelengkap.Save(ctx, &models.PesertaDidikPelengkap{
				PelengkapSiswaId: uuid.New(),
				PesertaDidikId:   pesertaDidikId,
				StatusDalamKel:   safeGet(row, 19),
				AnakKe:           safeGet(row, 20),
				SekolahAsal:      safeGet(row, 21),
				DiterimaKelas:    safeGet(row, 22),
				AlamatOrtu:       safeGet(row, 23),
				TeleponOrtu:      safeGet(row, 24),
				AlamatWali:       safeGet(row, 25),
				TeleponWali:      safeGet(row, 26),
			}, param.schemaname)

			if err != nil {
				// Jika duplikat NISN, ambil pesertaDidikId dari database
				if strings.Contains(err.Error(), "failed to save record") {
					log.Printf("Duplikat peserta_didik_id ditemukan baris %d: %s — mengambil ID dari DB", i+1, row[3])
				} else {
					return fmt.Errorf("gagal menyimpan peserta didik: %v", err)
				}
			}
		}

		// Jika kelas tersedia, simpan ke rombel anggota
		for j := 1; j <= 2; j++ {
			// rombonganBelajarId := utils.StringToUUID(row[27])
			conditions := map[string]any{
				"tabel_kelas.nm_kelas":    row[1],
				"tabel_kelas.semester_id": fmt.Sprintf("%s%d", param.semesterId, j),
			}
			rombonganBelajarId, err1 := s.repoKelas.FindWithPreloadAndJoinsOrigin(ctx, param.schemaname, nil, nil, conditions, nil) //utils.StringToUUID(row[27])
			if err1 != nil {
				return err1
			}
			err = s.repoKelasAnggota.Save(ctx, &models.RombelAnggota{
				AnggotaRombelId:    uuid.New(),
				PesertaDidikId:     pesertaDidikId,
				SemesterId:         fmt.Sprintf("%s%d", param.semesterId, j),
				RombonganBelajarId: utils.UUIDToPointer(rombonganBelajarId[0].RombonganBelajarId),
			}, param.schemaname)

			if err != nil {
				return fmt.Errorf("gagal menyimpan anggota rombel: %v", err)
			}

		}
	}

	return nil
}

// func (h *UploadServiceServer) DownloadTemplateHTTP(w http.ResponseWriter, r *http.Request) {
// 	templateType := r.URL.Query().Get("template_type")
// 	if templateType == "" || r.URL.Query().Get("schemaname") == "" {
// 		http.Error(w, "template-type or schemaname is required", http.StatusBadRequest)
// 		return
// 	}

// 	if templateType == "nilai" {
// 		if r.URL.Query().Get("kurikulum_id") == "" {
// 			http.Error(w, "kurikulum_id is required", http.StatusBadRequest)
// 			return
// 		}
// 	}
// 	param := ParamTemplate{
// 		schemaname:   r.FormValue("schemaname"),
// 		semesterId:   r.FormValue("semesterId"),
// 		KurikulumId:  r.FormValue("kurikulum_id"),
// 		templateType: templateType,
// 	}

// 	// Generate template langsung ke memori (tidak simpan ke file)
// 	f, err := GenerateTemplateV2(param, config.DB)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Gagal membuat template: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// Tulis file ke buffer memory
// 	var buf bytes.Buffer
// 	if err := f.Write(&buf); err != nil {
// 		http.Error(w, fmt.Sprintf("Gagal menulis template ke buffer: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// Header respons
// 	filename := fmt.Sprintf("template_%s_%s_%s.xlsx", templateType, param.schemaname, param.semesterId)
// 	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
// 	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
// 	w.Header().Set("Content-Length", fmt.Sprintf("%d", buf.Len()))
// 	w.Header().Set("Cache-Control", "no-cache")
// 	w.Header().Set("Pragma", "no-cache")

// 	// Kirim file langsung dari buffer
// 	if _, err := io.Copy(w, &buf); err != nil {
// 		http.Error(w, fmt.Sprintf("Gagal mengirim file: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// }

// func (s *UploadServiceServer) processUploadSiswa(
// 	ctx context.Context,
// 	param ParamTemplate,
// 	uploadFunc func(*ParamTemplate) ([][]string, error),
// ) error {
// 	var err error
// 	// Ambil data dari file
// 	data, err := uploadFunc(&param)
// 	if err != nil {
// 		return fmt.Errorf("gagal memproses file: %v", err)
// 	}
// 	// log.Printf("pembacaan di luar loop %s", data[0][4])
// 	for i, row := range data {
// 		pesertaDidikId := uuid.New()

// 		if err != nil {
// 			return err
// 		}

// 		var tglLahir, tglDiterima *time.Time
// 		if len(row[7]) != 0 {
// 			cek, err := utils.StringToTime(data[i][7], "02/01/2006")
// 			if err != nil {
// 				log.Printf("gagal parsing tanggal lahir %s", row[7])
// 			}
// 			tglLahir = utils.TimeToPointer(cek.Format("2006-01-02"))
// 		}
// 		if len(row[11]) != 0 {
// 			cek, err := utils.StringToTime(data[i][11], "02/01/2006")
// 			if err != nil {
// 				log.Printf("gagal parsing tanggal mendaftar %s", row[11])
// 				return nil
// 			}
// 			tglDiterima = utils.TimeToPointer(cek.Format("2006-01-02"))
// 		}

// 		// Simpan ke tabel peserta didik (siswa)
// 		err = s.repoSiswa.Save(ctx, &models.PesertaDidik{
// 			PesertaDidikId:  pesertaDidikId,
// 			Nis:             safeGet(row, 2),
// 			Nisn:            safeGet(row, 3),
// 			NmSiswa:         row[4],
// 			JenisKelamin:    safeGet(row, 5),
// 			TempatLahir:     safeGet(row, 6),
// 			TanggalLahir:    tglLahir,
// 			Agama:           safeGet(row, 8),
// 			AlamatSiswa:     safeGet(row, 9),
// 			TeleponSiswa:    safeGet(row, 10),
// 			DiterimaTanggal: tglDiterima,
// 			NmAyah:          safeGet(row, 12),
// 			NmIbu:           safeGet(row, 13),
// 			PekerjaanAyah:   safeGet(row, 14),
// 			PekerjaanIbu:    safeGet(row, 15),
// 			NmWali:          safeGet(row, 16),
// 			PekerjaanWali:   safeGet(row, 17),
// 			Nik:             safeGet(row, 18),
// 		}, param.schemaname)

// 		if err != nil {
// 			// Jika duplikat NISN, ambil pesertaDidikId dari database
// 			if strings.Contains(err.Error(), "failed to save record") {
// 				log.Printf("Duplikat NISN ditemukan baris %d: %s — mengambil ID dari DB", i+1, row[3])

// 				// Ambil ID dari database berdasarkan NISN
// 				pd, getErr := s.repoSiswa.FindByID(ctx, row[3], param.schemaname, "nisn")
// 				if getErr != nil {
// 					return fmt.Errorf("duplikat NISN, tetapi gagal ambil ID: %v", getErr)
// 				}
// 				pesertaDidikId = pd.PesertaDidikId
// 			} else {
// 				return fmt.Errorf("gagal menyimpan peserta didik: %v", err)
// 			}
// 		} else {
// 			// Simpan ke tabel pelengkap siswa (opsional)
// 			err = s.repoSiswaPelengkap.Save(ctx, &models.PesertaDidikPelengkap{
// 				PelengkapSiswaId: uuid.New(),
// 				PesertaDidikId:   pesertaDidikId,
// 				StatusDalamKel:   safeGet(row, 19),
// 				AnakKe:           safeGet(row, 20),
// 				SekolahAsal:      safeGet(row, 21),
// 				DiterimaKelas:    safeGet(row, 22),
// 				AlamatOrtu:       safeGet(row, 23),
// 				TeleponOrtu:      safeGet(row, 24),
// 				AlamatWali:       safeGet(row, 25),
// 				TeleponWali:      safeGet(row, 26),
// 			}, param.schemaname)

// 			if err != nil {
// 				// Jika duplikat NISN, ambil pesertaDidikId dari database
// 				if strings.Contains(err.Error(), "failed to save record") {
// 					log.Printf("Duplikat peserta_didik_id ditemukan baris %d: %s — mengambil ID dari DB", i+1, row[3])
// 				} else {
// 					return fmt.Errorf("gagal menyimpan peserta didik: %v", err)
// 				}
// 			}
// 		}

// 		// Jika kelas tersedia, simpan ke rombel anggota
// 		for j := 1; j <= 2; j++ {
// 			// rombonganBelajarId := utils.StringToUUID(row[27])
// 			conditions := map[string]any{
// 				"tabel_kelas.nm_kelas":    row[1],
// 				"tabel_kelas.semester_id": fmt.Sprintf("%s%d", param.semesterId, j),
// 			}
// 			rombonganBelajarId, err1 := s.repoKelas.FindWithPreloadAndJoinsOrigin(ctx, param.schemaname, nil, nil, conditions, nil) //utils.StringToUUID(row[27])
// 			if err1 != nil {
// 				return err1
// 			}
// 			err = s.repoKelasAnggota.Save(ctx, &models.RombelAnggota{
// 				AnggotaRombelId:    uuid.New(),
// 				PesertaDidikId:     pesertaDidikId,
// 				SemesterId:         fmt.Sprintf("%s%d", param.semesterId, j),
// 				RombonganBelajarId: utils.UUIDToPointer(rombonganBelajarId[0].RombonganBelajarId),
// 			}, param.schemaname)

// 			if err != nil {
// 				return fmt.Errorf("gagal menyimpan anggota rombel: %v", err)
// 			}

// 		}
// 	}

// 	return nil
// }

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
		tanggalLahirStr := data[i][3]
		if tanggalLahirStr != "" {
			tanggalLahir, err = time.Parse("01-02-06", tanggalLahirStr)
			if err == nil {
				// tanggalLahir = &tanggalLahir
			}
		}
		err := s.repoGuru.Save(ctx, &models.TabelPTK{
			PtkID:        ptkId,
			Nama:         data[i][0],
			JenisKelamin: &data[i][1],
			TempatLahir:  &data[i][2],
			TanggalLahir: &tanggalLahir,
			Agama:        &data[i][4],
			JenisPtkID:   4,
		}, param.schemaname)
		if err != nil {
			return err
		}
		err = s.repoGuruTerdaftar.Save(ctx, &models.PTKTerdaftar{
			PtkTerdaftarId: ptkTerdaftarId,
			PtkID:          ptkId,
			TahunAjaranId:  param.semesterId[:4],
		}, param.schemaname)

		if err != nil {
			return err
		}
		err = s.repoGuruPelengkap.Save(ctx, &models.PtkPelengkap{
			PtkPelengkapId: uuid.New(),
			PtkId:          ptkId,
			GelarDepan:     &data[i][5],
			GelarBelakang:  &data[i][6],
			Nik:            &data[i][7],
			NoKk:           &data[i][8],
			Nuptk:          &data[i][9],
			Niy:            &data[i][10],
			Nip:            &data[i][11],
			AlamatJalan:    &data[i][12],
			Rt:             &data[i][13],
			Rw:             &data[i][14],
			DesaKelurahan:  &data[i][15],
			Kec:            &data[i][16],
			KabKota:        &data[i][17],
			Propinsi:       &data[i][18],
			KodePos:        &data[i][19],
			NoTeleponRumah: &data[i][20],
			NoHp:           &data[i][21],
			Email:          &data[i][22],
			// GelarBelakang: ,
		}, param.schemaname)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *UploadServiceServer) processUploadKelas(
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
	for _, row := range data {
		tingkat, err1 := strconv.Atoi(row[1])
		if err1 != nil {
			return err
		}
		// Jika kelas tersedia, simpan ke rombel anggota
		for j := 1; j <= 2; j++ {
			err := s.repoKelas.Save(ctx, &models.RombonganBelajar{
				RombonganBelajarId:  uuid.New(),
				SemesterId:          fmt.Sprintf("%s%d", param.semesterId, j),
				NmKelas:             row[0],
				TingkatPendidikanId: int32(tingkat),
				JenisRombel:         utils.Int32ToPointer(1),
			}, param.schemaname)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (s *UploadServiceServer) processUploadTranskrip(
	ctx context.Context,
	param ParamTemplate,
	uploadFunc func(*ParamTemplate) ([][]string, error),
) error {
	// Ambil data dari file
	// data, err := uploadFunc(&param)
	// if err != nil {
	// 	return fmt.Errorf("gagal memproses file: %v", err)
	// }

	// Iterasi data dan simpan ke database
	// for _, row := range data {
	// tingkat, err1 := strconv.Atoi(row[1])
	// if err1 != nil {
	// 	return err
	// }
	// // Jika kelas tersedia, simpan ke rombel anggota
	// for j := 1; j <= 2; j++ {
	// 	err := s.repoKelas.Save(ctx, &models.RombonganBelajar{
	// 		RombonganBelajarId:  uuid.New(),
	// 		SemesterId:          fmt.Sprintf("%s%d", param.semesterId, j),
	// 		NmKelas:             row[0],
	// 		TingkatPendidikanId: int32(tingkat),
	// 		JenisRombel:         utils.Int32ToPointer(1),
	// 	}, param.schemaname)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// }
	return nil
}

// fungsi helper untuk mengecekan panjang slice len(row) sebelum akses indeks tertentu.
func safeGet(row []string, idx int) *string {
	if idx < len(row) && len(row[idx]) > 0 {
		return &row[idx]
	}
	return nil
}
