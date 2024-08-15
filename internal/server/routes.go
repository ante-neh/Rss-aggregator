package server

import "net/http"

func (s *Server) Router() *http.ServeMux{
	mux := http.NewServeMux()
	return mux 
}