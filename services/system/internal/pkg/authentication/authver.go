package authentication

import (
	"context"
	"infra/services/system/internal/store"
)

func GetAuthVer(ctx context.Context, gid string) (string, error) {
	return store.RedisClient.Get(ctx, string(BuildAuthVerKey(gid))).Result()
}

func UpdateAuthVer(ctx context.Context, gid string) error {
	_, err := store.RedisClient.Incr(ctx, string(BuildAuthVerKey(gid))).Result()
	if err != nil {
		return err
	}
	return nil
}
