package middleware

import "net/http"

// Chain menumpuk banyak middleware: paling kiri dieksekusi paling luar.
func Chain(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
