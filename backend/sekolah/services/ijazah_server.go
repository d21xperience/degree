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
	"time"

	"github.com/google/uuid"
)

type IjazahServiceServer struct {
	pb.UnimplementedIjazahServiceServer
	repo             repositories.GenericRepository[models.Ijazah]
	repoKelas        repositories.GenericRepository[models.RombonganBelajar]
	repoAnggotaKelas repositories.GenericRepository[models.RombelAnggota]
	repoSiswa        repositories.GenericRepository[models.PesertaDidik]
	repoSekolah      repositories.SekolahRepository
	repoDns          repositories.GenericRepository[models.DataNominasiSementara]
	repoInfoIjazah   repositories.GenericRepository[models.TabelInformasiIjazah]
}

func NewIjazahServiceServer() *IjazahServiceServer {
	repoIjazah := repositories.NewIjazahRepository(config.DB)
	repoKelas := repositories.NewrombonganBelajarRepository(config.DB)
	repoAnggotaKelas := repositories.NewRombelAnggotaRepository(config.DB)
	repoSiswa := repositories.NewSiswaRepository(config.DB)
	repoSekolah := repositories.NewSekolahRepository(config.DB)
	repoDns := repositories.NewDnsRepository(config.DB)
	repoInfoIjazah := repositories.NewInfoIjazahRepository(config.DB)

	return &IjazahServiceServer{
		repo:             *repoIjazah,
		repoKelas:        *repoKelas,
		repoAnggotaKelas: *repoAnggotaKelas,
		repoSiswa:        *repoSiswa,
		repoDns:          *repoDns,
		repoSekolah:      repoSekolah,
		repoInfoIjazah:   *repoInfoIjazah,
	}
}

