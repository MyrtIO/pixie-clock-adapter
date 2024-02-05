package app

import (
	"encoding/json"
	"net/http"
	"pixie_adapter/internal/dto"
)

const deviceNotFoundMessage = "device is not connected"

func (s *Service) SetState(w http.ResponseWriter, r *http.Request) {
	var req dto.LightsStateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		jsonErrorResponse(w, 400, err.Error())
	}
	tx := s.provider.Get()
	if tx == nil {
		jsonErrorResponse(w, 503, deviceNotFoundMessage)
		return
	}
	err = s.usecase.Lights.SetState(tx, &req)
	if err != nil {
		jsonErrorResponse(w, 500, err.Error())
	}
	w.WriteHeader(204)
}

func (s *Service) DisableLights(w http.ResponseWriter, r *http.Request) {
	tx := s.provider.Get()
	if tx == nil {
		jsonErrorResponse(w, 503, deviceNotFoundMessage)
		return
	}
	err := s.usecase.Lights.TurnOff(tx)
	if err != nil {
		jsonErrorResponse(w, 500, err.Error())
	}
	w.WriteHeader(204)
}

func (s *Service) GetState(w http.ResponseWriter, r *http.Request) {
	tx := s.provider.Get()
	if tx == nil {
		jsonErrorResponse(w, 503, deviceNotFoundMessage)
		return
	}
	state, err := s.usecase.Lights.GetState(tx)
	if err != nil {
		jsonErrorResponse(w, 500, err.Error())
	}
	jsonResponse(w, 200, state)
}
