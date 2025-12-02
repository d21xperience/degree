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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SemesterServiceServer struct {
	pb.UnimplementedSemesterServiceServer
	repo repositories.GenericRepository[models.Semester]
}

func NewSemesterService() *SemesterServiceServer {
	repoSemester := repositories.NewSemesterRepository(config.DB)
	return &SemesterServiceServer{
		repo: *repoSemester,
	}
}

// **CreateSemester**
func (s *SemesterServiceServer) CreateSemester(ctx context.Context, req *pb.CreateSemesterRequest) (*pb.CreateSemesterResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Semester", "TahunAjaranId", "Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	semester := req.GetSemester()
	// tahunAjaranId := req.GetTahunAjaranId()
	tglMulai, err := utils.StringToTime(semester.TanggalMulai, "2006-01-01")
	if err != nil {
		return nil, err
	}
	tglSelesai, err := utils.StringToTime(semester.TanggalSelesai, "2006-01-01")
	if err != nil {
		return nil, err
	}
	semesterModel := &models.Semester{
		SemesterID:     semester.SemesterId,
		Nama:           semester.NamaSemester,
		TahunAjaranID:  semester.TahunAjaranId,
		Semester:       semester.Semester,
		PeriodeAktif:   semester.PeriodeAktif,
		TanggalMulai:   tglMulai,
		TanggalSelesai: tglSelesai,
	}

	err = s.repo.Save(ctx, semesterModel, req.GetSchemaname())
	if err != nil {
		log.Printf("Gagal menyimpan semester: %v", err)
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &pb.CreateSemesterResponse{
		Message: "Semester berhasil ditambahkan",
		Status:  true,
	}, nil
}

// **GetSemester**
func (s *SemesterServiceServer) GetSemester(ctx context.Context, req *pb.GetSemesterRequest) (*pb.GetSemesterResponse, error) {
	requiredFields := []string{"Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	conditions := make(map[string]any)
	orderBy := []string{"tahun_ajaran_id DESC", "semester DESC"}
	SemesterModels, err := s.repo.FindAllByConditions(ctx, req.GetSchemaname(), conditions, 500, 0, orderBy)
	if err != nil {
		return &pb.GetSemesterResponse{
			Status:  false,
			Message: fmt.Sprintf("gagal mendapatkan data semester pada schema '%s': %v", req.GetSchemaname(), err),
		}, nil
	}
	// Konversi hasil ke response protobuf
	SemesterList := utils.ConvertModelsToPB(SemesterModels, func(model *models.Semester) *pb.Semester {
		return &pb.Semester{
			SemesterId:     model.SemesterID,
			TahunAjaranId:  model.TahunAjaranID,
			NamaSemester:   model.Nama,
			Semester:       model.Semester,
			PeriodeAktif:   model.PeriodeAktif,
			TanggalMulai:   utils.TimeToString(model.TanggalMulai, "2006-01-01"),
			TanggalSelesai: utils.TimeToString(model.TanggalSelesai, "2006-01-01"),
		}
	})
	// Return response
	return &pb.GetSemesterResponse{
		Semester: SemesterList,
		Status:   true,
		Message:  "Berhasil mendapatkan data semester",
	}, nil

}

// **UpdateSemester**
func (s *SemesterServiceServer) UpdateSemester(ctx context.Context, req *pb.UpdateSemesterRequest) (*pb.UpdateSemesterResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Semester", "Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	semesterReq := req.Semester
	tglMulai, err := time.Parse(time.RFC3339, semesterReq.TanggalMulai)
	if err != nil {
		return nil, err
	}
	tglSelesai, err := time.Parse(time.RFC3339, semesterReq.TanggalSelesai)
	if err != nil {
		return nil, err
	}

	SemesterModel := &models.Semester{
		SemesterID:     semesterReq.SemesterId,
		TahunAjaranID:  semesterReq.TahunAjaranId,
		Semester:       semesterReq.Semester,
		Nama:           semesterReq.NamaSemester,
		PeriodeAktif:   semesterReq.PeriodeAktif,
		TanggalMulai:   tglMulai,
		TanggalSelesai: tglSelesai,
	}
	err = s.repo.Update(ctx, SemesterModel, req.GetSchemaname(), "semester_id", SemesterModel.SemesterID)
	if err != nil {
		log.Printf("Gagal memperbarui tahun ajaran: %v", err)
		return nil, fmt.Errorf("gagal memperbarui tahun ajaran: %w", err)
	}
	return &pb.UpdateSemesterResponse{
		Message: "Semester berhasil diperbarui",
		Status:  true,
	}, nil
}

// // **DeleteSemester**
func (s *SemesterServiceServer) DeleteSemester(ctx context.Context, req *pb.DeleteSemesterRequest) (*pb.DeleteSemesterResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"SemesterIds", "Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	// Filter empty strings
	var validIds []string
	for _, id := range req.GetSemesterIds() {
		if id != "" {
			validIds = append(validIds, id)
		}
	}
	if len(validIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "semester_ids cannot be empty")
	}

	err = s.repo.DeleteBatch(ctx, validIds, req.GetSchemaname(), "semester_id", repositories.ValidateString)
	if err != nil {
		log.Printf("Gagal menghapus semester: %v", err)
		return &pb.DeleteSemesterResponse{
			Message: "Semester gagal dihapus",
			Status:  false,
		}, nil
	}

	return &pb.DeleteSemesterResponse{
		Message: "Semester berhasil dihapus",
		Status:  true,
	}, nil
}
