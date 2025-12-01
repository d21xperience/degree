package server

import (
	auth "auth_service/generated"
	"auth_service/http_handler"
	"auth_service/jwks"
	"auth_service/jwt"
	"auth_service/utils"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func StartGRPCServer() {
	utils.LoadEnvFiles()
	gRPCPort := utils.GetIntEnv("GRPC_PORT", 50051)
	httpPort := utils.GetIntEnv("HTTP_PORT", 8182)
	// frontend := utils.GetEnv("FRONTEND", "http://localhost:3000")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Tangani sinyal OS
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	// =========================================================

	jwtManager := jwt.NewManager()

	// Jalankan gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := RunGRPCServer(jwtManager)
	// =========================================================
	grpcServerEndpoint := fmt.Sprintf("%s:%d", "localhost", gRPCPort)
	// Inisialisasi HTTP mux (gRPC-Gateway)
	gwmux := runtime.NewServeMux(
		// 🔑 Ini kunci utamanya: konversi cookie → metadata gRPC
		runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			md := metadata.MD{}

			// Ambil access_token dari cookie
			if c, err := r.Cookie("access_token"); err == nil && c != nil {
				md.Set("authorization", "Bearer "+c.Value)
			}

			// Opsional: juga kirim refresh_token (misal untuk refresh otomatis di interceptor)
			if c, err := r.Cookie("refresh_token"); err == nil && c != nil {
				md.Set("x-refresh-token", c.Value)
			}

			return md
		}),
	)
	method, pattern := createPattern("GET", ".well-known", "jwks.json")
	gwmux.Handle(method, pattern, jwks.JWKSHandler(jwtManager))
	// 🔌 Buat gRPC client (untuk HTTP handler)
	conn, err := grpc.NewClient(grpcServerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial gRPC: %v", err)
	}
	defer conn.Close()

	// ✅ Override /login, /refresh, /logout dengan handler custom (untuk Set-Cookie)
	httpHandler := &http_handler.HTTPHandler{AuthClient: auth.NewAuthServiceClient(conn)}
	method, pattern = createPattern("POST", "api", "v1", "as", "auth", "web", "login")
	gwmux.Handle(method, pattern, httpHandler.HandlerLoginHTTP())
	method, pattern = createPattern("POST", "api", "v1", "as", "auth", "web", "refresh")
	gwmux.Handle(method, pattern, httpHandler.HandlerRefreshHTTP())
	method, pattern = createPattern("POST", "api", "v1", "as", "auth", "web", "logout")
	gwmux.Handle(method, pattern, httpHandler.HandlerLogoutHTTP())

	httpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", httpPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// =========================================================
	// Jalankan server
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		log.Printf("🚀 gRPC server running on :%d", gRPCPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()
	// Tunggu sedikit agar gRPC server benar-benar siap
	time.Sleep(500 * time.Millisecond)
	// =========================================================
	// Daftarkan gRPC-Gateway (otomatis)
	// grpcServerEndpoint := fmt.Sprintf("%s:%d", "localhost", gRPCPort)
	RunHTTPGateway(ctx, gwmux, grpcServerEndpoint, fmt.Sprintf("%d", httpPort))
	// =========================================================
	go func() {
		defer wg.Done()
		log.Printf("🌐 HTTP Gateway running on :%d", httpPort)
		if err := http.Serve(httpListener, gwmux); err != nil {
			log.Fatalf("Failed to serve HTTP Gateway: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-signalChan
		log.Println("🛑 Shutting down servers...")

		grpcServer.GracefulStop()
		httpListener.Close()
		cancel()
	}()

	wg.Wait()
	log.Println("✅ Server shutdown complete")
}

func CustomHeaderMatcher(header string) (string, bool) {
	switch strings.ToLower(header) {
	case "authorization":
		return header, true
	default:
		return header, false
	}
}

func createPattern(method string, pathSegments ...string) (string, runtime.Pattern) {
	pattern := runtime.MustPattern(
		runtime.NewPattern(1, generatePatternIndexes(len(pathSegments)), pathSegments, ""),
	)
	return method, pattern
}

// generatePatternIndexes membantu membuat pola angka yang sesuai dengan jumlah segment
func generatePatternIndexes(segmentCount int) []int {
	indexes := []int{}
	for i := 0; i < segmentCount; i++ {
		indexes = append(indexes, 2, i)
	}
	return indexes
}
