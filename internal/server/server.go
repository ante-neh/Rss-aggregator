package server

import (
	"log"
	"net/http"
)


type Server struct{
	InfoLogger *log.Logger
	ErrorLogger *log.Logger 
	Address string 
}


func NewServer(infoLogger, errorLogger *log.Logger, address string) *Server{
	return &Server{
		InfoLogger: infoLogger,
		ErrorLogger: errorLogger,
		Address: address,

	}
}


func (s *Server) Start()*http.Server{
	return &http.Server{
		Addr:s.Address,
		Handler:s.Router(),
		ErrorLog:s.ErrorLogger,
	}
}