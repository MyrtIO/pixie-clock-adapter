package server

import (
	"encoding/json"
	"log"
	"net/http"
	"pixie_adapter/internal/dto"
	"pixie_adapter/internal/pixie"
)

// @title chi-swagger example APIs
// @version 1.0
// @description chi-swagger example APIs
// @BasePath /
func (s *Server) handleSetState(w http.ResponseWriter, r *http.Request) {
	var req dto.ClockStateRequest
	d := json.NewDecoder(r.Body)
	err := d.Decode(&req)
	if err != nil {
		log.Printf("Error on parsing request: %v\n", err)
		w.WriteHeader(400)
	}
	tx := s.connection.Get()
	if tx == nil {
		w.WriteHeader(503)
		return
	}
	log.Println(req)
	_, err = pixie.SetPower(tx, true)
	if err != nil {
		log.Printf("Error on enabling: %v\n", err)
		w.WriteHeader(400)
	}
	_, err = pixie.SetColor(tx, req.Color[0], req.Color[1], req.Color[2])
	if err != nil {
		log.Printf("Error on setting color: %v\n", err)
		w.WriteHeader(500)
		return
	}
	_, err = pixie.SetBrightness(tx, req.Brightness)
	if err != nil {
		log.Printf("Error on setting brightness: %v\n", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

// @title Disable
// @version 1.0
// @description chi-swagger example APIs
// @BasePath /
func (s *Server) handleDisable(w http.ResponseWriter, r *http.Request) {
	tx := s.connection.Get()
	if tx == nil {
		w.WriteHeader(503)
		return
	}
	_, err := pixie.SetPower(tx, false)
	if err != nil {
		log.Printf("Error on setting color: %v\n", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func (s *Server) handleGetState(w http.ResponseWriter, r *http.Request) {
	tx := s.connection.Get()
	if tx == nil {
		w.WriteHeader(503)
		return
	}
	colorValue, err := pixie.GetColor(tx)
	if err != nil {
		log.Printf("Error on reading color: %v\n", err)
		w.WriteHeader(500)
		return
	}
	brightness, err := pixie.GetBrightness(tx)
	if err != nil {
		log.Printf("Error on reading brightness: %v\n", err)
		w.WriteHeader(500)
		return
	}
	enabled, err := pixie.GetPower(tx)
	if err != nil {
		log.Printf("Error on reading power state: %v\n", err)
		w.WriteHeader(500)
		return
	}
	var color dto.ClockColor = colorValue
	jsonResponse(w, 200, dto.ClockStateResponse{
		Color:      color.Ints(),
		Brightness: brightness,
		Enabled:    enabled,
	})
}
