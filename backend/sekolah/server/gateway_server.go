package server

import (
	"context"
	"log"
	pb "sekolah/generated"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunHTTPGateway(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint, httpPort string) {
	// Gunakan insecure.NewCredentials() sebagai pengganti grpc.WithInsecure()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register gRPC service ke HTTP Gateway
	err := pb.RegisterSekolahServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	err = pb.RegisterTahunAjaranServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Tahun Ajaran Gateway: %v", err)
	}

	err = pb.RegisterSemesterServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Semester Gateway: %v", err)
	}
	err = pb.RegisterSiswaServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Siswa Gateway: %v", err)
	}
	err = pb.RegisterKelasServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Kelas Gateway: %v", err)
	}
	err = pb.RegisterAnggotaKelasServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Kelas Gateway: %v", err)
	}
	err = pb.RegisterNilaiAkhirServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Nilai akhir Gateway: %v", err)
	}
	err = pb.RegisterUploadDataSekolahServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Upload data Sekolah Gateway: %v", err)
	}
	err = pb.RegisterPTKServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Upload PTK service Gateway: %v", err)
	}
	err = pb.RegisterPTKTerdaftarServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Upload PTK Terdaftar service Gateway: %v", err)
	}
	err = pb.RegisterDashboardServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Dashboard Service Gateway: %v", err)
	}
	err = pb.RegisterReferensiServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Reference Table Service Gateway: %v", err)
	}
	err = pb.RegisterPembelajaranServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Pembelajaran Service Gateway: %v", err)
	}
	err = pb.RegisterKenaikanServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Kenaikan Service Gateway: %v", err)
	}
	err = pb.RegisterIjazahServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Ijazah Service Gateway: %v", err)
	}
}
