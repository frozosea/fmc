package city

import (
	"context"
	"encoding/json"
	"fmc-newest/internal/cache"
	"fmc-newest/internal/logging"
	"fmt"
)

const CitiesRedisKey = "allCities"

type Controller struct {
	repo   *repository
	logger logging.ILogger
	cache  cache.ICache
}

func NewController(repo *repository, logger logging.ILogger, cache cache.ICache) *Controller {
	return &Controller{repo: repo, logger: logger, cache: cache}
}
func (c *Controller) AddCity(ctx context.Context, city BaseCity) error {
	go func() {
		jsonRepr, err := json.Marshal(city)
		if err != nil {
			return
		}
		c.logger.InfoLog(fmt.Sprintf(`city %v was add`, jsonRepr))
	}()
	if err := c.repo.Add(ctx, city); err != nil {
		return err
	}
	return c.cache.Del(ctx, CitiesRedisKey)
}
func (c *Controller) GetAll(ctx context.Context) ([]*City, error) {
	cacheCh := make(chan []*City, 1)
	go func() {
		var s []*City
		if err := c.cache.Get(ctx, CitiesRedisKey, &s); err != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`get all cities from cache failed: %s`, err.Error()))
			close(cacheCh)
			return
		}
		cacheCh <- s
	}()
	repoCh := make(chan []*City, 1)
	go func() {
		res, err := c.repo.GetAll(ctx)
		if err != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`get all city from repo failed: %s`, err.Error()))
			close(repoCh)
			return
		}
		repoCh <- res
	}()
	select {
	case <-ctx.Done():
		var cities []*City
		return cities, ctx.Err()
	case result := <-cacheCh:
		return result, nil
	case repoResult := <-repoCh:
		return repoResult, c.cache.Set(ctx, CitiesRedisKey, repoResult)
	}
}
