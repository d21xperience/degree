package server

import (
	"log"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/services"
	"sekolah/utils"

	"google.golang.org/grpc"
)

var UploadService *services.UploadServiceServer

func RunGRPCServer() *grpc.Server {
	publicKey, err := utils.LoadPublicKey()
	if err != nil {
		log.Printf("Error load public key: %v", err)
	}
	allowlist := map[string]bool{
		"/auth.AuthService/Login":        true,
		"/auth.AuthService/RefreshToken": true,
		"/auth.AuthService/JWKS":         true,
	}

	// method -> allowed roles
	methodRoles := map[string][]string{
		// no role restriction, but requires authentication:
		"/auth.AuthService/GetUser": {models.RoleAdmin},

		// strict role requirement:
		// "/sc.v1.SmartContractService/IssueDegree": {"admin", "issuer"},
		// if omitted or empty slice => any authenticated user allowed
	}

	rbac := &utils.RBACInterceptorConfig{
		MethodAllowlist: allowlist,
		MethodRoles:     methodRoles,
		PublicKey:       publicKey,
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(utils.AuthInterceptor(rbac)),
	)

	sekolahService := services.NewSekolahService()
	pb.RegisterSekolahServiceServer(grpcServer, sekolahService)

	tahunAjaranService := services.NewTahunAjararanService()
	pb.RegisterTahunAjaranServiceServer(grpcServer, tahunAjaranService)

	semesterService := services.NewSemesterService()
	pb.RegisterSemesterServiceServer(grpcServer, semesterService)

	// REGISTER SISWA
	siswaService := services.NewSiswaServiceServer()
	pb.RegisterSiswaServiceServer(grpcServer, siswaService)

	// REGISTER KELAS
	kelasService := services.NewRombelServiceServer()
	pb.RegisterKelasServiceServer(grpcServer, kelasService)

	// REGISTER ANGGOTA KELAS
	anggotaKelasService := services.NewRombelAnggotaService()
	pb.RegisterAnggotaKelasServiceServer(grpcServer, anggotaKelasService)

	// REGISTER ANGGOTA KELAS
	nilaiAkhirService := services.NewNilaiAkhirServiceServer()
	pb.RegisterNilaiAkhirServiceServer(grpcServer, nilaiAkhirService)

	// REGISTER UPLOAD SERVICE
	UploadService := services.NewUploadServiceServer()
	pb.RegisterUploadDataSekolahServiceServer(grpcServer, UploadService)

	// REGISTER PTKSERVICE
	ptkService := services.NewPTKServiceServer()
	pb.RegisterPTKServiceServer(grpcServer, ptkService)

	// REGISTER PTK TERDAFTAR SERVICE
	ptkTerdaftarService := services.NewPTKTerdaftarServiceServer()
	pb.RegisterPTKTerdaftarServiceServer(grpcServer, ptkTerdaftarService)

	// REGISTER DASHBOARD SERVICE
	dashboardService := services.NewDashboardServiceServer()
	pb.RegisterDashboardServiceServer(grpcServer, dashboardService)

	// REGISTER REFERENSI TABEL SERVICE
	referensiService := services.NewReferensiServiceServer()
	pb.RegisterReferensiServiceServer(grpcServer, referensiService)

	// REGISTER PEMBELAJARAN SERVICE
	pembelajaranService := services.NewPembelajaranServiceServer()
	pb.RegisterPembelajaranServiceServer(grpcServer, pembelajaranService)

	// REGISTER KENAIKAN SERVICE
	kenaikanService := services.NewKenaikanServiceServer()
	pb.RegisterKenaikanServiceServer(grpcServer, kenaikanService)

	// REGISTER PEMBELAJARAN SERVICE
	ijazahService := services.NewIjazahServiceServer()
	pb.RegisterIjazahServiceServer(grpcServer, ijazahService)

	return grpcServer
}
