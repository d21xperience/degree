package middleware

// import (
// 	"net"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"golang.org/x/time/rate"
// )

// func RateLimit(rps int, burst int) func(http.Handler) http.Handler {
// 	limiterStore := struct {
// 		sync.Mutex
// 		clients map[string]*rate.Limiter
// 	}{clients: make(map[string]*rate.Limiter)}

// 	getLimiter := func(ip string) *rate.Limiter {
// 		limiterStore.Lock()
// 		defer limiterStore.Unlock()
// 		if l, exists := limiterStore.clients[ip]; exists {
// 			return l
// 		}
// 		l := rate.NewLimiter(rate.Every(time.Second/time.Duration(rps)), burst)
// 		limiterStore.clients[ip] = l
// 		return l
// 	}

// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
// 			if !getLimiter(ip).Allow() {
// 				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
// 				return
// 			}
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }
