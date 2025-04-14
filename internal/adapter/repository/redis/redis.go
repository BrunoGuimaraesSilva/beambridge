package redis

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	client *redis.Client
}

func New(addr string) *RedisRepo {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisRepo{client: client}
}

func (r *RedisRepo) UpdateProgress(ctx context.Context, uploadID string, percent int) error {
	return r.client.Set(ctx, "progress:"+uploadID, percent, 0).Err()
}

func (r *RedisRepo) GetProgress(ctx context.Context, uploadID string) (int, error) {
	val, err := r.client.Get(ctx, "progress:"+uploadID).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}
