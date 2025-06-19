package services

import (
	"context"
	"fmt"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"
	"strings"
)

type ReferensiServiceServer struct {
	pb.UnimplementedReferensiServiceServer
	repoBentukPendidikan       repositories.GenericRepository[models.BentukPendidikan]
	repoJenjangPendidikan      repositories.GenericRepository[models.JenjangPendidikan]
	repoTingkatPendidikan      repositories.GenericRepository[models.TingkatPendidikan]
	repoStatusKepemilikan      repositories.GenericRepository[models.StatusKepemilikan]
	repoKurikulum              repositories.GenericRepository[models.Kurikulum]
	repoJurusan                repositories.GenericRepository[models.Jurusan]
	repoMapel                  repositories.GenericRepository[models.MataPelajaran]
	repoGelarAkademik          repositories.GenericRepository[models.GelarAkademik]
	repoMataPelajaranKurikulum repositories.GenericRepository[models.MataPelajaranKurikulum]
}

func NewReferensiServiceServer() *ReferensiServiceServer {
	repoBentukPendidikan := repositories.NewBentukPendidikanRepository(config.DB)
	repoJenjangPendidikan := repositories.NewJenjangPendidikanRepository(config.DB)
	repoTingkatPendidikan := repositories.NewTingkatPendidikanRepository(config.DB)
	repoStatusKepemilikan := repositories.NewStatusKepemilikanRepository(config.DB)
	repoKurikulum := repositories.NewKurikulumRepository(config.DB)
	repoJurusan := repositories.NewJurusanRepository(config.DB)
	repoMapel := repositories.NewMapelRepository(config.DB)
	repoGelarAkademik := repositories.NewGelarAkademikRepository(config.DB)
	repoMapelKurikulum := repositories.NewMapelKurikulumRepository(config.DB)

	return &ReferensiServiceServer{
		repoBentukPendidikan:       *repoBentukPendidikan,
		repoJenjangPendidikan:      *repoJenjangPendidikan,
		repoTingkatPendidikan:      *repoTingkatPendidikan,
		repoStatusKepemilikan:      *repoStatusKepemilikan,
		repoKurikulum:              *repoKurikulum,
		repoJurusan:                *repoJurusan,
		repoMapel:                  *repoMapel,
		repoGelarAkademik:          *repoGelarAkademik,
		repoMataPelajaranKurikulum: *repoMapelKurikulum,
	}
}

