package functional

import (
	"fmt"
	"testing"
)

func TestEitherResult(t *testing.T) {
	either := EitherFromResult(&[]string{"test-result"}[0])

	if either.IsError() {
		t.Errorf("Either object must be no error")
	}

	if !either.IsResult() {
		t.Errorf("Either object must be a result")
	}

	if *either.GetResult() != "test-result" {
		t.Errorf("Either object must have test-result as result")
	}

	if either.GetError() != nil {
		t.Errorf("Either object must wrap a nil error")
	}

	if *either.GetOrElse("alternative") == "alternative" {
		t.Errorf("Either object must not return a default result if result is not nil")
	}

}

func TestEitherError(t *testing.T) {
	either := EitherFromError[string](fmt.Errorf("test-error"))

	if !either.IsError() {
		t.Errorf("Either object must be error")
	}

	if either.IsResult() {
		t.Errorf("Either object must not be a result")
	}

	if either.GetError() == nil {
		t.Errorf("Either object must wrap an error")
	}

	if *either.GetOrElse("alternative") != "alternative" {
		t.Errorf("Either object must return a default result if result is not nil")
	}

}

func TestEitherErrorWhenPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	either := EitherFromError[string](fmt.Errorf("test-error"))
	either.GetResult()

}
