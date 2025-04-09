package functional

import (
	"fmt"
	"testing"
	"time"
)

func TestPromise(t *testing.T) {
	var input int = 5

	promise1 := NewPromise(func() (*string, error) {
		time.Sleep(5 * time.Second)
		output := fmt.Sprintf("%d!", input)
		return &output, nil
	})

	promise2 := NewPromise(func() (*string, error) {
		time.Sleep(5 * time.Second)
		output := fmt.Sprintf("%d!!", input)
		return &output, nil
	})
	promise2.Compute()

	promise3 := NewPromise(func() (*string, error) {
		time.Sleep(5 * time.Second)
		output := fmt.Sprintf("%d!!!", input)
		return &output, nil
	})

	promise1.Compute()
	promise2.Compute()
	promise3.Compute()

	either1 := promise1.WaitForResult()
	either2 := promise2.WaitForResult()
	either3 := promise3.WaitForResult()

	res1 := either1.GetResult()
	if *res1 != "5!" {
		t.Errorf("wrong return for promise1")
	}
	res2 := either2.GetResult()
	if *res2 != "5!!" {
		t.Errorf("wrong return for promise1")
	}
	res3 := either3.GetResult()
	if *res3 != "5!!!" {
		t.Errorf("wrong return for promise1")
	}
}
