package logic

import (
	"big-genius/internal/app/models/redis"
	"context"
	goredis "github.com/redis/go-redis/v9"
	"time"
)

func RedisLock(ctx context.Context, key string, value any, dur time.Duration) (bool, error) {
	return redis.RDB.SetNX(ctx, key, value, dur).Result()
}

func RedisGet(ctx context.Context, key string) (string, error) {
	result, err := redis.RDB.Get(ctx, key).Result()
	if err == goredis.Nil {
		return redis.KeyNotExist, nil
	} else if err != nil {
		return "", err
	}
	return result, err
}

func RedisSetNX(ctx context.Context, key string, value any, dur time.Duration) (bool, error) {
	return redis.RDB.SetNX(ctx, key, value, dur).Result()
}

func RedisDelete(ctx context.Context, key string) (int64, error) {
	return redis.RDB.Del(ctx, key).Result()
}