func (s *IjazahServiceServer) CreateDns(ctx context.Context, req *pb.CreateDnsRequest) (*pb.CreateDnsResponse, error) {
	log.Printf("createDns request: %+v\n", req)

	var err error
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "TahunAjaranId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}
	pbDns := req.GetDataNominasiSementara()
	dns := utils.ConvertPBToModels(pbDns, func(item *pb.Dns) *models.DataNominasiSementara {
		tglLahir, err := utils.StringToTime(item.TanggalLahir, "2006-01-02")
		if err != nil {
			return nil
		}
		var tglIjazah time.Time
		if item.TanggalIjazah != "" {
			tglIjazah, err = utils.StringToTime(item.TanggalIjazah, "2006-01-02")
			if err != nil {
				return nil
			}
		}
		return &models.DataNominasiSementara{
			ID:                          uuid.New(),
			PesertaDidikId:              utils.StringToUUID(item.PesertaDidikId),
			RombonganBelajarId:          utils.StringToUUID(item.RombonganBelajarId),
			Nama:                        item.Nama,
			NIS:                         item.Nis,
			NISN:                        item.Nisn,
			TempatLahir:                 item.TempatLahir,
			TanggalLahir:                tglLahir,
			NamaOrtuWali:                item.NamaOrtuWali,
			NPSN:                        item.Npsn,
			ProgramKeahlian:             item.ProgramKeahlian,
			KabupatenKota:               item.KabupatenKota,
			SekolahPenyelenggaraUjianUS: item.SekolahPenyelenggaraUjianUs,
			SekolahPenyelenggaraUjianUN: item.SekolahPenyelenggaraUjianUn,
			PaketKeahlian:               item.PaketKeahlian,
			Provinsi:                    item.Provinsi,
			SekolahID:                   item.SekolahId,
			AsalSekolah:                 item.AsalSekolah,
			NomorIjazah:                 item.NomorIjazah,
			TempatIjazah:                item.TempatIjazah,
			TanggalIjazah:               tglIjazah,
			TahunAjaranId:               req.GetTahunAjaranId(),
			// IsComplete:                  true,
		}
	})
	err = s.repoDns.SaveMany(ctx, schemaName, dns, 100)
	if err != nil {
		return nil, err
	}
	return &pb.CreateDnsResponse{
		Message: "Data nominasi sementara berhasil disimpan",
		Status:  true,
	}, nil
}
func (s *IjazahServiceServer) UpdateDns(ctx context.Context, req *pb.UpdateDnsRequest) (*pb.UpdateDnsResponse, error) {
	var err error
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "SemesterId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}
	pbDns := req.GetDataNominasiSementara()
	dns := utils.ConvertPBToModels(pbDns, func(item *pb.Dns) *models.DataNominasiSementara {
		tglLahir, err := utils.StringToTime(item.TanggalLahir, "2006-01-02")
		if err != nil {
			return nil
		}
		tglIjazah, err := utils.StringToTime(item.TanggalIjazah, "2006-01-02")
		if err != nil {
			return nil
		}
		return &models.DataNominasiSementara{
			PesertaDidikId:              utils.StringToUUID(item.PesertaDidikId),
			RombonganBelajarId:          utils.StringToUUID(item.RombonganBelajarId),
			Nama:                        item.Nama,
			NIS:                         item.Nis,
			NISN:                        item.Nisn,
			TempatLahir:                 item.TempatLahir,
			TanggalLahir:                tglLahir,
			NamaOrtuWali:                item.NamaOrtuWali,
			NPSN:                        item.Npsn,
			ProgramKeahlian:             item.ProgramKeahlian,
			KabupatenKota:               item.KabupatenKota,
			SekolahPenyelenggaraUjianUS: item.SekolahPenyelenggaraUjianUn,
			SekolahPenyelenggaraUjianUN: item.SekolahPenyelenggaraUjianUn,
			PaketKeahlian:               item.PaketKeahlian,
			Provinsi:                    item.Provinsi,
			SekolahID:                   item.SekolahId,
			AsalSekolah:                 item.AsalSekolah,
			NomorIjazah:                 item.NomorIjazah,
			TempatIjazah:                item.TempatIjazah,
			TanggalIjazah:               tglIjazah,
			// IsComplete:                  true,
		}
	})
	err = s.repoDns.SaveMany(ctx, schemaName, dns, 100)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateDnsResponse{
		Message: "Data nominasi sementara berhasil disimpan",
		Status:  true,
	}, nil
}
func (s *IjazahServiceServer) DeleteDns(ctx context.Context, req *pb.DeleteDnsRequest) (*pb.DeleteDnsResponse, error) {
	var err error
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "PesertaDidikId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	pesertaDidikId := req.GetPesertaDidikId()

	err = s.repo.Delete(ctx, pesertaDidikId, schemaName, "peserta_didik_id")
	if err != nil {
		log.Printf("Gagal menghapus DNS: %v", err)
		return nil, fmt.Errorf("gagal menghapus DNS: %w", err)
	}

	return &pb.DeleteDnsResponse{
		Message: "DNS berhasil dihapus",
		Status:  true,
	}, nil
}
func (s *IjazahServiceServer) GetDns(ctx context.Context, req *pb.GetDnsRequest) (*pb.GetDnsResponse, error) {
	var err error
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "TahunAjaranId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}
	joins := []string{
		"JOIN tabel_kelas ON tabel_kelas.rombongan_belajar_id = data_nominasi_sementara.rombongan_belajar_id",
	}
	preloads := []string{"RombonganBelajar", "PesertaDidik"}
	conditions := map[string]any{
		"data_nominasi_sementara.tahun_ajaran_id": req.TahunAjaranId,
		"data_nominasi_sementara.is_complete":     false,
	}
	orderBy := []string{"tabel_kelas.nm_kelas ASC", "data_nominasi_sementara.nama ASC"}
	DnsModels, err := s.repoDns.FindWithPreloadAndJoins(ctx, schemaName, joins, preloads, conditions, nil, orderBy, false)
	if err != nil {
		return nil, err
	}
	dns := utils.ConvertModelsToPB(DnsModels, func(item models.DataNominasiSementara) *pb.Dns {
		return &pb.Dns{
			PesertaDidikId:              item.PesertaDidikId.String(),
			RombonganBelajarId:          item.RombonganBelajarId.String(),
			Nama:                        item.Nama,
			Nis:                         item.NIS,
			Nisn:                        item.NISN,
			TempatLahir:                 item.TempatLahir,
			TanggalLahir:                item.TanggalLahir.Format("2006-01-02"),
			NamaOrtuWali:                item.NamaOrtuWali,
			Npsn:                        item.NPSN,
			ProgramKeahlian:             item.ProgramKeahlian,
			KabupatenKota:               item.KabupatenKota,
			SekolahPenyelenggaraUjianUs: item.SekolahPenyelenggaraUjianUS,
			SekolahPenyelenggaraUjianUn: item.SekolahPenyelenggaraUjianUN,
			PaketKeahlian:               item.PaketKeahlian,
			Provinsi:                    item.Provinsi,
			SekolahId:                   item.SekolahID,
			AsalSekolah:                 item.AsalSekolah,
			NomorIjazah:                 item.NomorIjazah,
			TempatIjazah:                item.TempatIjazah,
			TanggalIjazah:               item.TanggalLahir.Format("2006-01-02"),
			IsComplete:                  item.IsComplete,
			Kelas: &pb.Kelas{
				NmKelas: item.RombonganBelajar.NmKelas,
			},
			Siswa: &pb.Siswa{
				JenisKelamin: item.PesertaDidik.JenisKelamin,
			},
		}
	})

	return &pb.GetDnsResponse{
		DataNominasiSementara: dns,
		Message:               "Data nominasi sementara berhasil diambil",
		Status:                true,
	}, nil
}

