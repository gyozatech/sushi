package functional

// Future struct is a monad implementing a parallel task to be performed
type Future[T any, V any] struct {
	ch       chan Either[V]
	fn       Function[T, V]
	input    T
	output   *Either[V]
	executed bool
}

func ProcessAsync[T any, V any](fn Function[T, V], input T) *Future[T, V] {
	return NewFuture(fn, input).Process()
}

func NewFuture[T any, V any](fn Function[T, V], input T) *Future[T, V] {
	return &Future[T, V]{
		ch:       make(chan Either[V], 1),
		fn:       fn,
		input:    input,
		output:   nil,
		executed: false,
	}
}

// Process function performs the wrapped function in another Goroutine and returns a Future with the wrapped result
// It can also used as a "void" because the wrapped chan is pointed by a Pointer
func (future *Future[T, V]) Process() *Future[T, V] {
	if future.executed {
		return future
	}
	go channelifyProcess(future.fn, future.input, &future.ch)
	future.executed = true
	return future
}

func channelifyProcess[T any, V any](fn Function[T, V], input T, ch *chan Either[V]) {
	defer close(*ch)
	output, err := fn(input)
	if err != nil {
		*ch <- EitherFromError[V](err)
		return
	}
	*ch <- EitherFromResult(output)
}

// WaitForResult waits and gets the result to the main Goroutine in the form of an Either object
func (future *Future[T, V]) WaitForResult() Either[V] {
	if future.output != nil {
		return *future.output
	}
	either := <-future.ch
	future.output = &either
	return either
}
