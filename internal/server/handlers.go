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
