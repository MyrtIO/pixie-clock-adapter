package dto

type LightsStateResponse struct {
	Color      []int `json:"color"`
	Brightness byte  `json:"brightness"`
	Enabled    bool  `json:"enabled"`
}

type LightsStateRequest struct {
	Color      []byte `json:"color"`
	Brightness byte   `json:"brightness"`
}