func (s *ReferensiServiceServer) GetBentukPendidikan(ctx context.Context, req *pb.Empty) (*pb.GetBentukPendidikanResponse, error) {
	mod, err := s.repoBentukPendidikan.FindAll(ctx, "ref", 100, 0)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(re *models.BentukPendidikan) *pb.BentukPendidikan {
		return &pb.BentukPendidikan{
			BentukPendidikanId:   uint32(re.BentukPendidikanID),
			Nama:                 re.Nama,
			JenjangPaud:          uint32(re.JenjangPaud),
			JenjangTk:            uint32(re.JenjangTk),
			JenjangSd:            uint32(re.JenjangSd),
			JenjangSmp:           uint32(re.JenjangSmp),
			JenjangSma:           uint32(re.JenjangSma),
			JenjangTinggi:        uint32(re.JenjangTinggi),
			DirektoratPembinaan:  *re.DirektoratPembinaan,
			Aktif:                uint32(re.Aktif),
			FormalitasPendidikan: re.FormalitasPendidikan,
		}
	})
	return &pb.GetBentukPendidikanResponse{
		BentukPendidikan: res,
	}, nil
}
func (s *ReferensiServiceServer) GetJenjang(ctx context.Context, req *pb.Empty) (*pb.GetJenjangResponse, error) {
	mod, err := s.repoJenjangPendidikan.FindAll(ctx, "ref", 100, 0)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(re *models.JenjangPendidikan) *pb.Jenjang {
		return &pb.Jenjang{
			JenjangPendidikanId: uint32(re.JenjangPendidikanID),
			Nama:                re.Nama,
			JenjangLembaga:      uint32(re.JenjangLembaga),
			JenjangOrang:        uint32(re.JenjangOrang),
		}
	})
	return &pb.GetJenjangResponse{
		Jenjang: res,
	}, nil
}
func (s *ReferensiServiceServer) GetTingkatPendidikan(ctx context.Context, req *pb.GetTingkatPendidikanRequest) (*pb.GetTingkatPendidikanResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"JenjangPendidikanId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	conditions := map[string]interface{}{
		"jenjang_pendidikan_id": req.GetJenjangPendidikanId(),
	}
	mod, err := s.repoTingkatPendidikan.FindAllByConditions(ctx, "ref", conditions, 100, 0, nil)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(re *models.TingkatPendidikan) *pb.TingkatPendidikan {
		return &pb.TingkatPendidikan{
			TingkatPendidikanId: int32(re.TingkatPendidikanID),
			JenjangPendidikanId: uint32(re.JenjangPendidikanID),
			Kode:                re.Kode,
			Nama:                re.Nama,
		}
	})
	return &pb.GetTingkatPendidikanResponse{
		TingkatPendidikan: res,
	}, nil
}
func (s *ReferensiServiceServer) GetStatusKepemilikan(ctx context.Context, req *pb.Empty) (*pb.GetStatusKepemilikanResponse, error) {
	mod, err := s.repoStatusKepemilikan.FindAll(ctx, "ref", 100, 0)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(re *models.StatusKepemilikan) *pb.StatusKepemilikan {
		return &pb.StatusKepemilikan{
			StatusKepemilikanId: uint32(re.StatusKepemilikanID),
			Nama:                re.Nama,
		}
	})
	return &pb.GetStatusKepemilikanResponse{
		StatusKepemilikan: res,
	}, nil
}
func (s *ReferensiServiceServer) GetKurikulum(ctx context.Context, req *pb.GetKurikulumRequest) (*pb.GetKurikulumResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"JenjangPendidikanId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	// jika SMK
	/* SELECT * FROM kurikulum k
	WHERE k.jenjang_pendidikan_id = 6 AND LENGTH(k.jurusan_id) = 8 AND k.nama_kurikulum LIKE '%SMK%'
	ORDER BY k.mulai_berlaku ASC */
	// SD, SMP, dll. field jurusan_id IS NULL

	conditions := map[string]any{
		"kurikulum.jenjang_pendidikan_id": req.GetJenjangPendidikanId(),
	}
	var costumConditions []struct {
		Query string
		Args  []any
	}
	if req.JenjangPendidikanStr != nil {
		switch strings.ToUpper(req.GetJenjangPendidikanStr()) {
		case "SMA":
			costumConditions = append(costumConditions, struct {
				Query string
				Args  []any
			}{
				Query: "ref.kurikulum.nama_kurikulum ILIKE " + fmt.Sprintf("'%%%s%%'", req.GetJenjangPendidikanStr()),
				Args:  []any{},
			})
		case "SMK":
			costumConditions = append(costumConditions, struct {
				Query string
				Args  []any
			}{
				Query: "ref.kurikulum.nama_kurikulum ILIKE ?",
				Args:  []any{fmt.Sprintf("%%%s%%", req.GetJenjangPendidikanStr())},
			}, struct {
				Query string
				Args  []any
			}{
				Query: "LENGTH(ref.kurikulum.jurusan_id) = ?",
				Args:  []any{8},
			},
			)

		}

	}

	orderBy := []string{"ref.kurikulum.mulai_berlaku DESC"}
	modelKurikulum, err := s.repoKurikulum.FindWithRelations(ctx, "ref", nil, nil, conditions, costumConditions, nil, orderBy)
	if err != nil {
		return &pb.GetKurikulumResponse{
			Status:    false,
			Message:   fmt.Sprintf("Gagal mengambil data,  %s", err),
			Kurikulum: nil,
		}, err
	}
	res := utils.ConvertModelsToPB(modelKurikulum, func(item models.Kurikulum) *pb.Kurikulum {
		return &pb.Kurikulum{
			KurikulumId:         uint32(item.KurikulumID),
			NamaKurikulum:       item.NamaKurikulum,
			MulaiBerlaku:        item.MulaiBerlaku.Format("2006-01-02"),
			SistemSks:           uint32(item.SistemSKS),
			TotalSks:            uint32(item.TotalSKS),
			JenjangPendidikanId: uint32(item.JenjangPendidikanID),
			JurusanId:           utils.SafeString(item.JurusanID),
		}
	})
	return &pb.GetKurikulumResponse{
		Status:    true,
		Message:   "Berhasil mengambil kurikulum",
		Kurikulum: res,
	}, nil
}
func (s *ReferensiServiceServer) GetBidangKeahlian(ctx context.Context, req *pb.GetBidangKeahlianRequest) (*pb.GetBidangKeahlianResponse, error) {
	conditions := map[string]any{
		"ref.jurusan.level_bidang_id": "10",
	}
	if req.JurusanInduk != nil {
		conditions["ref.jurusan.jurusan_id"] = req.GetJurusanInduk()
	}
	orderBy := []string{"ref.jurusan.jurusan_id", "ref.jurusan.nama_jurusan"}
	mod, err := s.repoJurusan.FindAllByConditions(ctx, "ref", conditions, 1000, 0, orderBy)
	if err != nil {
		return &pb.GetBidangKeahlianResponse{
			Status:         false,
			Message:        "Gagal mengambil bidang keahlian",
			BidangKeahlian: nil,
		}, nil
	}
	res := utils.ConvertModelsToPB(mod, func(item *models.Jurusan) *pb.Jurusan {
		return &pb.Jurusan{
			JurusanId:           item.JurusanID,
			NamaJurusan:         item.NamaJurusan,
			UntukSma:            uint32(item.UntukSMA),
			UntukSmk:            uint32(item.UntukSMK),
			UntukPt:             uint32(item.UntukPT),
			UntukSlb:            uint32(item.UntukSLB),
			UntukSmklb:          uint32(item.UntukSMKLB),
			JenjangPendidikanId: utils.PointerToUint32(utils.Uint16ToUint32Pointer(item.JenjangPendidikanID)),
			JurusanInduk:        utils.SafeString(item.JurusanInduk),
			LevelBidangId:       item.LevelBidangID,
		}
	})
	return &pb.GetBidangKeahlianResponse{
		Status:         true,
		Message:        "Berhasil mengambil bidang keahlian",
		BidangKeahlian: res,
	}, nil
}
func (s *ReferensiServiceServer) GetProgramKeahlian(ctx context.Context, req *pb.GetProgramKeahlianRequest) (*pb.GetProgramKeahlianResponse, error) {
	conditions := map[string]any{
		"ref.jurusan.level_bidang_id": "11",
	}
	if req.JurusanInduk != nil {
		conditions["ref.jurusan.jurusan_induk"] = req.GetJurusanInduk()
	}
	orderBy := []string{"ref.jurusan.jurusan_id", "ref.jurusan.nama_jurusan"}
	mod, err := s.repoJurusan.FindAllByConditions(ctx, "ref", conditions, 1000, 0, orderBy)
	if err != nil {
		return &pb.GetProgramKeahlianResponse{
			Status:          false,
			Message:         "Gagal mengambil program keahlian",
			ProgramKeahlian: nil,
		}, nil
	}
	res := utils.ConvertModelsToPB(mod, func(item *models.Jurusan) *pb.Jurusan {
		return &pb.Jurusan{
			JurusanId:           item.JurusanID,
			NamaJurusan:         item.NamaJurusan,
			UntukSma:            uint32(item.UntukSMA),
			UntukSmk:            uint32(item.UntukSMK),
			UntukPt:             uint32(item.UntukPT),
			UntukSlb:            uint32(item.UntukSLB),
			UntukSmklb:          uint32(item.UntukSMKLB),
			JenjangPendidikanId: utils.PointerToUint32(utils.Uint16ToUint32Pointer(item.JenjangPendidikanID)),
			JurusanInduk:        utils.SafeString(item.JurusanInduk),
			LevelBidangId:       item.LevelBidangID,
		}
	})
	return &pb.GetProgramKeahlianResponse{
		Status:          true,
		Message:         "Berhasil mengambil program keahlian",
		ProgramKeahlian: res,
	}, nil
}
func (s *ReferensiServiceServer) GetJurusan(ctx context.Context, req *pb.GetJurusanRequest) (*pb.GetJurusanResponse, error) {
	conditions := map[string]any{
		"ref.jurusan.level_bidang_id": "12",
	}
	if req.JurusanInduk != nil {
		conditions["ref.jurusan.jurusan_induk"] = req.GetJurusanInduk()
	}
	orderBy := []string{"ref.jurusan.jurusan_id", "ref.jurusan.nama_jurusan"}
	mod, err := s.repoJurusan.FindAllByConditions(ctx, "ref", conditions, 1000, 0, orderBy)
	if err != nil {
		return &pb.GetJurusanResponse{
			Status:  false,
			Message: "Gagal mengambil data jurusan",
			Jurusan: nil,
		}, nil
	}
	res := utils.ConvertModelsToPB(mod, func(item *models.Jurusan) *pb.Jurusan {
		return &pb.Jurusan{
			JurusanId:           item.JurusanID,
			NamaJurusan:         item.NamaJurusan,
			UntukSma:            uint32(item.UntukSMA),
			UntukSmk:            uint32(item.UntukSMK),
			UntukPt:             uint32(item.UntukPT),
			UntukSlb:            uint32(item.UntukSLB),
			UntukSmklb:          uint32(item.UntukSMKLB),
			JenjangPendidikanId: utils.PointerToUint32(utils.Uint16ToUint32Pointer(item.JenjangPendidikanID)),
			JurusanInduk:        utils.SafeString(item.JurusanInduk),
			LevelBidangId:       item.LevelBidangID,
		}
	})
	return &pb.GetJurusanResponse{
		Status:  true,
		Message: "Berhasil mengambil jurusan",
		Jurusan: res,
	}, nil
}

