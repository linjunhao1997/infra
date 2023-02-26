package authentication

import (
	"context"
	"fmt"
	"testing"
)

func Test_Name(t *testing.T) {
	a := &a{
		ctx: context.Background(),
	}
	ctx := a.ctx
	ctx = context.WithValue(ctx, "a", "a")
	//ctx = lctx
	fmt.Println(a.ctx.Value("a"))
}

type a struct {
	ctx context.Context
}
