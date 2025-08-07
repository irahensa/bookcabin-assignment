package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/api/check", CORSMiddleware(http.HandlerFunc(h.CheckVoucherHandler)))
	mux.Handle("/api/generate", CORSMiddleware(http.HandlerFunc(h.GenerateVoucherHandler)))
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set common CORS headers for all responses
		w.Header().Set("Access-Control-Allow-Origin", "*") // Or specify allowed origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Add other required headers

		// Handle OPTIONS preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK) // Respond with 200 OK for preflight
			return
		}

		// Pass the request to the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