func (s *IjazahServiceServer) SearchDns(ctx context.Context, req *pb.SearchDnsRequest) (*pb.SearchDnsResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "TahunAjaranId", "PesertaDidikId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}
	pesertaDidikId := req.GetPesertaDidikId()
	joins := []string{
		"JOIN tabel_siswa ON tabel_siswa.peserta_didik_id = data_nominasi_sementara.peserta_didik_id",
	}
	preloads := []string{"PesertaDidik", "RombonganBelajar"}
	conditions := map[string]any{"data_nominasi_sementara.tahun_ajaran_id": req.GetTahunAjaranId(), "data_nominasi_sementara.peserta_didik_id": pesertaDidikId}
	pdModels, err := s.repoDns.FindWithPreloadAndJoins(ctx, schemaName, joins, preloads, conditions, nil, nil, false)
	if err != nil {
		return &pb.SearchDnsResponse{
			Status:  false,
			Message: "DNS gagal diambil",
		}, nil
	}
	if len(pdModels) == 0 {
		return &pb.SearchDnsResponse{
			Status:  false,
			Message: fmt.Sprintf("peserta_didik_id %s tidak ditemukan", pesertaDidikId),
		}, nil
	}

	return &pb.SearchDnsResponse{
		Status:  true,
		Message: "DNS berhasil diambil",
		DataNominasiSementara: &pb.Dns{
			PesertaDidikId:              pdModels[0].PesertaDidikId.String(),
			RombonganBelajarId:          pdModels[0].RombonganBelajarId.String(),
			Nama:                        pdModels[0].Nama,
			Nis:                         pdModels[0].NIS,
			Nisn:                        pdModels[0].NISN,
			TempatLahir:                 pdModels[0].TempatLahir,
			TanggalLahir:                pdModels[0].TanggalLahir.Format("2006-01-02"),
			NamaOrtuWali:                pdModels[0].NamaOrtuWali,
			Npsn:                        pdModels[0].NPSN,
			ProgramKeahlian:             pdModels[0].ProgramKeahlian,
			KabupatenKota:               pdModels[0].KabupatenKota,
			SekolahPenyelenggaraUjianUs: pdModels[0].SekolahPenyelenggaraUjianUS,
			SekolahPenyelenggaraUjianUn: pdModels[0].SekolahPenyelenggaraUjianUN,
			PaketKeahlian:               pdModels[0].PaketKeahlian,
			Provinsi:                    pdModels[0].Provinsi,
			SekolahId:                   pdModels[0].SekolahID,
			AsalSekolah:                 pdModels[0].AsalSekolah,
			NomorIjazah:                 pdModels[0].NomorIjazah,
			TempatIjazah:                pdModels[0].TempatIjazah,
			TanggalIjazah:               pdModels[0].TanggalLahir.Format("2006-01-02"),
			IsComplete:                  pdModels[0].IsComplete,

			Kelas: &pb.Kelas{
				RombonganBelajarId:  pdModels[0].RombonganBelajar.RombonganBelajarId.String(),
				SekolahId:           pdModels[0].RombonganBelajar.SekolahId.String(),
				SemesterId:          pdModels[0].RombonganBelajar.SemesterId,
				JurusanId:           pdModels[0].RombonganBelajar.JurusanId,
				PtkId:               pdModels[0].RombonganBelajar.PtkID.String(),
				NmKelas:             pdModels[0].RombonganBelajar.NmKelas,
				TingkatPendidikanId: pdModels[0].RombonganBelajar.TingkatPendidikanId,
				JenisRombel:         pdModels[0].RombonganBelajar.JenisRombel,
				NamaJurusanSp:       pdModels[0].RombonganBelajar.NamaJurusanSp,
				KurikulumId:         pdModels[0].RombonganBelajar.KurikulumId,
			},
			Siswa: &pb.Siswa{
				Nis:           pdModels[0].PesertaDidik.Nis,
				Nisn:          pdModels[0].PesertaDidik.Nisn,
				NmSiswa:       pdModels[0].PesertaDidik.NmSiswa,
				TempatLahir:   pdModels[0].PesertaDidik.TempatLahir,
				TanggalLahir:  pdModels[0].PesertaDidik.TanggalLahir.Format("2006-01-02"),
				JenisKelamin:  pdModels[0].PesertaDidik.JenisKelamin,
				Agama:         pdModels[0].PesertaDidik.Agama,
				AlamatSiswa:   utils.SafeString(pdModels[0].PesertaDidik.AlamatSiswa),
				TeleponSiswa:  pdModels[0].PesertaDidik.TeleponSiswa,
				NmAyah:        pdModels[0].PesertaDidik.NmAyah,
				NmIbu:         pdModels[0].PesertaDidik.NmIbu,
				PekerjaanAyah: pdModels[0].PesertaDidik.PekerjaanAyah,
				PekerjaanIbu:  pdModels[0].PesertaDidik.PekerjaanIbu,
				NmWali:        utils.SafeString(pdModels[0].PesertaDidik.NmWali),
				PekerjaanWali: utils.SafeString(pdModels[0].PesertaDidik.PekerjaanWali),
				// DiterimaTanggal: utils.TimeToString(*siswa.PesertaDidik.DiterimaTanggal, "2006-01-02"),
				// DiterimaTanggal: utils.SafeString(*siswa.PesertaDidik.DiterimaTanggal),
			},
		},
	}, nil

}

