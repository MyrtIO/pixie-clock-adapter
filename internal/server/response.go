package server

import (
	"encoding/json"
	"net/http"

	"log"
)

func jsonResponse(w http.ResponseWriter, code int, data interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error on formatting json response: %v\n", err)
	}
	w.WriteHeader(code)
	w.Write(j)
}
