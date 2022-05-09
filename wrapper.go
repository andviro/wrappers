package wrappers

import "context"

type Func[A, B any] func(ctx context.Context, req A) (resp B, err error)

type Wrapper[A, B any] func(src Func[A, B]) Func[A, B]

func Wrap[A, B any](src Func[A, B], wrappers ...Wrapper[A, B]) Func[A, B] {
	for i := len(wrappers) - 1; i >= 0; i-- {
		src = wrappers[i](src)
	}
	return src
}
