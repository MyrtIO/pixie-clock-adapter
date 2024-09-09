package pointer

// From creates a pointer from a value
func From[T any](value T) *T {
	return &value
}

// ToValue creates a value from a pointer.
// If the pointer is nil, it returns a default value.
func ToValue[T any](ptr *T) T {
	var dummy T
	if ptr == nil {
		return dummy
	}
	return *ptr
}
