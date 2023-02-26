package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func initRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.174.129:6379",
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic("连接redis失败, error=" + err.Error())
	}
}

func init() {
	initRedis()
}
