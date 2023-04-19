package functional

// Function type is the generic func which gets a T type and returns a V type and an error
type Function[T any, V any] func(t T) (*V, error)

// usage
//
//var int2str Function[int, string] = func(t int) (*string, error) {
//    return &fmt.Sprintf("%v",t), nil
//}
