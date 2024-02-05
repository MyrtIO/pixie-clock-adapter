package dto

type ClockColor []byte

func (c ClockColor) Bytes() []byte {
	bytes := make([]byte, len(c))
	for i := range []byte(c) {
		bytes[i] = byte(c[i])
	}
	return bytes
}

func (c ClockColor) Ints() []int {
	ints := make([]int, len(c))
	for i := range []byte(c) {
		ints[i] = int(c[i])
	}
	return ints
}

type ClockStateResponse struct {
	Color      []int `json:"color"`
	Brightness byte  `json:"brightness"`
	Enabled    bool  `json:"enabled"`
}

type ClockStateRequest struct {
	Color      []byte `json:"color"`
	Brightness byte   `json:"brightness"`
}
