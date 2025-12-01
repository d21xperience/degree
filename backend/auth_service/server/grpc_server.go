package server

import (
	pb "auth_service/generated"
	"auth_service/interceptor"
	"auth_service/jwt"
	"auth_service/services"

	"google.golang.org/grpc"
)

// gRPC Server
func RunGRPCServer(jwtManager *jwt.Manager) *grpc.Server {
	intercep := interceptor.NewAuthInterceptor(jwtManager)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(intercep.Unary()),
	)
	// Register gRPC services
	authServiceServer := services.NewAuthServiceServer(jwtManager)
	pb.RegisterAuthServiceServer(grpcServer, authServiceServer)

	// User
	userServiceServer := services.NewUserUserServiceServer()
	pb.RegisterUserServiceServer(grpcServer, userServiceServer)

	// User Profile
	userProfileServiceServer := services.NewUserProfileServiceServer()
	pb.RegisterUserProfileServiceServer(grpcServer, userProfileServiceServer)

	// Sekolah Tenant
	sekolahTenantServer := services.NewSekolahTenantServer()
	pb.RegisterSekolahTenantServiceServer(grpcServer, sekolahTenantServer)

	return grpcServer
}
