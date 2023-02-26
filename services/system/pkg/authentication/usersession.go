package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"infra/services/system/pkg/cache"
	"time"
)

var (
	errUserSessionIsNil = errors.New(`usersession is nil`)
)

type UserSession struct {
	Uid     string   `redis:"uid"`
	Gid     string   `redis:"gid"`
	Roles   []string `redis:"roles"`
	AuthVer string   `redis:"authVer"`
}

func (u *UserSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u *UserSession) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, u)
}

type UserSessionKey string

func (u UserSession) Cache(ctx context.Context, key UserSessionKey) error {
	if _, err := cache.RedisClient.Set(ctx, string(key), u, time.Minute*30).Result(); err != nil {
		return err
	}
	return nil
}

func GetUserSession(ctx context.Context, authToken string) (*UserSession, error) {
	var res UserSession
	if err := cache.RedisClient.Get(ctx, string(cache.BuildUserSessionKey(authToken))).Scan(&res); err != nil {
		return nil, err
	}
	if res.Uid == "" {
		return nil, errUserSessionIsNil
	}
	return &res, nil
}
