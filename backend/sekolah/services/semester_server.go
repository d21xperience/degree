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
	requiredFields := []string{"Semester", "TahunAjaranId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	semester := req.GetSemester()
	// tahunAjaranId := req.GetTahunAjaranId()
	semesterModel := &models.Semester{
		SemesterID:     semester.SemesterId,
		Nama:           semester.NamaSemester,
		TahunAjaranID:  semester.TahunAjaranId,
		Semester:       semester.Semester,
		PeriodeAktif:   semester.PeriodeAktif,
		TanggalMulai:   semester.TanggalMulai,
		TanggalSelesai: semester.TanggalSelesai,
	}

	err = s.repo.Save(ctx, semesterModel, "ref")
	if err != nil {
		log.Printf("Gagal menyimpan semester: %v", err)
		return nil, fmt.Errorf("gagal menyimpan semester: %w", err)
	}

	return &pb.CreateSemesterResponse{
		Message: "Semester berhasil ditambahkan",
		Status:  true,
	}, nil
}

// **GetSemester**
func (s *SemesterServiceServer) GetSemester(ctx context.Context, req *pb.GetSemesterRequest) (*pb.GetSemesterResponse, error) {
	// Cek apakah harus mengambil semua data atau data spesifik berdasarkan SemesterId
	SemesterID := req.GetSemesterId()
	findAll := SemesterID == ""

	if findAll {
		// Ambil data spesifik berdasarkan SemesterId
		conditions := map[string]interface{}{
			"periode_aktif": 1,
		}
		orderBy := []string{
			"semester_id",
		}
		SemesterModels, err := s.repo.FindAllByConditions(ctx, "ref", conditions, 100, 0, orderBy)
		if err != nil {
			log.Printf("[ERROR] Gagal menemukan tahun ajaran di schema '%s': %v", "ref", err)
			return nil, fmt.Errorf("gagal menemukan tahun ajaran di schema '%s': %w", "ref", err)
		}
		// Konversi hasil ke response protobuf
		SemesterList := utils.ConvertModelsToPB(SemesterModels, func(model *models.Semester) *pb.Semester {
			return &pb.Semester{
				SemesterId:     model.SemesterID,
				TahunAjaranId:  model.TahunAjaranID,
				NamaSemester:   model.Nama,
				Semester:       model.Semester,
				PeriodeAktif:   model.PeriodeAktif,
				TanggalMulai:   model.TanggalMulai,
				TanggalSelesai: model.TanggalSelesai,
			}
		})
		// Return response
		return &pb.GetSemesterResponse{
			Semester: SemesterList,
		}, nil
	}

	// Ambil data spesifik berdasarkan SemesterId

	SemesterModel, err := s.repo.FindByID(ctx, SemesterID, "ref", "semester_id")
	if err != nil {
		log.Printf("[ERROR] Gagal menemukan tahun ajaran dengan ID '%s' di schema '%s': %v", SemesterID, "ref", err)
		return nil, fmt.Errorf("gagal menemukan tahun ajaran dengan ID '%s': %w", SemesterID, err)
	}

	return &pb.GetSemesterResponse{
		Semester: []*pb.Semester{
			utils.ConvertModelToPB(SemesterModel, func(model *models.Semester) *pb.Semester {
				return &pb.Semester{
					SemesterId:     model.SemesterID,
					TahunAjaranId:  model.TahunAjaranID,
					NamaSemester:   model.Nama,
					Semester:       model.Semester,
					PeriodeAktif:   model.PeriodeAktif,
					TanggalMulai:   model.TanggalMulai,
					TanggalSelesai: model.TanggalSelesai,
				}
			}),
		},
	}, nil
}

// **UpdateSemester**
func (s *SemesterServiceServer) UpdateSemester(ctx context.Context, req *pb.UpdateSemesterRequest) (*pb.UpdateSemesterResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Semester"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	semesterReq := req.Semester
	SemesterModel := &models.Semester{
		SemesterID:     semesterReq.SemesterId,
		TahunAjaranID:  semesterReq.TahunAjaranId,
		Semester:       semesterReq.Semester,
		Nama:           semesterReq.NamaSemester,
		PeriodeAktif:   semesterReq.PeriodeAktif,
		TanggalMulai:   semesterReq.TanggalMulai,
		TanggalSelesai: semesterReq.TanggalSelesai,
	}
	err = s.repo.Update(ctx, SemesterModel, "ref", "semester_id", SemesterModel.SemesterID)
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
	requiredFields := []string{"SemesterId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	SemesterID := req.GetSemesterId()

	err = s.repo.Delete(ctx, SemesterID, "ref", "semester_id")
	if err != nil {
		log.Printf("Gagal menghapus tahun ajaran: %v", err)
		return nil, fmt.Errorf("gagal menghapus tahun ajaran: %w", err)
	}

	return &pb.DeleteSemesterResponse{
		Message: "Tahun ajaran berhasil dihapus",
		Status:  true,
	}, nil
}
