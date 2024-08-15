package server

import "net/http"

func (s *Server) Router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /api/v1/healthz", http.HandlerFunc(s.handleHealthz))
	mux.Handle("POST /api/v1/users", http.HandlerFunc(s.handleUserCreate))
	mux.Handle("GET /api/v1/users/", s.authMiddleware((s.handleGetUser)))

	return s.secureHeaders(mux)
}
