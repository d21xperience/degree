package server

import (
	pb "auth_service/generated"
	"auth_service/models"
	"auth_service/services"
	"auth_service/utils"
	"log"

	"google.golang.org/grpc"
)

// gRPC Server
func RunGRPCServer(pvKey any) *grpc.Server {
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
		"/auth.AuthService/GetUser": {models.RoleSiswa},

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
