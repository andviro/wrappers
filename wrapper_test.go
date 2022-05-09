package wrappers_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/andviro/wrappers"
)

type testA struct {
	x int
}

type testB struct {
	y string
}

func testWrapper[A, B any](dest io.Writer) wrappers.Wrapper[A, B] {
	return func(src wrappers.Func[A, B]) wrappers.Func[A, B] {
		return func(ctx context.Context, a A) (B, error) {
			fmt.Fprintf(dest, "calling with %#v\n", a)
			res, err := src(ctx, a)
			fmt.Fprintf(dest, "got: %#v, %+v", res, err)
			return res, err
		}
	}
}
func TestWrappers_Wrap(t *testing.T) {
	f := func(ctx context.Context, req *testA) (resp *testB, err error) {
		return &testB{
			y: fmt.Sprintf("(%d)", req.x),
		}, nil
	}
	buf := new(bytes.Buffer)
	f2 := wrappers.Wrap(f, testWrapper(buf))
	f2(context.TODO(), &testA{10})
	t.Logf("%s", buf.String())
}
