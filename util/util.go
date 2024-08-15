package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}){
	response, err := json.Marshal(payload)
	if err != nil{
		log.Println("Faild to marshal payload: ", payload, err)
		w.WriteHeader(500)
		return 
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}

func ResponseWithError(w http.ResponseWriter, code int, message string ){
	if code > 499{
		log.Println("Server Stops with status 5XX", message )
	}

	type errorResponse struct{
		Error string `json:"error"`
	}

	ResponseWithJson(w, code, errorResponse{
		Error:message,
	})
	
}