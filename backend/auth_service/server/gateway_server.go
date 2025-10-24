package server

import (
	pb "auth_service/generated"
	"context"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunHTTPGateway(ctx context.Context, gatewayMux *runtime.ServeMux, grpcServerEndpoint, httpPort string) {
	// log.Printf("Connecting HTTP Gateway to gRPC server at %s...", grpcServerEndpoint)
	// conn, err := grpc.Dial(grpcServerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("❌ Cannot connect to gRPC server: %v", err)
	// }
	// defer conn.Close()
	// log.Println("✅ HTTP Gateway successfully connected to gRPC server.")

	// log.Println("Registering AuthService handler...")
	// Gunakan insecure.NewCredentials() sebagai pengganti grpc.WithInsecure()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register gRPC-Gateway handlers
	err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, gatewayMux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register Auth service gRPC Gateway: %v", err)
	}
	// log.Println("AuthService handler registered successfully")
	err = pb.RegisterUserProfileServiceHandlerFromEndpoint(ctx, gatewayMux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register User profile service gRPC Gateway: %v", err)
	}

	err = pb.RegisterSekolahIndonesiaServiceHandlerFromEndpoint(ctx, gatewayMux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register User profile service gRPC Gateway: %v", err)
	}

}
