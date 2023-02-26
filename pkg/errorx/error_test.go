package errorx

import (
	"errors"
	"fmt"
	"testing"
)

func TestWrapError(t *testing.T) {
	err := NewExpectError(123, "123")
	//err = err.WrapErr(NewExpectError("456", "456").WrapErr(errors.New("789")))
	fmt.Println(err.DeepestExpectError())
	var expect *ExpectError
	fmt.Println(errors.As(err, &expect))
}
