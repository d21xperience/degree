package services

import (
	"context"
	"fmt"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"
)

type DashboardServiceServer struct {
	pb.UnimplementedDashboardServiceServer
	repoSiswa repositories.GenericRepository[models.RombelAnggota]
	repoGuru  repositories.GenericRepository[models.PTKTerdaftar]
	repoKelas repositories.GenericRepository[models.RombonganBelajar]
}

func NewDashboardServiceServer() *DashboardServiceServer {
	repoSiswa := repositories.NewRombelAnggotaRepository(config.DB)
	repoGuru := repositories.NewPTKTerdaftarRepository(config.DB)
	repoKelas := repositories.NewrombonganBelajarRepository(config.DB)
	return &DashboardServiceServer{
		repoSiswa: *repoSiswa,
		repoGuru:  *repoGuru,
		repoKelas: *repoKelas,
	}
}

func (s *DashboardServiceServer) GetDashboard(ctx context.Context, req *pb.GetDashboardRequest) (*pb.GetDashboardResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "SemesterId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	semesterId := req.GetSemesterId()
	tahunAjaranId := semesterId[:4]
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}
	countSiswa, err := s.repoSiswa.CountRows(ctx, schemaName, map[string]any{"semester_id": semesterId})
	if err != nil {
		return nil, err
	}
	countGuru, err := s.repoGuru.CountRows(ctx, schemaName, map[string]any{"tahun_ajaran_id": tahunAjaranId})
	if err != nil {
		return nil, err
	}
	countKelas, err := s.repoKelas.CountRows(ctx, schemaName, map[string]any{"semester_id": semesterId, "jenis_rombel": 1})
	if err != nil {
		return nil, err
	}

	return &pb.GetDashboardResponse{
		CountSiswa: countSiswa,
		CountGuru:  countGuru,
		CountKelas: countKelas,
	}, nil
}
