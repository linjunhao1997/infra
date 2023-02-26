package utils

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type RedLockUtil struct {
	redSync *redsync.Redsync
}

func NewRedLockUtil(cli *redis.Client) *RedLockUtil {
	return &RedLockUtil{
		redSync: redsync.New(goredis.NewPool(cli)),
	}
}

func (util *RedLockUtil) Lock(ctx context.Context, key string, expire time.Duration) (*redsync.Mutex, error) {
	mutex := util.redSync.NewMutex(key, redsync.WithExpiry(expire))
	if err := mutex.LockContext(ctx); err != nil {
		return nil, err
	}
	return mutex, nil
}
