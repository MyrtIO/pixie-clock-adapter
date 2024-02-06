package dto

type LightsStateResponse struct {
	Color      []int `json:"color"`
	Brightness byte  `json:"brightness"`
	Enabled    bool  `json:"enabled"`
	Effect     byte  `json:"effect"`
}

type LightsStateRequest struct {
	Color      []byte `json:"color"`
	Brightness byte   `json:"brightness"`
	Effect     byte   `json:"effect"`
}
