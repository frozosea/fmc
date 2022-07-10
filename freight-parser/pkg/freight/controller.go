package freight

import (
	"context"
	"encoding/json"
	"fmc-newest/internal/cache"
	"fmc-newest/internal/logging"
	"fmt"
)

type controllerUtils struct{}

func (c *controllerUtils) getCacheKeyByFreightRequest(freight *GetFreight) string {
	return fmt.Sprintf(`%d;%d;%d`, freight.FromCityId, freight.ToCityId, freight.ContainerTypeId)
}

type controller struct {
	repo   IRepository
	cache  cache.ICache
	logger logging.ILogger
	controllerUtils
}

func (c *controller) getFreights(ctx context.Context, freight GetFreight) ([]BaseFreight, error) {
	cacheCh := make(chan []BaseFreight)
	go func() {
		var cacheFreights []BaseFreight
		chErr := c.cache.Get(ctx, c.getCacheKeyByFreightRequest(&freight), &cacheFreights)
		if chErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`getFromCacheError: %s`, chErr.Error()))
		}
		cacheCh <- cacheFreights
	}()
	repoCh := make(chan []BaseFreight)
	go func() {
		result, err := c.repo.Get(ctx, freight)
		if err != nil {
			go c.logger.ExceptionLog(fmt.Sprintf(`get information from database error: %s`, err.Error()))
		}
		go c.logger.InfoLog(fmt.Sprintf(`get info from database has result: %v`, result))
		repoCh <- result
	}()
	select {
	case res := <-repoCh:
		return res, c.cache.Set(ctx, c.getCacheKeyByFreightRequest(&freight), res)
	case result := <-cacheCh:
		return result, nil
	}
}
func (c *controller) AddFreight(ctx context.Context, freight AddFreight) error {
	if err := c.repo.Add(ctx, freight); err != nil {
		go func() {
			jsonRepr, marshErr := json.Marshal(freight)
			if marshErr != nil {
				c.logger.ExceptionLog(fmt.Sprintf(`add freight: %v failed: %s`, jsonRepr, err.Error()))
				return
			}
			c.logger.InfoLog(fmt.Sprintf(`freight: %v was add to repository`, jsonRepr))
		}()
		return err
	}
	return c.cache.Del(ctx, c.getCacheKeyByFreightRequest(&GetFreight{
		FromCityId:      freight.FromCityId,
		ToCityId:        freight.ToCityId,
		ContainerTypeId: int64(freight.ContainerTypeId),
		Limit:           1,
	}))
}
func (c *controller) GetFreights(ctx context.Context, freight GetFreight) ([]BaseFreight, error) {
	return c.getFreights(ctx, freight)
}

func NewController(freightRepo IRepository, logger logging.ILogger, cache cache.ICache) *controller {
	return &controller{repo: freightRepo, logger: logger, cache: cache, controllerUtils: controllerUtils{}}
}
