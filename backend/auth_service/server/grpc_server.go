package server

import (
	pb "auth_service/generated"
	"auth_service/services"

	"google.golang.org/grpc"
)

// var UploadService *services.UploadServiceServer

func RunGRPCServer(pvKey any) *grpc.Server {
	// gRPC Server
	grpcServer := grpc.NewServer()
	// Register gRPC services
	authServiceServer := services.NewAuthServiceServer(pvKey)
	pb.RegisterAuthServiceServer(grpcServer, authServiceServer)

	// User
	userServiceServer := services.NewUserUserServiceServer()
	pb.RegisterUserServiceServer(grpcServer, userServiceServer)

	// User Profile
	userProfileServiceServer := services.NewUserProfileServiceServer()
	pb.RegisterUserProfileServiceServer(grpcServer, userProfileServiceServer)

	// sekolahIndonesiaServer := services.NewSekolahIndonesiaServer()
	// pb.RegisterSekolahServiceServer(grpcServer, sekolahIndonesiaServer)

	return grpcServer
}
