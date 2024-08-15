package util

import "net/http"

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) error{
	return nil 
}

func ResponseWithError(w http.ResponseWriter, code int, message string ){
	
}