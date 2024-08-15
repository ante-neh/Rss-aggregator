package server

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/ante-neh/Rss-aggregator/util"
	"github.com/google/uuid"
)

func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request){
	util.ResponseWithJson(w, 200,map[string]string{"message":"server is running "})
}


func (s *Server) handleUserCreate(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
	}

	params := parameters{}
	json.NewDecoder(r.Body).Decode(&params)

	userId, err := s.DB.Createuser(uuid.New(), time.Now().UTC(), time.Now().UTC(), params.Name)

	if err != nil{
		s.ErrorLogger.Println(err)
		util.ResponseWithError(w, 500, "Faild to create user")
		return 
	}

	util.ResponseWithJson(w, 201, userId)

}