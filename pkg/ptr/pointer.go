package ptr

func From[T any](value T) *T {
	return &value
}

func ToValue[T any](ptr *T) T {
	var dummy T
	if ptr == nil {
		return dummy
	}
	return *ptr
}