// =========================================
// info ijazah
func (s *IjazahServiceServer) CreateInfoIjazah(ctx context.Context, req *pb.CreateInfoIjazahRequest) (*pb.CreateInfoIjazahResponse, error) {
	log.Printf("createDns request: %+v\n", req)

	var err error
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "InfoIjazah"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}
	infoIjazah := req.GetInfoIjazah()
	tglDikeluarkanIjazah, err := utils.StringToTime(infoIjazah.TglDikeluarkanIjazah, "2006-01-02")
	if err != nil {
		return nil, err
	}
	infoIjazahModel := &models.TabelInformasiIjazah{
		TahunAjaranID:            &infoIjazah.TahunAjaranId,
		TempatDikeluarkanIjazah:  &infoIjazah.TempatDikeluarkanIjazah,
		TanggalDikeluarkanIjazah: &tglDikeluarkanIjazah,
		PtkID:                    utils.UUIDToPointer(utils.StringToUUID(infoIjazah.PtkId)),
		KopSekolahURL:            &infoIjazah.KopSekolahUrl,
	}
	err = s.repoInfoIjazah.Save(ctx, infoIjazahModel, schemaName)
	if err != nil {
		return &pb.CreateInfoIjazahResponse{
			Status:  false,
			Message: "Info sertifikat gagal disimpan",
		}, nil
	}
	return &pb.CreateInfoIjazahResponse{
		Status:  true,
		Message: "Berhahsil",
	}, nil
}
