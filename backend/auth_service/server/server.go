package server

import (
	"auth_service/jwks"
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
	// Load private key
	priv, err := utils.LoadPrivateKey(os.Getenv("JWT_PRIVATE_PATH"))
	if err != nil {
		log.Fatal("failed load private key:", err)
	}
	pub, kid, err := jwks.ParseKeyFile(os.Getenv("JWT_PRIVATE_PATH"))
	if err != nil {
		panic(err)
	}
	// Jalankan gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := RunGRPCServer(priv)
	// =========================================================

	// Inisialisasi HTTP mux (gRPC-Gateway)
	gwmux := runtime.NewServeMux()
	method, pattern := createPattern("GET", ".well-known", "jwks.json")
	gwmux.Handle(method, pattern, jwks.JWKSHandler(pub, kid))

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
	grpcServerEndpoint := fmt.Sprintf("%s:%d", "localhost", gRPCPort)
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
