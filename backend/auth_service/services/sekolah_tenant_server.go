package services

import (
	"auth_service/config"
	pb "auth_service/generated"
	"auth_service/repositories"
	"auth_service/utils"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SekolahTenantServer struct {
	pb.UnimplementedSekolahTenantServiceServer
	repo repositories.SekolahTenantRepository
}

func NewSekolahTenantServer() *SekolahTenantServer {
	repoSekolahTenant := repositories.NewSekolahTenantRepository(config.DB)
	return &SekolahTenantServer{
		repo: repoSekolahTenant,
	}
}

func (s *SekolahTenantServer) GetSekolahTenant(ctx context.Context, req *pb.GetSekolahTenantRequest) (*pb.SekolahTenant, error) {
	log.Printf("sekolah_tenant/GetSekolahTenant received data from request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"npsn"}
	// Validasi request
	requiredFieldsResponse := utils.ValidateFields(req, requiredFields)
	if requiredFieldsResponse != nil {
		return nil, requiredFieldsResponse
	}
	npsn := req.GetNpsn()
	res, err := s.repo.GetSekolahTenantByNPSN(npsn)
	if err != nil {
		return nil, status.Error(codes.NotFound, "sekolah tidak ditemukan!")
	}

	return &pb.SekolahTenant{
		Id:            res.ID,
		NamaSekolah:   res.NamaSekolah,
		Npsn:          res.NPSN,
		EnkripId:      res.EnkripID,
		Kecamatan:     res.Kecamatan,
		Kabupaten:     res.Kabupaten,
		Propinsi:      res.Propinsi,
		KodeKecamatan: res.KodeKecamatan,
		KodeKab:       res.KodeKab,
		KodeProp:      res.KodeProp,
		AlamatJalan:   res.AlamatJalan,
		LogoUrl:       res.LogoUrl,
	}, nil

}

// func (s *SekolahTenantServer) ListSekolahTenant(ctx context.Context, req *pb.ListSekolahTenantRequest) (*pb.ListSekolahTenantResponse, error) {

// }
// func (s *SekolahTenantServer) SearchSekolahTenant(ctx context.Context, req *pb.SearchSekolahTenantRequest) (*pb.ListSekolahTenantResponse, error) {

// }
