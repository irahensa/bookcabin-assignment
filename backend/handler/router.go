package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/api/check", CORSMiddleware(http.HandlerFunc(h.CheckVoucherHandler)))
	mux.Handle("/api/generate", CORSMiddleware(http.HandlerFunc(h.GenerateVoucherHandler)))
}
