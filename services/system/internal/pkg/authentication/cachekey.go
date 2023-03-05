package authentication

import (
	"context"
	"fmt"
	"infra/pkg/cache"
	"infra/services/system/internal/store"
	"time"
)

type userSessionCacheRegistry struct {
	authToken string
	key       string
	Value     UserSession
}

func NewUserSessionCacheRegistry(authToken string) *userSessionCacheRegistry {
	return &userSessionCacheRegistry{
		authToken: authToken,
	}
}

func (cache *userSessionCacheRegistry) BuildKey(ctx context.Context) (string, error) {
	if cache.key == "" {
		cache.key = fmt.Sprintf(`system:userSession:%s`, cache.authToken)
	}
	return cache.key, nil
}

func (cache *userSessionCacheRegistry) Cache(ctx context.Context, value UserSession) error {
	_, err := cache.BuildKey(ctx)
	if err != nil {
		return err
	}
	if _, err := store.RedisClient.Set(ctx, cache.authToken, value, time.Minute*30).Result(); err != nil {
		return err
	}
	return nil
}

func (cache *userSessionCacheRegistry) Get(ctx context.Context) (userSession UserSession, err error) {
	if cache.Value.Uid != "" {
		return cache.Value, nil
	}
	if _, err = cache.BuildKey(ctx); err != nil {
		return
	}
	var res UserSession
	if err = store.RedisClient.Get(ctx, cache.key).Scan(&res); err != nil {
		return
	}
	if res.Uid == "" {
		return UserSession{}, errUserSessionIsNil
	}
	cache.Value = res
	return cache.Value, nil
}

func (cache *userSessionCacheRegistry) Update(ctx context.Context, _ UserSession) error {
	if _, err := cache.BuildKey(ctx); err != nil {
		return err
	}
	if _, err := store.RedisClient.Del(ctx, cache.key).Result(); err != nil {
		return err
	}
	return nil
}

var _ cache.Regsitry[UserSession] = (*userSessionCacheRegistry)(nil)

type authVerCacheRegistry struct {
	gid   string
	Value string
}

func NewAuthVerCacheRegistry(gid string) *authVerCacheRegistry {
	return &authVerCacheRegistry{
		gid: gid,
	}
}

func (ver *authVerCacheRegistry) BuildKey(ctx context.Context) (string, error) {
	key := fmt.Sprintf(`system:auth:%s:ver`, ver.gid)
	return key, nil
}

func (ver *authVerCacheRegistry) Cache(ctx context.Context, _ string) error {
	key, err := ver.BuildKey(ctx)
	if err != nil {
		return err
	}
	if _, err := store.RedisClient.Incr(ctx, key).Result(); err != nil {
		return err
	}
	return nil
}

func (ver *authVerCacheRegistry) Get(ctx context.Context) (string, error) {
	if ver.Value != "" {
		return ver.Value, nil
	}
	key, err := ver.BuildKey(ctx)
	if err != nil {
		return "", err
	}
	res, err := store.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	ver.Value = res
	return ver.Value, nil
}

func (ver *authVerCacheRegistry) Update(ctx context.Context, _ string) error {
	return ver.Cache(ctx, "")
}

var _ cache.Regsitry[string] = (*authVerCacheRegistry)(nil)
