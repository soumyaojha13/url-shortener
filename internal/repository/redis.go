package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

func (r *RedisRepo) Set(key, value string) {
	r.Client.Set(context.Background(), key, value, 0)
}

func (r *RedisRepo) Get(key string) (string, error) {
	return r.Client.Get(context.Background(), key).Result()
}
