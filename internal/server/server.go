package server

import (
	"database/sql"
	"github.com/ante-neh/Rss-aggregator/internal/database"
	"log"
	"net/http"
)

type Server struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	Address     string
	DB          *database.Postgres
}

func NewServer(infoLogger, errorLogger *log.Logger, address string, db *sql.DB) *Server {
	return &Server{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
		Address:     address,
		DB: &database.Postgres{
			DB: db,
		},
	}
}

func (s *Server) Start() *http.Server {
	return &http.Server{
		Addr:     s.Address,
		Handler:  s.Router(),
		ErrorLog: s.ErrorLogger,
	}
}
