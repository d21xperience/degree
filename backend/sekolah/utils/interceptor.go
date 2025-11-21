package utils

import (
	"context"
	"crypto/rsa"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// RBACInterceptorConfig menyimpan aturan RBAC
type RBACInterceptorConfig struct {
	// MethodAllowlist : method-method yang tidak perlu auth (full method names)
	MethodAllowlist map[string]bool

	// MethodRoles : mapping fullMethodName -> allowed roles (any of these suffices)
	// e.g. "/auth.v1.AuthService/GetProfile": []string{"user","admin"}
	MethodRoles map[string][]string

	// PublicKey RSA public key untuk verifikasi JWT
	PublicKey *rsa.PublicKey
}

// containsRole checks whether roles contain at least one required role
func containsAnyRole(userRoles []string, allowed []string) bool {
	if len(allowed) == 0 {
		// empty allowed means only authentication required (no role check)
		return true
	}
	m := make(map[string]struct{}, len(userRoles))
	for _, r := range userRoles {
		m[r] = struct{}{}
	}
	for _, a := range allowed {
		if _, ok := m[a]; ok {
			return true
		}
	}
	return false
}

func AuthInterceptor(rbac *RBACInterceptorConfig) grpc.UnaryServerInterceptor {
	if rbac == nil {
		panic("rbac tidak boleh bernilai nil")
	}
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		// 1) Allowlist (public methods)
		if rbac.MethodAllowlist != nil && rbac.MethodAllowlist[info.FullMethod] {
			// skip auth for these methods
			return handler(ctx, req)
		}

		// Extract token from metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		authHeader := md["authorization"]
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "no authorization header")
		}

		tokenStr, err := ExtractToken(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "no token")
		}

		// 3) parse token
		claims, err := ParseAndVerify(tokenStr, rbac.PublicKey)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "error, "+err.Error())
		}

		// 4) role check if required
		allowedRoles := rbac.MethodRoles[info.FullMethod] // nil or empty => no role restriction
		if !containsAnyRole(claims.Roles, allowedRoles) {
			return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("required roles: %v", allowedRoles))
		}

		// 5) inject user info into context
		newCtx := context.WithValue(ctx, CtxUserID, claims.UserID)
		newCtx = context.WithValue(newCtx, CtxRoles, claims.Roles)

		// 6) call handler with newCtx
		return handler(newCtx, req)

	}
}
