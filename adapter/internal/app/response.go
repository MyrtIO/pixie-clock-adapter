package app

import (
	"encoding/json"
	"net/http"
	"pixie_adapter/internal/dto"

	"log"
)

func jsonResponse(w http.ResponseWriter, code int, data interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error on formatting json response: %v\n", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(j)
}

func jsonErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	j, err := json.Marshal(&dto.SystemResponse{
		Message: message,
		Code:    code,
	})
	if err != nil {
		log.Printf("Error on formatting json response: %v\n", err)
	} else {
		w.Write(j)
	}
}
