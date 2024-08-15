package server

import "net/http"

func (s *Server) Router() *http.ServeMux{
	mux := http.NewServeMux()
	mux.Handle("GET /api/v1/healthz", http.HandlerFunc(s.handleHealthz))
	mux.Handle("POST /ap/v1/users", http.HandlerFunc(s.handleUserCreate))
	
	return mux 
}