func (s *ReferensiServiceServer) GetMapel(ctx context.Context, req *pb.GetMapelRequest) (*pb.GetMapelResponse, error) {
	// // Daftar field yang wajib diisi
	// requiredFields := []string{"Param"}
	// // Validasi request
	// err := utils.ValidateFields(req, requiredFields)
	// if err != nil {
	// 	return nil, err
	// }
	conditions := map[string]any{}
	mapelId := req.GetMapelId()
	if mapelId != "" {
		conditions["mapel_id"] = mapelId
	}
	mod, err := s.repoMapel.FindAllByConditions(ctx, "ref", conditions, 7000, 0, nil)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(item *models.MataPelajaran) *pb.Mapel {
		return &pb.Mapel{
			MataPelajaranId:     item.MataPelajaranID,
			Nama:                item.Nama,
			PilihanSekolah:      item.PilihanSekolah,
			PilihanBuku:         item.PilihanBuku,
			PilihanKepengawasan: item.PilihanKepengawasan,
			PilihanEvaluasi:     item.PilihanEvaluasi,
			JurusanId:           utils.SafeString(item.JurusanID),
		}
	})
	return &pb.GetMapelResponse{
		Mapel: res,
	}, nil
}

func (s *ReferensiServiceServer) GetGelarAkademik(ctx context.Context, req *pb.GetGelarAkademikRequest) (*pb.GetGelarAkademikResponse, error) {
	gelarAkademikModel, err := s.repoGelarAkademik.FindAll(ctx, "ref", 1000, 0)
	if err != nil {
		return nil, err
	}
	results := utils.ConvertModelsToPB(gelarAkademikModel, func(item *models.GelarAkademik) *pb.GelarAkademik {
		return &pb.GelarAkademik{
			Kode:        item.Kode,
			Nama:        item.Nama,
			PosisiGelar: item.PosisiGelar,
		}
	})
	return &pb.GetGelarAkademikResponse{
		GelarAkademik: results,
	}, nil
}

