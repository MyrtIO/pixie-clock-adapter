package usecase

func bytesToInts(b []byte) []int {
	ints := make([]int, 0, len(b))
	for _, v := range b {
		ints = append(ints, int(v))
	}
	return ints
}
