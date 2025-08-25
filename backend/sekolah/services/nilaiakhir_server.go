package services

import (
	"context"
	"fmt"
	"log"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type NilaiAkhirServiceServer struct {
	pb.UnimplementedNilaiAkhirServiceServer
	repo repositories.GenericRepository[models.NilaiAkhir]
	// repoNilaiSiswa repositories.GenericRepository[models.NilaiSiswa]
	repoRombelAnggota repositories.GenericRepository[models.RombelAnggota]
}

func NewNilaiAkhirServiceServer() *NilaiAkhirServiceServer {
	repoNilaiAkhir := repositories.NewNilaiAkhirRepository(config.DB)
	// repoNilaiSiswa := repositories.NewN
	repoRombelAnggota := repositories.NewRombelAnggotaRepository(config.DB)
	return &NilaiAkhirServiceServer{
		repo:              *repoNilaiAkhir,
		repoRombelAnggota: *repoRombelAnggota,
	}
}

func (s *NilaiAkhirServiceServer) CreateNilaiAkhir(ctx context.Context, req *pb.CreateNilaiAkhirRequest) (*pb.CreateNilaiAkhirResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "NilaiAkhir"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	nilaiAkhirReq := req.GetNilaiAkhir()

	nilaiAkhir := utils.ConvertPBToModels(nilaiAkhirReq, func(nilai *pb.NilaiAkhir) *models.NilaiAkhir {
		anggotaRombelID, err := uuid.Parse(nilai.AnggotaRombelId)
		if err != nil {
			log.Printf("Invalid UUID for AnggotaRombelID: %v", err)
		}
		pesertaDidikID, err := uuid.Parse(nilai.AnggotaRombelId)
		if err != nil {
			log.Printf("Invalid UUID for AnggotaRombelID: %v", err)
		}
		return &models.NilaiAkhir{
			IdNilaiAkhir:    uuid.New(),
			AnggotaRombelId: anggotaRombelID,
			MataPelajaranId: &nilai.MataPelajaranId,
			SemesterId:      nilai.SemesterId,
			NilaiPeng:       &nilai.NilaiPeng,
			PredikatPeng:    nilai.PredikatPeng,
			NilaiKet:        &nilai.NilaiKet,
			PredikatKet:     nilai.PredikatKet,
			NilaiSik:        &nilai.NilaiSik,
			PredikatSik:     nilai.PredikatSik,
			NilaiSikSos:     &nilai.NilaiSiksos,
			PredikatSikSos:  nilai.PredikatSiksos,
			PesertaDidikId:  pesertaDidikID,
			IDMinat:         nilai.IdMinat,
		}
	})
	err = s.repo.SaveMany(ctx, schemaName, nilaiAkhir, 100)
	if err != nil {
		log.Printf("Gagal menyimpan Nilai akhir: %v", err)
		return nil, fmt.Errorf("gagal menyimpan Nilai akhir: %w", err)
	}

	return &pb.CreateNilaiAkhirResponse{
		Message: "Nilai akhir berhasil ditambahkan",
		Status:  true,
	}, nil
}

func (s *NilaiAkhirServiceServer) SearchNilaiAkhir(ctx context.Context, req *pb.SearchNilaiAkhirRequest) (*pb.SearchNilaiAkhirResponse, error) {
	log.Printf("nilaiakhir_server/SearchNilaiAkhir received data from request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"SchemaName", "SemesterId", "PesertaDidikId"}
	// Validasi request
	requiredFieldsResponse := utils.ValidateFields(req, requiredFields)
	if requiredFieldsResponse != nil {
		return nil, requiredFieldsResponse
	}

	schemaName := req.GetSchemaname()
	pesertaDidikId := utils.StringToUUID(req.GetPesertaDidikId())
	modelNilaiAkhir, jenjang, err := GetNilaiSiswaPerSemester(config.DB, schemaName, req.SemesterId, pesertaDidikId)
	if err != nil {
		return &pb.SearchNilaiAkhirResponse{
			Status:  false,
			Message: "Gagal mendapatkan nilai",
		}, nil
	}
	results := utils.ConvertModelsToPB(modelNilaiAkhir, func(item models.NilaiMapel) *pb.NilaiMapel {
		return &pb.NilaiMapel{
			MataPelajaran: item.MataPelajaran,
			Semester1:     utils.SafeUint32(item.Semester1),
			Semester2:     utils.SafeUint32(item.Semester2),
			Semester3:     utils.SafeUint32(item.Semester3),
			Semester4:     utils.SafeUint32(item.Semester4),
			Semester5:     utils.SafeUint32(item.Semester5),
			Semester6:     utils.SafeUint32(item.Semester6),
			Semester7:     utils.SafeUint32(item.Semester7),
			Semester8:     utils.SafeUint32(item.Semester8),
			Semester9:     utils.SafeUint32(item.Semester9),
			Semester10:    utils.SafeUint32(item.Semester10),
			Semester11:    utils.SafeUint32(item.Semester11),
			Semester12:    utils.SafeUint32(item.Semester12),
		}
	})

	return &pb.SearchNilaiAkhirResponse{
		Status:  true,
		Message: "Berhasil mendapatkan nilai",
		Nilai: &pb.NilaiSiswa{
			NmSiswa:    "",
			NmKelas:    "",
			Jenjang:    jenjang,
			NilaiMapel: results,
		},
	}, nil
}

// **GetNilai akhir**
// func (s *NilaiAkhirServiceServer) GetNilaiSiswa(ctx context.Context, req *pb.GetNilaiSiswaRequest) (*pb.GetNilaiSiswaResponse, error) {
// 	requiredFields := []string{"Schemaname", "SemesterId"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaname()
// 	rombelId := req.GetRombonganBelajarId()
// 	handlerNilai := handler.NewNilaiSiswaHandler()
// 	nilaiSiswa, err := handlerNilai.GetNilaiSiswa(ctx, schemaName, req.SemesterId, rombelId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	NilaiAkhirList := utils.ConvertModelsToPB(nilaiSiswa, func(item handler.NilaiSiswa) *pb.NilaiSiswa {
// 		return &pb.NilaiSiswa{
// 			NmSiswa:             item.NmSiswa,
// 			NmKelas:             item.NmKelas,
// 			TingkatPendidikanId: item.TingkatPendidikanId,
// 			SemesterId:          item.SemesterId,
// 			NilaiAkhir: utils.ConvertModelsToPB(*item.NilaiMapel, func(item handler.NilaiMapel) *pb.NilaiMapel {
// 				return &pb.NilaiMapel{
// 					NmMapel:         item.NmMapel,
// 					MataPelajaranId: item.MataPelajaranId,
// 					SemesterId:      item.SemesterId,
// 					NilaiPeng:       item.NilaiPeng,
// 				}
// 			}),
// 		}
// 	})
// 	return &pb.GetNilaiSiswaResponse{
// 		NilaiSiswa: NilaiAkhirList,
// 		Status:     true,
// 		Message:    "berhasil mengambil data",
// 	}, nil
// }

// **UpdateNilai akhir**
// func (s *NilaiAkhirServiceServer) UpdateNilai akhir(ctx context.Context, req *pb.UpdateNilai akhirRequest) (*pb.UpdateNilai akhirResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received UpdateUserProfile request: %+v\n", req)
// 	schemaName := req.GetSchemaname()
// 	Nilai akhirReq := req.GetNilai akhir()
// 	Nilai akhirPelenReq := req.GetNilai akhirPelengkap()
// 	Nilai akhir := &models.PesertaDidik{
// 		PesertaDidikID:  Nilai akhirReq.PesertaDidikID,
// 		NIS:             Nilai akhirReq.NIS,
// 		NISN:            Nilai akhirReq.NISN,
// 		NamaNilai akhir:       Nilai akhirReq.NamaNilai akhir,
// 		TempatLahir:     Nilai akhirReq.TempatLahir,
// 		TanggalLahir:    Nilai akhirReq.TanggalLahir,
// 		JenisKelamin:    Nilai akhirReq.JenisKelamin,
// 		Agama:           Nilai akhirReq.Agama,
// 		AlamatNilai akhir:     &Nilai akhirReq.AlamatNilai akhir,
// 		TeleponNilai akhir:    Nilai akhirReq.TeleponNilai akhir,
// 		DiterimaTanggal: Nilai akhirReq.DiterimaTanggal,
// 		NamaAyah:        Nilai akhirReq.NamaAyah,
// 		NamaIbu:         Nilai akhirReq.NamaIbu,
// 		PekerjaanAyah:   Nilai akhirReq.PekerjaanAyah,
// 		PekerjaanIbu:    Nilai akhirReq.PekerjaanIbu,
// 		NamaWali:        &Nilai akhirReq.NamaWali,
// 		PekerjaanWali:   &Nilai akhirReq.PekerjaanWali,
// 	}
// 	Nilai akhirPelenkap := &models.PesertaDidikPelengkap{
// 		PelengkapNilai akhirID: Nilai akhirPelenReq.PelengkapNilai akhirID,
// 		PesertaDidikID:   &Nilai akhirPelenReq.PesertaDidikID,
// 		StatusDalamKel:   &Nilai akhirPelenReq.StatusDalamKel,
// 		AnakKe:           &Nilai akhirPelenReq.AnakKe,
// 		SekolahAsal:      Nilai akhirPelenReq.SekolahAsal,
// 		DiterimaNilai akhir:    &Nilai akhirPelenReq.DiterimaNilai akhir,
// 		AlamatOrtu:       &Nilai akhirPelenReq.AlamatOrtu,
// 		TeleponOrtu:      &Nilai akhirPelenReq.TeleponOrtu,
// 		AlamatWali:       &Nilai akhirPelenReq.AlamatWali,
// 		TeleponWali:      &Nilai akhirPelenReq.TeleponWali,
// 		FotoNilai akhir:        &Nilai akhirPelenReq.FotoNilai akhir,
// 	}
// 	err := s.pesertaDidikService.Update(ctx, Nilai akhir, Nilai akhirPelenkap, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal memperbarui Nilai akhir: %v", err)
// 		return nil, fmt.Errorf("gagal memperbarui Nilai akhir: %w", err)
// 	}

// 	return &pb.UpdateNilai akhirResponse{
// 		Message: "Nilai akhir berhasil diperbarui",
// 		Status:  true,
// 	}, nil
// }

// // // **DeleteNilai akhir**
// func (s *NilaiAkhirServiceServer) DeleteNilai akhir(ctx context.Context, req *pb.DeleteNilai akhirRequest) (*pb.DeleteNilai akhirResponse, error) {
// 	schemaName := req.GetSchemaname()
// 	Nilai akhirID := req.GetNilai akhirId()

// 	err := s.pesertaDidikService.Delete(ctx, Nilai akhirID, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menghapus Nilai akhir: %v", err)
// 		return nil, fmt.Errorf("gagal menghapus Nilai akhir: %w", err)
// 	}

// 	return &pb.DeleteNilai akhirResponse{
// 		Message: "Nilai akhir berhasil dihapus",
// 		Status:  true,
// 	}, nil
// }

// // UploadNilai akhir mengunggah data Nilai akhir dari file Excel
// func (s *NilaiAkhirServiceServer) UploadNilai akhir(ctx context.Context, req *pb.UploadNilai akhirRequest) (*pb.UploadNilai akhirResponse, error) {
// 	schemaName := req.GetSchemaname()
// 	fileData := req.GetFile() // File dalam bentuk byte array

// 	// Simpan file ke sementara
// 	tempFile := "/tmp/uploaded_Nilai akhir.xlsx"
// 	err := saveFile(tempFile, fileData)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menyimpan file sementara: %w", err)
// 	}

// 	// Baca file Excel
// 	f, err := excelize.OpenFile(tempFile)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
// 	}
// 	defer f.Close()

// 	// Ambil semua data dari sheet pertama
// 	rows, err := f.GetRows(f.GetSheetName(0))
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
// 	}

// 	// Pastikan ada data
// 	if len(rows) < 2 {
// 		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
// 	}

// 	// Validasi header
// 	expectedHeaders := []string{"NIS", "NISN", "NamaNilai akhir", "TempatLahir", "TanggalLahir", "JenisKelamin", "Agama"}
// 	for i, expected := range expectedHeaders {
// 		if rows[0][i] != expected {
// 			return nil, fmt.Errorf("format kolom tidak sesuai, kolom '%s' seharusnya ada di posisi %d", expected, i+1)
// 		}
// 	}

// 	var Nilai akhirList []*models.PesertaDidik

// 	// Mulai dari baris kedua karena baris pertama adalah header
// 	for _, row := range rows[1:] {
// 		if len(row) < len(expectedHeaders) {
// 			log.Println("Skipping row due to insufficient data:", row)
// 			continue
// 		}

// 		// Konversi data sesuai dengan model
// 		namaNilai akhir := row[2]
// 		nis := row[0]
// 		nisn := row[1]
// 		tempatLahir := row[3]
// 		tanggalLahir := row[4]
// 		jenisKelamin := row[5]
// 		agama := row[6]

// 		// Validasi data
// 		if nis == "" || namaNilai akhir == "" || nisn == "" {
// 			log.Println("Skipping row due to missing required fields:", row)
// 			continue
// 		}

// 		// Konversi angka
// 		nisInt, err := strconv.Atoi(nis)
// 		if err != nil {
// 			log.Printf("Format NIS tidak valid: %s", nis)
// 			continue
// 		}

// 		nisnInt, err := strconv.Atoi(nisn)
// 		if err != nil {
// 			log.Printf("Format NISN tidak valid: %s", nisn)
// 			continue
// 		}

// 		// Masukkan ke dalam list
// 		Nilai akhir := &models.PesertaDidik{
// 			NIS:          strconv.Itoa(nisInt),
// 			NISN:         strconv.Itoa(nisnInt),
// 			NamaNilai akhir:    namaNilai akhir,
// 			TempatLahir:  tempatLahir,
// 			TanggalLahir: tanggalLahir,
// 			JenisKelamin: jenisKelamin,
// 			Agama:        agama,
// 		}
// 		Nilai akhirList = append(Nilai akhirList, Nilai akhir)
// 	}

// 	// Simpan ke database
// 	err = s.pesertaDidikService.BatchSave(ctx, Nilai akhirList, schemaName)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menyimpan data Nilai akhir ke database: %w", err)
// 	}

//		return &pb.UploadNilai akhirResponse{
//			Message: "Nilai akhir berhasil diunggah",
//			Total:   int32(len(Nilai akhirList)),
//			Status:  true,
//		}, nil
//	}
func hitungUrutanSemester(semesterID string, tahunMasuk int) int {
	if len(semesterID) < 5 {
		return 0
	}

	tahun := cast.ToInt(semesterID[:4])
	sem := cast.ToInt(semesterID[4:])

	return (tahun-tahunMasuk)*2 + sem
}
func GetNilaiSiswaPerSemester(db *gorm.DB, schemaname, semesterId string, pesertaDidikID uuid.UUID) ([]models.NilaiMapel, string, error) {
	var result []models.NilaiMapel
	var jenjang string

	// Dapatkan jenjang dan tahun masuk dari rombel
	var rombel struct {
		Jenjang    string
		TahunMasuk int
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaname))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// query := tx.Table(fmt.Sprintf("%s.tabel_kelas", strings.ToLower(schemaname)))
		query := tx.Table("tabel_kelas")

		joins := []string{
			"JOIN tabel_anggotakelas ak ON ak.rombongan_belajar_id = tabel_kelas.rombongan_belajar_id ",
			"JOIN tabel_sekolah ts ON ts.sekolah_id = tabel_kelas.sekolah_id",
			"JOIN ref.jenjang_pendidikan jp ON jp.jenjang_pendidikan_id = ts.jenjang_pendidikan_id",
			"JOIN tabel_siswa tsis ON tsis.peserta_didik_id = ak.peserta_didik_id",
		}
		for _, join := range joins {
			query = query.Joins(join)
		}

		query = query.Where("ak.peserta_didik_id = ?", pesertaDidikID).
			Group("jp.nama, tsis.diterima_tanggal").
			Select("jp.nama as jenjang, EXTRACT(YEAR FROM tsis.diterima_tanggal) as tahun_masuk")
			// Scan(&rombel).Error
		if err := query.Scan(&rombel).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, "", err
	}
	jenjang = rombel.Jenjang
	tahunMasuk := rombel.TahunMasuk // Misal: 2022

	// Tentukan jumlah semester maksimal
	totalSemesters := map[string]int{
		"PAUD": 2, "TK / sederajat": 2, "SD / sederajat": 12, "SMP / sederajat": 6, "SMA / sederajat": 6,
	}[jenjang]
	if totalSemesters == 0 {
		totalSemesters = 6
	}

	// Ambil data nilai + nama mata pelajaran
	var rows []struct {
		SemesterID    string
		MataPelajaran string
		NilaiPeng     uint32
	}
	err = db.Table("tabel_nilaiakhir n").
		Joins("JOIN ref.mata_pelajaran m ON m.mata_pelajaran_id = n.mata_pelajaran_id").
		Where("n.peserta_didik_id = ?", pesertaDidikID).Where("n.semester_id <= ?", semesterId).
		Select("n.semester_id, m.nama as mata_pelajaran, COALESCE(n.nilai_peng, 0) as nilai_peng").
		Scan(&rows).Error

	if err != nil {
		return nil, "", err
	}

	// Mapping nilai ke struct
	mapping := make(map[string]*models.NilaiMapel)
	for _, row := range rows {
		urutan := hitungUrutanSemester(row.SemesterID, tahunMasuk)

		// Hanya proses semester 1 sampai max (misal 6 untuk SMA)
		if urutan < 1 || urutan > totalSemesters {
			continue
		}

		if _, exists := mapping[row.MataPelajaran]; !exists {
			mapping[row.MataPelajaran] = &models.NilaiMapel{
				MataPelajaran: row.MataPelajaran,
			}
		}

		// Assign ke field semester ke-N
		target := mapping[row.MataPelajaran]
		switch urutan {
		case 1:
			target.Semester1 = &row.NilaiPeng
		case 2:
			target.Semester2 = &row.NilaiPeng
		case 3:
			target.Semester3 = &row.NilaiPeng
		case 4:
			target.Semester4 = &row.NilaiPeng
		case 5:
			target.Semester5 = &row.NilaiPeng
		case 6:
			target.Semester6 = &row.NilaiPeng
		case 7:
			target.Semester7 = &row.NilaiPeng
		case 8:
			target.Semester8 = &row.NilaiPeng
		case 9:
			target.Semester9 = &row.NilaiPeng
		case 10:
			target.Semester10 = &row.NilaiPeng
		case 11:
			target.Semester11 = &row.NilaiPeng
		case 12:
			target.Semester12 = &row.NilaiPeng
		}
	}

	// Konversi map ke slice
	for _, v := range mapping {
		result = append(result, *v)
	}

	return result, jenjang, nil
}
func GetTranskripNilai(db *gorm.DB, schemaname string, pesertaDidikID uuid.UUID) ([]models.NilaiMapel, string, error) {
	var result []models.NilaiMapel
	var jenjang string

	// Dapatkan jenjang dan tahun masuk dari rombel
	var rombel struct {
		Jenjang    string
		TahunMasuk int
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaname))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// query := tx.Table(fmt.Sprintf("%s.tabel_kelas", strings.ToLower(schemaname)))
		query := tx.Table("tabel_kelas")

		joins := []string{
			"JOIN tabel_anggotakelas ak ON ak.rombongan_belajar_id = tabel_kelas.rombongan_belajar_id ",
			"JOIN tabel_sekolah ts ON ts.sekolah_id = tabel_kelas.sekolah_id",
			"JOIN ref.jenjang_pendidikan jp ON jp.jenjang_pendidikan_id = ts.jenjang_pendidikan_id",
			"JOIN tabel_siswa tsis ON tsis.peserta_didik_id = ak.peserta_didik_id",
		}
		for _, join := range joins {
			query = query.Joins(join)
		}

		query = query.Where("ak.peserta_didik_id = ?", pesertaDidikID).
			Group("jp.nama, tsis.diterima_tanggal").
			Select("jp.nama as jenjang, EXTRACT(YEAR FROM tsis.diterima_tanggal) as tahun_masuk")
			// Scan(&rombel).Error
		if err := query.Scan(&rombel).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, "", err
	}
	jenjang = rombel.Jenjang
	tahunMasuk := rombel.TahunMasuk // Misal: 2022

	// Tentukan jumlah semester maksimal
	totalSemesters := map[string]int{
		"PAUD": 2, "TK / sederajat": 2, "SD / sederajat": 12, "SMP / sederajat": 6, "SMA / sederajat": 6,
	}[jenjang]
	if totalSemesters == 0 {
		totalSemesters = 6
	}

	// Ambil data nilai + nama mata pelajaran
	var rows []struct {
		SemesterID    string
		MataPelajaran string
		NilaiPeng     uint32
	}
	err = db.Table("tabel_nilaiakhir n").
		Joins("JOIN ref.mata_pelajaran m ON m.mata_pelajaran_id = n.mata_pelajaran_id").
		Where("n.peserta_didik_id = ?", pesertaDidikID).
		Select("n.semester_id, m.nama as mata_pelajaran, COALESCE(n.nilai_peng, 0) as nilai_peng").
		Scan(&rows).Error

	if err != nil {
		return nil, "", err
	}

	// Mapping nilai ke struct
	mapping := make(map[string]*models.NilaiMapel)
	for _, row := range rows {
		urutan := hitungUrutanSemester(row.SemesterID, tahunMasuk)

		// Hanya proses semester 1 sampai max (misal 6 untuk SMA)
		if urutan < 1 || urutan > totalSemesters {
			continue
		}

		if _, exists := mapping[row.MataPelajaran]; !exists {
			mapping[row.MataPelajaran] = &models.NilaiMapel{
				MataPelajaran: row.MataPelajaran,
			}
		}

		// Assign ke field semester ke-N
		target := mapping[row.MataPelajaran]
		switch urutan {
		case 1:
			target.Semester1 = &row.NilaiPeng
		case 2:
			target.Semester2 = &row.NilaiPeng
		case 3:
			target.Semester3 = &row.NilaiPeng
		case 4:
			target.Semester4 = &row.NilaiPeng
		case 5:
			target.Semester5 = &row.NilaiPeng
		case 6:
			target.Semester6 = &row.NilaiPeng
		case 7:
			target.Semester7 = &row.NilaiPeng
		case 8:
			target.Semester8 = &row.NilaiPeng
		case 9:
			target.Semester9 = &row.NilaiPeng
		case 10:
			target.Semester10 = &row.NilaiPeng
		case 11:
			target.Semester11 = &row.NilaiPeng
		case 12:
			target.Semester12 = &row.NilaiPeng
		}
	}

	// Konversi map ke slice
	for _, v := range mapping {
		result = append(result, *v)
	}

	return result, jenjang, nil
}
