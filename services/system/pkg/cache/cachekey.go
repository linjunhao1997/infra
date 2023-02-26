package cache

import (
	"context"
	"fmt"
	"time"
)

type AuthVerKey string

func (authVerKey AuthVerKey) Cache(ctx context.Context, gid string) error {
	key := string(BuildAuthVerKey(gid))
	ver, err := RedisClient.Incr(ctx, key).Result()
	if err != nil {
		return err
	}
	if _, err := RedisClient.Set(ctx, key, ver, time.Minute*30).Result(); err != nil {
		return err
	}
	return nil
}

func (authVerKey AuthVerKey) GetVer(ctx context.Context, gid string) (string, error) {
	ver, err := RedisClient.Get(ctx, string(BuildAuthVerKey(gid))).Result()
	if err != nil {
		return "", err
	}
	return ver, nil
}

type UserSessionKey string

func BuildUserSessionKey(authToken string) UserSessionKey {
	return (UserSessionKey)(fmt.Sprintf(`system:userSession:%s`, authToken))
}

func BuildAuthVerKey(gid string) AuthVerKey {
	return (AuthVerKey)(fmt.Sprintf(`system:auth:%s:ver`, gid))
}
