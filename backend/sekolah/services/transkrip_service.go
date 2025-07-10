package services

import (
	"context"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
)

type TranskripNilaiService struct {
	pb.UnimplementedTranskripNilaiServiceServer
	repo repositories.GenericRepository[models.NilaiAkhir]
}

func NewTranskripNilaiService() *TranskripNilaiService {
	repoTranskrip := repositories.NewNilaiAkhirRepository(config.DB)
	return &TranskripNilaiService{
		repo: *repoTranskrip,
	}
}

func (s *TranskripNilaiService) CreateTranskripNilai(ctx context.Context, req *pb.CreateTranskripNilaiRequest) (*pb.CreateTranskripNilaiResponse, error) {

	return &pb.CreateTranskripNilaiResponse{
		Status:  true,
		Message: "sukses",
	}, nil

}
func (s *TranskripNilaiService) GetTranskripNilai(ctx context.Context, req *pb.GetTranskripNilaiRequest) (*pb.GetTranskripNilaiResponse, error) {
	// Debugging: Cek nilai request yang diterima
	// log.Printf("transkrip_service/GetTranskripNilai => Received Sekolah data request: %+v\n", req)
	// // Daftar field yang wajib diisi
	// requiredFields := []string{"Schemaname", "SemesterId"}
	// // Validasi request
	// err := utils.ValidateFields(req, requiredFields)
	// if err != nil {
	// 	return nil, err
	// }
	// schemaName := req.GetSchemaname()
	// semesterId := req.GetSemesterId()
	// joins := []string{
	// 	"JOIN tabel_kelas tk ON tk.rombongan_belajara_id =  ",
	// }
	// preloads := []string{""}
	// conditions := map[string]any{
	// 	"tabel_nilaiakhir.semester_id": semesterId,
	// }

	// transkripModel, err := s.repo.FindWithPreloadAndJoins(ctx, schemaName, joins, preloads, conditions, nil, nil, false)
	// if err != nil {
	// 	return nil, err
	// }
	// transkripPb := utils.ConvertModelsToPB(transkripModel, func(item models.NilaiAkhir) *pb.NilaiAkhir {
	// 	return &pb.NilaiAkhir{
	// 		SemesterId:      item.SemesterId,
	// 		MataPelajaranId: *item.MataPelajaranId,
	// 	}
	// })
	return &pb.GetTranskripNilaiResponse{
		// TranskripNilai: transkripPb,
		Status:  true,
		Message: "sukses",
	}, nil

}
func (s *TranskripNilaiService) UpdateTranskripNilai(ctx context.Context, req *pb.UpdateTranskripNilaiRequest) (*pb.UpdateTranskripNilaiResponse, error) {

	return &pb.UpdateTranskripNilaiResponse{
		Status:  true,
		Message: "sukses",
	}, nil

}
func (s *TranskripNilaiService) DeleteTranskripNilai(ctx context.Context, req *pb.DeleteTranskripNilaiRequest) (*pb.DeleteTranskripNilaiResponse, error) {

	return &pb.DeleteTranskripNilaiResponse{
		Status:  true,
		Message: "sukses",
	}, nil

}
func (s *TranskripNilaiService) UploadTranskripNilai(ctx context.Context, req *pb.UploadTranskripNilaiRequest) (*pb.UploadTranskripNilaiResponse, error) {

	return &pb.UploadTranskripNilaiResponse{
		Status:  true,
		Message: "sukses",
	}, nil

}
