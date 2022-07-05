package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type ICache interface {
	Get(ctx context.Context, key string, dest interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
}

type cache struct {
	client *redis.Client
	ttl    time.Duration
}

func (c *cache) Get(ctx context.Context, key string, dest interface{}) error {
	res := c.client.Get(ctx, key)
	if res.Err() != nil {
		return res.Err()
	}
	bytes, getErr := res.Bytes()
	if getErr != nil {
		return getErr
	}
	return json.Unmarshal(bytes, &dest)
}
func (c *cache) Set(ctx context.Context, key string, value interface{}) error {
	jsonRepr, marshErr := json.Marshal(value)
	if marshErr != nil {
		return marshErr
	}
	return c.client.Set(ctx, key, jsonRepr, c.ttl).Err()
}

func NewCache(redisCli *redis.Client, ttl time.Duration) *cache {
	return &cache{client: redisCli, ttl: ttl}
}
