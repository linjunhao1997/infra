package authentication

import (
	"context"
	"infra/services/system/pkg/cache"
)

func GetAuthVer(ctx context.Context, gid string) (string, error) {
	return cache.RedisClient.Get(ctx, string(cache.BuildAuthVerKey(gid))).Result()
}
