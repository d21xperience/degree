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

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PTKTerdaftarServiceServer struct {
	pb.UnimplementedPTKTerdaftarServiceServer
	repo             repositories.GenericRepository[models.PTKTerdaftar]
	repoPTK          repositories.GenericRepository[models.TabelPTK]
	repoPTKPelengkap repositories.GenericRepository[models.PtkPelengkap]
}

func NewPTKTerdaftarServiceServer() *PTKTerdaftarServiceServer {
	repoPTKTerdaftar := repositories.NewPTKTerdaftarRepository(config.DB)
	repoPTK := repositories.NewPTKRepository(config.DB)
	repoPTKPelengkap := repositories.NewPTKPelengkapRepository(config.DB)
	return &PTKTerdaftarServiceServer{
		repo:             *repoPTKTerdaftar,
		repoPTK:          *repoPTK,
		repoPTKPelengkap: *repoPTKPelengkap,
	}
}

// **CreatePTKTerdaftar**
func (s *PTKTerdaftarServiceServer) CreatePTKTerdaftar(ctx context.Context, req *pb.CreatePTKTerdaftarRequest) (*pb.CreatePTKTerdaftarResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("ptk_terdaftar_server/CreatePTKTerdaftar => Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "PTKTerdaftar"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	PTKTerdaftar := req.GetPtkTerdaftar()
	if req.PtkTerdaftar[0].Ptk.Nama == "" || req.PtkTerdaftar[0].Ptk.Nama == "\"\"" {
		return &pb.CreatePTKTerdaftarResponse{
			Message: "Data yang diinputkan tidak boleh kosong!",
			Status:  false,
		}, nil
	}

	var ptkTerdaftarId uuid.UUID
	var ptkId uuid.UUID
	PTKTerdaftarModel := utils.ConvertPBToModels(PTKTerdaftar, func(item *pb.PTKTerdaftar) *models.PTKTerdaftar {
		// cek apakah data dari dapodik atau bukan
		if !item.Ptk.IsDapodik {
			ptkTerdaftarId = uuid.New()
			ptkId = uuid.New()
		} else {
			ptkTerdaftarId = utils.StringToUUID(item.PtkTerdaftarId)
			ptkId = utils.StringToUUID(item.PtkId)
		}

		return &models.PTKTerdaftar{
			PtkTerdaftarId: ptkTerdaftarId,
			TahunAjaranId:  item.TahunAjaranId,
			PtkID:          ptkId,
			JenisKeluarId:  &item.JenisKeluarId,
			IsDapodik:      item.IsDapodik,
			// PtkTerdaftarIdDapodik: ,
			PTK: models.TabelPTK{
				PtkID:             utils.StringToUUID(item.PtkId),
				Nama:              item.Ptk.Nama,
				JenisPtkID:        item.Ptk.JenisPtkId,
				JenisKelamin:      &item.Ptk.JenisKelamin,
				TempatLahir:       &item.Ptk.TempatLahir,
				TanggalLahir:      utils.TimeToPointer(item.Ptk.TanggalLahir),
				StatusKeaktifanID: item.Ptk.StatusKeaktifanId,
				Agama:             &item.Ptk.Agama,
				IsDapodik:         item.Ptk.IsDapodik,
			},
			PTKPelengkap: models.PtkPelengkap{
				PtkPelengkapId: uuid.New(),
				PtkId:          ptkId,
				GelarDepan:     &item.PtkPelengkap.GelarDepan,
				GelarBelakang:  &item.PtkPelengkap.GelarBelakang,
				Nik:            &item.PtkPelengkap.Nik,
				NoKk:           &item.PtkPelengkap.No_KK,
				Nuptk:          &item.PtkPelengkap.Nuptk,
				Niy:            &item.PtkPelengkap.Niy,
				Nip:            &item.PtkPelengkap.Nip,
				AlamatJalan:    &item.PtkPelengkap.AlamatJalan,
				Rt:             &item.PtkPelengkap.Rt,
				Rw:             &item.PtkPelengkap.Rw,
				DesaKelurahan:  &item.PtkPelengkap.DesaKelurahan,
				Kec:            &item.PtkPelengkap.Kec,
				KabKota:        &item.PtkPelengkap.KabKota,
				Propinsi:       &item.PtkPelengkap.Propinsi,
				KodePos:        &item.PtkPelengkap.KodePos,
				NoTeleponRumah: &item.PtkPelengkap.NoTeleponRumah,
				NoHp:           &item.PtkPelengkap.NoHp,
				Email:          &item.PtkPelengkap.Email,
			},
		}
	})

	// var daftarPTK []models.TabelPTK
	// for _, v := range PTKTerdaftarModel {
	// 	daftarPTK = append(daftarPTK, v.PTK)
	// }
	// conflicts, err := s.repoPTK.SaveManyWithConflictCheck(ctx, schemaName, utils.SliceToPointer(daftarPTK), "PtkID", "ptk_id", 100)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "insert failed: %v", err)
	// }

	// conflictProto := repositories.ConvertConflictsToProto(conflicts, "PtkID", "Nama")

	// simpan ke tabel_ptk_terdaftar
	conflicts1, err := s.repo.SaveManyWithConflictCheck(ctx, schemaName, PTKTerdaftarModel, "PtkTerdaftarId", "ptk_terdaftar_id", 100, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insert failed: %v", err)
	}

	conflictProto2 := repositories.ConvertConflictsToProto(conflicts1, "PtkTerdaftarId", "PtkID")
	// conflictProto2 = append(conflictProto2, conflictProto...)
	fmt.Print(conflictProto2)
	return &pb.CreatePTKTerdaftarResponse{
		Message: "PTKTerdaftar berhasil ditambahkan",
		Status:  true,
		Conflicts: &pb.ConflictResponse{
			Message:       "Sebagian data berhasil disimpan",
			Conflicts:     conflictProto2,
			TotalConflict: int32(len(conflictProto2)),
		},
	}, nil
}

// func (s *PTKTerdaftarServiceServer) CreateBanyakPTKTerdaftar(ctx context.Context, req *pb.CreateBanyakPTKTerdaftarRequest) (*pb.CreateBanyakPTKTerdaftarResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Schemaname", "PTKTerdaftar"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaname()
// 	PTKTerdaftar := req.PTKTerdaftar

// 	PTKTerdaftarModels := ConvertPBToModels(PTKTerdaftar, func(rom *pb.PTKTerdaftar) *models.RombonganBelajar {
// 		return &models.RombonganBelajar{
// 			RombonganBelajarId:  rom.RombonganBelajarId,
// 			SekolahId:           rom.SekolahId,
// 			SemesterId:          rom.SemesterId,
// 			JurusanId:           rom.JurusanId,
// 			PtkId:               rom.PtkId,
// 			NmPTKTerdaftar:      rom.NmPTKTerdaftar,
// 			TingkatPendidikanId: rom.TingkatPendidikanId,
// 			JenisPTKTerdaftar:   rom.JenisPTKTerdaftar,
// 			NamaJurusanSp:       rom.NamaJurusanSp,
// 			JurusanSpId:         rom.JurusanSpId,
// 			KurikulumId:         rom.KurikulumId,
// 		}
// 	})
// 	err = s.repo.SaveMany(ctx, schemaName, PTKTerdaftarModels, 100)
// 	if err != nil {
// 		log.Printf("Gagal menyimpan PTKTerdaftar: %s", err)
// 		return nil, fmt.Errorf("gagal menyimpan PTKTerdaftar: %w", err)
// 	}

// 	return &pb.CreateBanyakPTKTerdaftarResponse{
// 		Message: "PTKTerdaftar berhasil ditambahkan",
// 		Status:  true,
// 	}, nil
// }

// **GetPTKTerdaftar**
func (s *PTKTerdaftarServiceServer) GetPTKTerdaftar(ctx context.Context, req *pb.GetPTKTerdaftarRequest) (*pb.GetPTKTerdaftarResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "TahunAjaranId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}

	// Cek apakah harus mengambil semua data atau data spesifik berdasarkan SemesterId
	joins := []string{
		"JOIN tabel_ptk ON tabel_ptk.ptk_id = tabel_ptk_terdaftar.ptk_id",
	}
	preloads := []string{"PTK", "PTKPelengkap"}

	conditions := map[string]interface{}{
		"tabel_ptk_terdaftar.tahun_ajaran_id": req.GetTahunAjaranId(),
	}
	orderBy := []string{"nama ASC"}
	exactConditions := []struct {
		Query string
		Args  []interface{}
	}{
		{"jenis_ptk_id != ?", []any{9}},
		{"jenis_ptk_id != ?", []any{11}},
		{"jenis_ptk_id != ?", []any{42}},
		{"jenis_ptk_id != ?", []any{44}},
		// {"created_at BETWEEN ? AND ?", []interface{}{startDate, endDate}},
	}
	// groupByColumns := []string{"tabel_ptk_terdaftar.ptk_terdaftar_id"} // Hindari duplikasi
	// PTKTerdaftarModel, err := s.repo.FindWithPreloadAndJoinsOrigin(ctx, schemaName, joins, preloads, conditions, orderBy)
	PTKTerdaftarModel, err := s.repo.FindWithRelations(ctx, schemaName, joins, preloads, conditions, exactConditions, nil, orderBy)
	if err != nil {
		return nil, err
	}
	// Konversi ke protobuf
	ptkTerdaftarPB := utils.ConvertModelsToPB(PTKTerdaftarModel, func(item models.PTKTerdaftar) *pb.PTKTerdaftar {
		return &pb.PTKTerdaftar{
			PtkTerdaftarId: item.PtkTerdaftarId.String(),
			TahunAjaranId:  item.TahunAjaranId,
			Ptk: &pb.PTK{
				PtkId:             item.PTK.PtkID.String(),
				Nama:              item.PTK.Nama,
				JenisKelamin:      utils.SafeString(item.PTK.JenisKelamin),
				TempatLahir:       utils.SafeString(item.PTK.TempatLahir),
				JenisPtkId:        item.PTK.JenisPtkID,
				TanggalLahir:      item.PTK.TanggalLahir.Format("2006-01-02"),
				StatusKeaktifanId: item.PTK.StatusKeaktifanID,
			},
			PtkPelengkap: &pb.PtkPelengkap{
				GelarDepan:     utils.SafeString(item.PTKPelengkap.GelarDepan),
				GelarBelakang:  utils.SafeString(item.PTKPelengkap.GelarBelakang),
				Nik:            utils.SafeString(item.PTKPelengkap.Nik),
				No_KK:          utils.SafeString(item.PTKPelengkap.Nik),
				Nuptk:          utils.SafeString(item.PTKPelengkap.Nuptk),
				Niy:            utils.SafeString(item.PTKPelengkap.Niy),
				Nip:            utils.SafeString(item.PTKPelengkap.Nip),
				AlamatJalan:    utils.SafeString(item.PTKPelengkap.AlamatJalan),
				Rt:             utils.SafeString(item.PTKPelengkap.Rt),
				Rw:             utils.SafeString(item.PTKPelengkap.Rw),
				DesaKelurahan:  utils.SafeString(item.PTKPelengkap.DesaKelurahan),
				Kec:            utils.SafeString(item.PTKPelengkap.Kec),
				KabKota:        utils.SafeString(item.PTKPelengkap.KabKota),
				Propinsi:       utils.SafeString(item.PTKPelengkap.Propinsi),
				KodePos:        utils.SafeString(item.PTKPelengkap.KodePos),
				NoTeleponRumah: utils.SafeString(item.PTKPelengkap.NoTeleponRumah),
				NoHp:           utils.SafeString(item.PTKPelengkap.NoHp),
				Email:          utils.SafeString(item.PTKPelengkap.Email),
			},
		}
	})

	return &pb.GetPTKTerdaftarResponse{
		PtkTerdaftar: ptkTerdaftarPB,
		Message:      "Sukses",
	}, nil
}

// // **UpdatePTKTerdaftar**
func (s *PTKTerdaftarServiceServer) UpdatePTKTerdaftar(ctx context.Context, req *pb.UpdatePTKTerdaftarRequest) (*pb.UpdatePTKTerdaftarResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "PtkTerdaftar"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	PTKTerdaftar := req.GetPtkTerdaftar()
	if len(PTKTerdaftar) == 1 {

		PTKModel := &models.TabelPTK{
			Nama:         PTKTerdaftar[0].Ptk.Nama,
			JenisKelamin: &PTKTerdaftar[0].Ptk.JenisKelamin,
			TempatLahir:  &PTKTerdaftar[0].Ptk.TempatLahir,
			Agama:        &PTKTerdaftar[0].Ptk.Agama,
			TanggalLahir: utils.TimeToPointer(PTKTerdaftar[0].Ptk.TanggalLahir),
		}
		err = s.repoPTK.Update(ctx, PTKModel, schemaName, "ptk_id", PTKTerdaftar[0].Ptk.PtkId)
		if err != nil {
			log.Printf("Gagal memperbaharui PTK: %s", err)
			return &pb.UpdatePTKTerdaftarResponse{
				Message: fmt.Sprintf("gagal memperbaharui PTK: %s", err),
				Status:  false,
			}, nil
		}
		// Cek apakah data sudah tersimpan di tabel_ptk_pelengkap?

		PTKPelengkap := &models.PtkPelengkap{
			// PtkPelengkapId: uuid.New(),
			PtkId:          utils.StringToUUID(PTKTerdaftar[0].Ptk.PtkId),
			Rt:             &PTKTerdaftar[0].PtkPelengkap.Rt,
			Rw:             &PTKTerdaftar[0].PtkPelengkap.Rw,
			DesaKelurahan:  &PTKTerdaftar[0].PtkPelengkap.DesaKelurahan,
			Kec:            &PTKTerdaftar[0].PtkPelengkap.Kec,
			KabKota:        &PTKTerdaftar[0].PtkPelengkap.KabKota,
			Propinsi:       &PTKTerdaftar[0].PtkPelengkap.Propinsi,
			GelarDepan:     &PTKTerdaftar[0].PtkPelengkap.GelarDepan,
			GelarBelakang:  &PTKTerdaftar[0].PtkPelengkap.GelarBelakang,
			Nik:            &PTKTerdaftar[0].PtkPelengkap.Nik,
			NoKk:           &PTKTerdaftar[0].PtkPelengkap.No_KK,
			Nuptk:          &PTKTerdaftar[0].PtkPelengkap.Nuptk,
			Niy:            &PTKTerdaftar[0].PtkPelengkap.Niy,
			Nip:            &PTKTerdaftar[0].PtkPelengkap.Nip,
			AlamatJalan:    &PTKTerdaftar[0].PtkPelengkap.AlamatJalan,
			KodePos:        &PTKTerdaftar[0].PtkPelengkap.KodePos,
			NoTeleponRumah: &PTKTerdaftar[0].PtkPelengkap.NoTeleponRumah,
			NoHp:           &PTKTerdaftar[0].PtkPelengkap.NoHp,
			Email:          &PTKTerdaftar[0].PtkPelengkap.Email,
		}
		// Cek apakah data sudah tersimpan di tabel_ptk_pelengkap?
		cek, err := s.repoPTKPelengkap.FindOrCreateByID(ctx, PTKTerdaftar[0].Ptk.PtkId, schemaName, "ptk_id", func(id string) *models.PtkPelengkap {
			return &models.PtkPelengkap{
				PtkId:         utils.StringToUUID(id),
				Rt:            &PTKTerdaftar[0].PtkPelengkap.Rt,
				Rw:            &PTKTerdaftar[0].PtkPelengkap.Rw,
				DesaKelurahan: &PTKTerdaftar[0].PtkPelengkap.DesaKelurahan,
				Kec:           &PTKTerdaftar[0].PtkPelengkap.Kec,
				KabKota:       &PTKTerdaftar[0].PtkPelengkap.KabKota,
				Propinsi:      &PTKTerdaftar[0].PtkPelengkap.Propinsi,
			}
		})
		if err != nil {
			log.Printf("Gagal membuat PTK Pelengkap: %s", err)
			return &pb.UpdatePTKTerdaftarResponse{
				Message: fmt.Sprintf("gagal membuat PTK Pelengkap: %s", err),
				Status:  false,
			}, nil
		}
		if cek != nil {
			err := s.repoPTKPelengkap.Update(ctx, PTKPelengkap, schemaName, "ptk_id", PTKTerdaftar[0].Ptk.PtkId)
			if err != nil {
				log.Printf("Gagal memperbaharui PTK Pelengkap: %s", err)
				return &pb.UpdatePTKTerdaftarResponse{
					Message: fmt.Sprintf("gagal memperbaharui PTK Pelengkap: %s", err),
					Status:  false,
				}, nil
			}
		}
	}
	return &pb.UpdatePTKTerdaftarResponse{
		Message: "PTKTerdaftar berhasil diperbarui",
		Status:  true,
	}, nil
}

// // **DeletePTKTerdaftar**
func (s *PTKTerdaftarServiceServer) DeletePTKTerdaftar(ctx context.Context, req *pb.DeletePTKTerdaftarRequest) (*pb.DeletePTKTerdaftarResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "PtkTerdaftarId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	PTKTerdaftarID := req.GetPtkTerdaftarId()
	err = s.repo.Delete(ctx, PTKTerdaftarID, schemaName, "ptk_terdaftar_id")
	if err != nil {
		log.Printf("Gagal menghapus PTKTerdaftar: %s", err)
		return &pb.DeletePTKTerdaftarResponse{Message: "Gagal menghapus", Status: false}, fmt.Errorf("gagal menghapus PTKTerdaftar: %w", err)
	}

	return &pb.DeletePTKTerdaftarResponse{
		Message: "PTKTerdaftar berhasil dihapus",
		Status:  true,
	}, nil
}
func (s *PTKTerdaftarServiceServer) DeletBatchPTKTerdaftar(ctx context.Context, req *pb.DeleteBatchPTKTerdaftarRequest) (*pb.DeleteBatchPTKTerdaftarResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "PtkTerdaftarId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	PTKTerdaftarID := req.GetPtkTerdaftarId()
	err = s.repo.DeleteBatch(ctx, PTKTerdaftarID, schemaName, "ptk_terdaftar_id")
	if err != nil {
		log.Printf("Gagal menghapus PTKTerdaftar: %s", err)
		return &pb.DeleteBatchPTKTerdaftarResponse{Message: "Gagal menghapus", Status: false}, fmt.Errorf("gagal menghapus PTKTerdaftar: %w", err)
	}

	return &pb.DeleteBatchPTKTerdaftarResponse{
		Message: "PTKTerdaftar berhasil dihapus",
		Status:  true,
	}, nil
}

func (s *PTKTerdaftarServiceServer) SearchPTKTerdaftar(ctx context.Context, req *pb.SearchPTKTerdaftarRequest) (*pb.SearchPTKTerdaftarResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "PtkTerdaftarId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}

	// Cek apakah harus mengambil semua data atau data spesifik berdasarkan SemesterId
	// joins := []string{
	// 	"JOIN tabel_ptk ON tabel_ptk.ptk_id = tabel_ptk_terdaftar.ptk_id",
	// }
	preloads := []string{"PTK", "PTKPelengkap"}

	conditions := map[string]any{
		"tabel_ptk_terdaftar.ptk_terdaftar_id": req.GetPtkTerdaftarId(),
	}
	PTKTerdaftarModel, err := s.repo.FindWithRelations(ctx, schemaName, nil, preloads, conditions, nil, nil, nil)
	if err != nil {
		return &pb.SearchPTKTerdaftarResponse{
			PtkTerdaftar: nil,
			Message:      fmt.Sprintf("Gagal mendapatkan %s", err),
			Status:       false,
		}, nil
	}
	// Konversi ke protobuf
	ptkTerdaftarPB := utils.ConvertModelsToPB(PTKTerdaftarModel, func(ptk models.PTKTerdaftar) *pb.PTKTerdaftar {
		return &pb.PTKTerdaftar{
			PtkTerdaftarId: ptk.PtkTerdaftarId.String(),
			TahunAjaranId:  ptk.TahunAjaranId,
			Ptk: &pb.PTK{
				PtkId:             ptk.PTK.PtkID.String(),
				Nama:              ptk.PTK.Nama,
				Agama:             utils.SafeString(ptk.PTK.Agama),
				JenisKelamin:      utils.SafeString(ptk.PTK.JenisKelamin),
				TempatLahir:       utils.SafeString(ptk.PTK.TempatLahir),
				JenisPtkId:        ptk.PTK.JenisPtkID,
				TanggalLahir:      ptk.PTK.TanggalLahir.Format("2006-01-02"),
				StatusKeaktifanId: ptk.PTK.StatusKeaktifanID,
			},
			PtkPelengkap: &pb.PtkPelengkap{
				GelarDepan:     utils.SafeString(ptk.PTKPelengkap.GelarDepan),
				GelarBelakang:  utils.SafeString(ptk.PTKPelengkap.GelarBelakang),
				Nik:            utils.SafeString(ptk.PTKPelengkap.Nik),
				No_KK:          utils.SafeString(ptk.PTKPelengkap.Nik),
				Nuptk:          utils.SafeString(ptk.PTKPelengkap.Nuptk),
				Niy:            utils.SafeString(ptk.PTKPelengkap.Niy),
				Nip:            utils.SafeString(ptk.PTKPelengkap.Nip),
				AlamatJalan:    utils.SafeString(ptk.PTKPelengkap.AlamatJalan),
				Rt:             utils.SafeString(ptk.PTKPelengkap.Rt),
				Rw:             utils.SafeString(ptk.PTKPelengkap.Rw),
				DesaKelurahan:  utils.SafeString(ptk.PTKPelengkap.DesaKelurahan),
				Kec:            utils.SafeString(ptk.PTKPelengkap.Kec),
				KabKota:        utils.SafeString(ptk.PTKPelengkap.KabKota),
				Propinsi:       utils.SafeString(ptk.PTKPelengkap.Propinsi),
				KodePos:        utils.SafeString(ptk.PTKPelengkap.KodePos),
				NoTeleponRumah: utils.SafeString(ptk.PTKPelengkap.NoTeleponRumah),
				NoHp:           utils.SafeString(ptk.PTKPelengkap.NoHp),
				Email:          utils.SafeString(ptk.PTKPelengkap.Email),
			},
		}
	})

	return &pb.SearchPTKTerdaftarResponse{
		PtkTerdaftar: ptkTerdaftarPB[0],
		Status:       true,
		Message:      "Behasil mendapatkan data ptk terdaftar",
	}, nil
}
