package functional

/*
import (
	"fmt"
)

func main() {

	opt := NewOptional(5)
	fmt.Printf("%+v\n", opt)
	opt.Map(func(i int) (*string, error) {})

}

type Fn func(t interface{}) (interface{}, error)

var f1 Fn = func(t int) (string, error) {
	return fmt.Sprintf("%v", t), nil
}

var f2 Fn = func(t string) (string, error) {
	return t + "-1234", nil
}

func (o *Optional) Map(f Fn)



type Optional[T any, V any] struct {
	value Either[T]
	fn    Function[T, V]
}

func NewOptional[T any](value T) *Optional[T, error] {
	return &Optional[T, error]{
		value: EitherFromResult(&value),
		fn:    nil,
	}
}

func (opt *Optional[T, V]) Map(fn Function[T, V]) *Optional[V, error] {
	if opt.fn == nil {
		return &Optional[V, error]{
			value: EitherFromError[V](fmt.Errorf("function is nil")),
			fn:    nil,
		}
	}

	if opt.value.IsResult() {
		res, err := fn(*opt.value.result)
		if err != nil && res != nil {
			if res == nil {
				return &Optional[V, error]{
					value: EitherFromError[V](fmt.Errorf("value is nil")),
					fn:    nil,
				}
			}
			return NewOptional(*res)
		} else {
			return &Optional[V, error]{
				value: EitherFromError[V](err),
				fn:    nil,
			}
		}
	}
	return &Optional[V, error]{
		value: EitherFromError[V](opt.value.err),
		fn:    nil,
	}
}
*/
