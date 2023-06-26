package logic

import (
	ctx "big-genius/core/context"
	"big-genius/internal/app/models/redis"
	goredis "github.com/redis/go-redis/v9"
	"time"
)

var KeyNotExist = "#KeyNotExist#"

func RedisLock(ctx ctx.Context, key string, value any, dur time.Duration) (bool, error) {
	return redis.RDB.SetNX(ctx.Context, key, value, dur).Result()
}

func RedisGet(ctx ctx.Context, key string) (string, error) {
	result, err := redis.RDB.Get(ctx.Context, key).Result()
	if err == goredis.Nil {
		return KeyNotExist, nil
	} else if err != nil {
		return "", err
	}
	return result, err
}
