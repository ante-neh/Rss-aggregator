package server

import (
	"net/http"
	"github.com/ante-neh/Rss-aggregator/util"
)

func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request){
	util.ResponseWithJson(w, 200,map[string]string{"message":"server is running "})

}


func (s *Server) handleUserCreate(w http.ResponseWriter, r *http.Request){
	
}