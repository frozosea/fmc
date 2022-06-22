package cache

import "context"

type ICache interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, ttl int64)
}
type cache struct {
}

func (c *cache) Get(ctx context.Context, key string) (interface{}, error) {
	return nil, nil
}
func (c *cache) Set(ctx context.Context, key string, value interface{}, ttl int64) {

}

func NewCache() *cache {
	return &cache{}
}
