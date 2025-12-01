package interceptor

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"sekolah/jwks"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var publicMethods = map[string]bool{
	"/sekolah.SekolahService/RegistrasiSekolah": true,
	"/sekolah.SekolahService/CreateSekolah":     true,
	// "/sekolah.SekolahService/RefreshToken":              true,
	// "/sekolah.SekolahTenantService/GetSekolahTenant": true,
	// "/health.Health/Check":                        true, // contoh service lain
}

func JWTInterceptor(authJWKSURL string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		// md, ok := metadata.FromIncomingContext(ctx)
		// if !ok {
		// 	return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		// }
		// // Ambil token dari header Authorization (prioritas tinggi)
		// authHeaders := md["authorization"]
		// var tokenStr string

		// if len(authHeaders) > 0 {
		// 	tokenStr = strings.TrimPrefix(authHeaders[0], "Bearer ")
		// 	if tokenStr == authHeaders[0] { // artinya tidak ada "Bearer "
		// 		tokenStr = "" // abaikan
		// 	}
		// }

		// // Jika tidak ada di header, coba dari cookie
		// if tokenStr == "" {
		// 	// Coba dari gRPC-Gateway
		// 	cookieHeaders := md["grpcgateway-cookie"]
		// 	if len(cookieHeaders) == 0 {
		// 		// fallback: raw HTTP cookie header (e.g., via envoy/ngnix)
		// 		cookieHeaders = md[":cookie"]
		// 	}

		// 	if len(cookieHeaders) > 0 {
		// 		// Parse cookie (format: "name1=value1; name2=value2")
		// 		cookies := strings.Split(cookieHeaders[0], ";")
		// 		for _, c := range cookies {
		// 			parts := strings.SplitN(strings.TrimSpace(c), "=", 2)
		// 			if len(parts) == 2 && parts[0] == "access_token" { // sesuaikan nama cookie!
		// 				tokenStr = parts[1]
		// 				break
		// 			}
		// 		}
		// 	}
		// }

		if publicMethods[info.FullMethod] {
			return handler(ctx, req)
		}
		// Extract token from context
		tokenStr, err := ExtractToken(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}

		if tokenStr == "" {
			return nil, status.Errorf(codes.Unauthenticated, "missing or empty token in Authorization header or access_token cookie")
		}

		// Fetch JWKS with caching
		jwksData, ok := jwks.JWKSCache.Get()
		if !ok {
			j, err := jwks.FetchJWKS(authJWKSURL)
			if err != nil {
				return nil, errors.New("cannot fetch JWKS")
			}
			jwks.JWKSCache.Set(j)
			jwksData = j
		}

		// Debug: decode JWT header manually
		parts := strings.Split(tokenStr, ".")
		if len(parts) < 2 {
			return nil, status.Error(codes.Unauthenticated, "malformed JWT")
		}
		headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid JWT header encoding: %v", err)
		}
		log.Printf("🔍 Decoded JWT header JSON: %s", string(headerBytes))

		// Parse JWT header for KID
		parser := jwt.NewParser()
		token, _, err := parser.ParseUnverified(tokenStr, jwt.MapClaims{})
		if err != nil {
			return nil, err
		}

		kidVal, ok := token.Header["kid"]
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing 'kid' in token header")
		}
		kid, ok := kidVal.(string)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "'kid' is not a string")
		}

		jwk, err := jwks.FindKey(jwksData, kid)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "JWK not found for kid %q: %v", kid, err)
		}

		rsaPub, err := jwks.JWKToRSA(jwk)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to convert JWK to RSA: %v", err)
		}

		verifiedToken, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			// Pastikan algoritma sesuai (contoh: RS256)
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return rsaPub, nil
		})
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "token verification failed: %v", err)
		}
		if !verifiedToken.Valid {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		return handler(ctx, req)
	}
}
