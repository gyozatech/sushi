package functional

// Either struct is the equivalent of functional Scala Either: a wrapper for the result/error response
type Either[V any] struct {
	result *V
	err    error
}

// EitherFromResult static function initializes an Either object with a result
func EitherFromResult[V any](result *V) Either[V] {
	return Either[V]{result, nil}
}

// EitherFromError static function initializes an Either object with a
func EitherFromError[V any](err error) Either[V] {
	return Either[V]{nil, err}
}

// IsError function returns true if Either wraps an error object
func (e Either[V]) IsError() bool {
	return e.err != nil
}

// IsResult function returns true if Either doesn't wrap an error object
func (e Either[V]) IsResult() bool {
	return e.err == nil
}

// GetResult function returns the wrapped Result
func (e Either[V]) GetResult() *V {
	if e.err != nil {
		panic("Either struct contains an error")
	}
	return e.result
}

// GetError function returns the wrapped error if any
func (e Either[V]) GetError() error {
	return e.err
}

// Get function return the wrapped result which can "either" be an error or another object
func (e Either[V]) Get() (*V, error) {
	return e.result, e.err
}

// GetOrElse function returns the wrapped result or if an error is present, a default value
func (e Either[V]) GetOrElse(fallback V) *V {
	if e.IsError() {
		return &fallback
	}
	return e.result
}
