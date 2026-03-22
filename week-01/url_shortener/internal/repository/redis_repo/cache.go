package redis_repo

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func NewCache() *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return &Cache{client: rdb}
}

func (c *Cache) Set(key string, value string) error {
	return c.client.Set(context.Background(), key, value, 0).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return c.client.Get(context.Background(), key).Result()
}