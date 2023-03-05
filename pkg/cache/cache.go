package cache

import "context"

type Regsitry[T any] interface {
	BuildKey(ctx context.Context) (string, error)
	Cache(ctx context.Context, value T) error
	Update(ctx context.Context, value T) error
	Get(ctx context.Context) (T, error)
}
