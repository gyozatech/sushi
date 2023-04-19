package utils

// returns a pointer to a value
func PointerTo[T any](t T) *T {
	return &t
}
