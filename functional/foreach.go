package functional

// ForEach allows to apply a transformation function over a slice with the following rules:
// - if one error occurs while processing one item, the whole ForEach function fails;
// - if processing one item returns a nil element, this will be discarded by the resulting slice
func ForEach[T any, V any](slice []T, f Function[T, V]) ([]V, error) {
	newSlice := []V{}
	for _, t := range slice {
		if v, err := f(t); err != nil {
			return nil, err
		} else if v != nil {
			newSlice = append(newSlice, *v)
		}
	}
	return newSlice, nil
}
