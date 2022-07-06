package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

type ICache interface {
	Get(ctx context.Context, key string, dest interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
}

type cache struct {
	client *redis.Client
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
	return nil
}

func NewCache(redisCli *redis.Client) *cache {
	return &cache{client: redisCli}
}
