package server

import "net/http"

func (s *Server) Router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /api/v1/healthz", http.HandlerFunc(s.handleHealthz))
	mux.Handle("POST /api/v1/users", http.HandlerFunc(s.handleUserCreate))
	mux.Handle("GET /api/v1/users/", s.authMiddleware((s.handleGetUser)))
	mux.Handle("POST /api/v1/feeds", s.authMiddleware(s.handleCreateFeeds))
	mux.Handle("GET /api/v1/feeds", http.HandlerFunc(s.handleGetFeeds))

	
	return s.secureHeaders(mux)
}
