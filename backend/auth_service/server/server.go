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
	// Menggunakan environment variable untuk fleksibilitas
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
		httpPort = "8081"
	}

	frontend := os.Getenv("FRONTEND")
	if frontend == "" {
		frontend = "http://localhost:5173"
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Menangani signal dari OS (Ctrl+C, SIGTERM)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	// gRPC Gateway
	// =========================================
	// Jalankan server gRPC dan gateway
	// gRPC Listener
	listener, err := net.Listen("tcp", ":"+gRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := RunGRPCServer() // gRPC di port 50052
	// HTTP Gateway
	// =========================================
	// Inisialisasi mux untuk HTTP Gateway
	// cookieCfg := utils.CookieCfg{
	// 	Domain: os.Getenv("COOKIE_DOMAIN"), // ".myapp.com" atau ""
	// 	Secure: os.Getenv("ENV") == "production",
	// }

	handlerHTTP := handler.NewHandlerHttp()
	loginHandler := handlerHTTP.HandlerLoginHTTP() // ambil handler
	refreshHandler := handlerHTTP.HandlerRefreshToken()
	logoutHandler := handlerHTTP.HandlerLogout()
	mux := runtime.NewServeMux()
	method, pattern := utils.CreatePattern("POST", "api", "v1", "auth", "web", "login")
	mux.Handle(method, pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		loginHandler(w, r)
	})
	method, pattern = utils.CreatePattern("POST", "api", "v1", "auth", "web", "logout")
	mux.Handle(method, pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		logoutHandler(w, r)
	})
	method, pattern = utils.CreatePattern("POST", "api", "v1", "auth", "web", "refresh")
	mux.Handle(method, pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		refreshHandler(w, r)
	})
	combinedHandler := middleware.Chain(
		mux,                      // handler paling dalam
		middleware.SecureHeaders, // ↓ urutan eksekusi
		middleware.Logging,
		middleware.RateLimit(5, 10),  // 5 req/detik, burst 10
		middleware.JWTAuthMiddleware, // cek token
		middleware.CORS(frontend),    // harus paling luar
	)
	grpcServerEndpoint := fmt.Sprintf("%s:%s", grpcHost, gRPCPort)
	RunHTTPGateway(ctx, mux, grpcServerEndpoint, httpPort) // HTTP gateway di port 8080
	// HTTP Listener
	httpListener, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Sync WaitGroup
	var wg sync.WaitGroup
	// wg.Add(2)
	wg.Add(2)

	// Menjalankan gRPC server dalam Goroutine
	go func() {
		defer wg.Done()
		log.Printf("gRPC server berjalan di :%s", gRPCPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Menjalankan HTTP Gateway dalam Goroutine
	go func() {
		defer wg.Done()
		log.Printf("HTTP gateway berjalan di :%s", httpPort)
		if err := http.Serve(httpListener, combinedHandler); err != nil {
			log.Fatalf("Failed to serve HTTP Gateway: %v", err)
		}
	}()

	// Menunggu sinyal shutdown
	wg.Add(1) // Tambahkan WaitGroup untuk shutdown goroutine
	go func() {
		defer wg.Done() // Pastikan WaitGroup diberi tahu setelah selesai
		<-signalChan
		fmt.Println("\nShutting down servers...")

		// Timeout shutdown dalam 5 detik
		_, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		// Matikan server gRPC
		grpcServer.GracefulStop()

		// Matikan HTTP Gateway
		if err := httpListener.Close(); err != nil {
			log.Printf("Error while closing HTTP listener: %v", err)
		}

		// Batalkan context utama agar semua operasi berhenti
		cancel()
	}()

	// Menunggu semua Goroutine selesai
	wg.Wait()
	fmt.Println("Server shutdown complete")
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
