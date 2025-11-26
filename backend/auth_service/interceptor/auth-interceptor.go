package interceptor

import (
	"auth_service/jwt"
	"context"

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

type contextKey string

const UserIDKey contextKey = "user_id"

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		// Skip auth for login endpoint
		if info.FullMethod == "/auth.v1.AuthService/Login" {
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

		// Add claims to context
		ctx = context.WithValue(ctx, UserIDKey, claims.ID)
		// ctx = context.WithValue(ctx, "username", claims.Username)

		return handler(ctx, req)
	}
}
