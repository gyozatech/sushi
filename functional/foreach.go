package functional

// ForEach allows to apply a transformation function over a slice
func ForEach[T any, V any](slice []T, f Function[T, V]) ([]V, error) {
	newSlice := []V{}
	for _, t := range slice {
		if v, err := f(t); err != nil {
			return nil, err
		} else {
			newSlice = append(newSlice, *v)
		}
	}
	return newSlice, nil
}
