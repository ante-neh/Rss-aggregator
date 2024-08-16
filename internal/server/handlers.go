package server

import (
	"encoding/json"
	"github.com/ante-neh/Rss-aggregator/types"
	"github.com/ante-neh/Rss-aggregator/util"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request) {
	util.ResponseWithJson(w, 200, map[string]string{"message": "server is running "})
}

func (s *Server) handleUserCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	json.NewDecoder(r.Body).Decode(&params)

	userId, err := s.DB.Createuser(uuid.New(), time.Now().UTC(), time.Now().UTC(), params.Name)

	if err != nil {
		s.ErrorLogger.Println(err)
		util.ResponseWithError(w, 500, "Faild to create user")
		return
	}

	util.ResponseWithJson(w, 201, userId)

}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request, user types.User) {
	util.ResponseWithJson(w, 200, user)
}


func (s *Server) handleCreateFeeds(w http.ResponseWriter, r *http.Request, user types.User){
	type parameters struct{
		Name string
		Url string 
	}

	params := parameters{} 
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil{
		util.ResponseWithError(w, 400, "Bad Request")
		return 
	}
	
	feed, err := s.DB.CreateFeeds(uuid.New(), user.ID, params.Name, params.Url, time.Now().UTC(), time.Now().UTC())

	if err != nil{
		util.ResponseWithError(w, 404, "Unable to create feed")
		return 
	}

	util.ResponseWithJson(w, 201, feed)
}