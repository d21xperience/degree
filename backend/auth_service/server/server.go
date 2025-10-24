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
	grpcHost := os.Getenv("GRPC_HOST")
	if grpcHost == "" {
		grpcHost = "localhost"
	}

	gRPCPort := os.Getenv("GRPC_PORT")
	if gRPCPort == "" {
		gRPCPort = "50051"
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8182"
	}

	frontend := os.Getenv("FRONTEND")
	if frontend == "" {
		frontend = "http://localhost:3000"
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Tangani sinyal OS
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// =========================================================
	// Jalankan gRPC server
	listener, err := net.Listen("tcp", ":"+gRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := RunGRPCServer()
	// =========================================================

	// Inisialisasi HTTP mux (gRPC-Gateway)
	mux := runtime.NewServeMux()

	// Daftarkan handler custom HTTP (bukan gRPC)
	handlerHTTP := handler.NewHandlerHttp()
	loginHandler := handlerHTTP.HandlerLoginHTTP()
	refreshHandler := handlerHTTP.HandlerRefreshToken()
	logoutHandler := handlerHTTP.HandlerLogout()
	authMeHandler := handlerHTTP.HandlerAuthMe()

	// === Manual route (non gRPC Gateway) ===
	method, pattern := utils.CreatePattern("POST", "api", "v1", "as", "auth", "web", "login")
	mux.Handle(method, pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		loginHandler(w, r)
	})

	method, pattern = utils.CreatePattern("POST", "api", "v1", "as", "auth", "web", "logout")
	mux.Handle(method, pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		logoutHandler(w, r)
	})

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

	httpListener, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// =========================================================
	// Jalankan server
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		log.Printf("🚀 gRPC server running on :%s", gRPCPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()
	// Tunggu sedikit agar gRPC server benar-benar siap
	time.Sleep(500 * time.Millisecond)
	// =========================================================
	// Daftarkan gRPC-Gateway (otomatis)
	grpcServerEndpoint := fmt.Sprintf("%s:%s", grpcHost, gRPCPort)
	RunHTTPGateway(ctx, mux, grpcServerEndpoint, httpPort)
	// =========================================================

	go func() {
		defer wg.Done()
		log.Printf("🌐 HTTP Gateway running on :%s", httpPort)
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

// func corsMiddleware(h http.Handler) http.Handler {
// 	frontend := os.Getenv("FRONTEND")
// 	if frontend == "" {
// 		frontend = "http://localhost:5173"
// 	}

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", frontend)
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 		w.Header().Set("Access-Control-Allow-Credentials", "true")

// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		h.ServeHTTP(w, r)
// 	})
// }

// // secureHeaders menambahkan header keamanan standar
// func SecureHeaders(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// HTTP security headers
// 		w.Header().Set("Content-Security-Policy", "default-src 'self'")
// 		w.Header().Set("X-Frame-Options", "DENY")
// 		w.Header().Set("X-Content-Type-Options", "nosniff")
// 		w.Header().Set("X-XSS-Protection", "1; mode=block")
// 		w.Header().Set("Referrer-Policy", "no-referrer")
// 		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

// 		next.ServeHTTP(w, r)
// 	})
// }
