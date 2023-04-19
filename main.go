package main

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
)

// PointerOf returns a pointer to a value
func PointerTo[T any](t T) *T {
	return &t
}

// ~~~~~~~~~~ EITHER ~~~~~~~~~~~~ //

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

// ~~~~~~~~~~~~~~~~~~~~~~

// Function type is the generic func which gets a T type and returns a V type and an error
type Function[T any, V any] func(t T) (*V, error)

// usage
//
//var int2str Function[int, string] = func(t int) (*string, error) {
//    return &fmt.Sprintf("%v",t), nil
//}

// ~~~~~~~~~~~~~~~~~~~~`

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

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`
