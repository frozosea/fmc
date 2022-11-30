package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	const KEY = "testCacheKey"
	var cacheValue = []struct {
		Value string
	}{{Value: "TEST_VALUE"}}
	var redisUrl string
	redisUrl = os.Getenv("REDIS_URL")
	if redisUrl != "" {
		redisUrl = "localhost:6379"
	}
	cli := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	cache := NewCache(cli, time.Duration(1))
	if err := cache.Set(ctx, KEY, cacheValue); err != nil {
		log.Fatal(err.Error())
		return
	}
	var dest []struct {
		Value string
	}
	if err := cache.Get(ctx, KEY, &dest); err != nil {
		log.Fatal(err.Error())
	}
	for index, v := range dest {
		assert.Equal(t, v.Value, cacheValue[index].Value)
	}
}