func (s *ReferensiServiceServer) GetMapelKurikulum(ctx context.Context, req *pb.GetMapelKurikulumRequest) (*pb.GetMapelKurikulumResponse, error) {
	requiredFields := []string{"KurikulumId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}

	joins := []string{
		"JOIN ref.mata_pelajaran ON ref.mata_pelajaran.mata_pelajaran_id = ref.mata_pelajaran_kurikulum.mata_pelajaran_id ",
	}
	conditions := map[string]any{
		"ref.mata_pelajaran_kurikulum.kurikulum_id": req.KurikulumId,
		// "ref.mata_pelajaran_kurikulum.tingkat_pendidikan_id": modelRombel.TingkatPendidikanId,
	}
	preloads := []string{"MataPelajaran"}

	modelsMapelKurikulum, err := s.repoMataPelajaranKurikulum.FindWithPreloadAndJoins(ctx, "ref", joins, preloads, conditions, nil, nil, false)
	if err != nil {
		return &pb.GetMapelKurikulumResponse{
			Mapel:   nil,
			Status:  true,
			Message: fmt.Sprintf("Gagal mengambil mapel kukrikulum %v", err),
		}, nil
	}
	pbMapelKurikulum := utils.ConvertModelsToPB(modelsMapelKurikulum, func(item models.MataPelajaranKurikulum) pb.MapelKurikulum {
		return pb.MapelKurikulum{
			KurikulumId:     item.KurikulumId,
			MataPelajaranId: item.MataPelajaranId,
			Nama:            item.MataPelajaran.Nama,
		}
	})
	return &pb.GetMapelKurikulumResponse{
		Mapel:   utils.SliceToPointer(pbMapelKurikulum),
		Status:  true,
		Message: "Berhasil mengambil mapel kukrikulum",
	}, nil
}
