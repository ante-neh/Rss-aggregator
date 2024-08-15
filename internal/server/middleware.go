package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ante-neh/Rss-aggregator/types"
	"github.com/ante-neh/Rss-aggregator/util"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user types.User)

func(s *Server) authMiddleware(next authHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		val := r.Header.Get("Authorization")
		if val == ""{
			util.ResponseWithError(w, 400, "No Authentication info found" )
			return 
		}

		vals := strings.Split(val, " ")
		if len(vals) != 2{
			util.ResponseWithError(w, 400, "Malformed auth")
			return 
		}

		if vals[0] != "Apikey"{
			util.ResponseWithError(w, 400, "Malformed  first part of auth")
			return 
		}

		user, err := s.DB.GetUser(vals[1])

		if err != nil{
			util.ResponseWithError(w, 403, fmt.Sprintf("User Not Found: %v", err))
			return 
		}

		next(w, r, user)
	}
}