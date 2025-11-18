package server

import (
	"auth_service/handler"
	"auth_service/middleware"
	"auth_service/utils"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func StartGRPCServer() {
	utils.LoadEnvFiles()
	gRPCPort := utils.GetIntEnv("GRPC_PORT", 50051)
	httpPort := utils.GetIntEnv("HTTP_PORT", 8182)
	frontend := utils.GetEnv("FRONTEND", "http://localhost:3000")
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
	// Jalankan gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := RunGRPCServer(priv)
	// =========================================================

	// Inisialisasi HTTP mux (gRPC-Gateway)
	mux := runtime.NewServeMux()

	// Daftarkan handler custom HTTP (bukan gRPC)
	handlerHTTP := handler.NewHandlerHttp()
	loginHandler := handlerHTTP.HandlerLoginHTTP()
	refreshHandler := handlerHTTP.HandlerRefreshToken()
	// logoutHandler := handlerHTTP.HandlerLogout()
	authMeHandler := handlerHTTP.HandlerAuthMe()

	// === Manual route (non gRPC Gateway) ===
	method, pattern := utils.CreatePattern("POST", "api", "v1", "as", "auth", "web", "login")
	mux.Handle(method, pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		loginHandler(w, r)
	})

	// method, pattern = utils.CreatePattern("POST", "api", "v1", "as", "auth", "web", "logout")
	// mux.Handle(method, pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	// 	logoutHandler(w, r)
	// })

	method, pattern = utils.CreatePattern("POST", "api", "v1", "as", "auth", "web", "refresh")
	mux.Handle(method, pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		refreshHandler(w, r)
	})

	method, pattern = utils.CreatePattern("GET", "api", "v1", "as", "auth", "web", "me")
	mux.Handle(method, pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		authMeHandler(w, r)
	})

	// Middleware Chain
	combinedHandler := middleware.Chain(
		mux,
		middleware.SecureHeaders,
		middleware.Logging,
		middleware.RateLimit(5, 10),
		middleware.JWTAuthMiddleware, // hanya untuk /auth/*
		middleware.CORS(frontend),
	)

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
	RunHTTPGateway(ctx, mux, grpcServerEndpoint, fmt.Sprintf("%d", httpPort))
	// =========================================================
	go func() {
		defer wg.Done()
		log.Printf("🌐 HTTP Gateway running on :%d", httpPort)
		if err := http.Serve(httpListener, combinedHandler); err != nil {
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
