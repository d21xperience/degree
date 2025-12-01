package interceptor

import (
	"auth_service/jwt"
	"auth_service/utils"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	JwtManager *jwt.Manager
}

func NewAuthInterceptor(jwtManager *jwt.Manager) *AuthInterceptor {
	return &AuthInterceptor{JwtManager: jwtManager}
}

var publicMethods = map[string]bool{
	"/auth.AuthService/Login":                     true,
	"/auth.AuthService/Register":                  true,
	"/auth.AuthService/RefreshToken":              true,
	"/auth.SekolahTenantService/GetSekolahTenant": true,
	// "/health.Health/Check":                        true, // contoh service lain
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if publicMethods[info.FullMethod] {
			return handler(ctx, req)
		}
		// Extract token from context
		token, err := jwt.ExtractToken(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}

		// Validate token
		claims, err := i.JwtManager.ValidateToken(token)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}
		// var newCtx context.Context

		ctx = context.WithValue(ctx, utils.CtxUserID, claims.UserID)
		ctx = context.WithValue(ctx, utils.CtxRole, claims.Role)
		log.Printf("Injected userID=%q", claims.UserID) // ← temporary log
		return handler(ctx, req)

	}
}
