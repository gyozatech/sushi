package functional

type Task[T any] func() (*T, error)

// Future struct is a monad implementing a parallel task to be performed
type Promise[T any] struct {
	ch       chan Either[T]
	task     Task[T]
	output   *Either[T]
	executed bool
}

func NewPromise[T any](task Task[T]) *Promise[T] {
	return &Promise[T]{
		ch:       make(chan Either[T], 1),
		task:     task,
		output:   nil,
		executed: false,
	}
}

func ComputeAsync[T any](task Task[T]) *Promise[T] {
	return NewPromise(task).Compute()
}

// Process function performs the wrapped function in another Goroutine and returns a Future with the wrapped result
// It can also used as a "void" because the wrapped chan is pointed by a Pointer
func (promise *Promise[T]) Compute() *Promise[T] {
	if promise.executed {
		return promise
	}
	go channelifyCompute(promise.task, &promise.ch)
	promise.executed = true
	return promise
}

func channelifyCompute[T any](task Task[T], ch *chan Either[T]) {
	defer close(*ch)
	output, err := task()
	if err != nil {
		*ch <- EitherFromError[T](err)
		return
	}
	*ch <- EitherFromResult(output)
}

// WaitForResult waits and gets the result to the main Goroutine in the form of an Either object
func (promise *Promise[T]) WaitForResult() Either[T] {
	if promise.output != nil {
		return *promise.output
	}
	either := <-promise.ch
	promise.output = &either
	return either
}
