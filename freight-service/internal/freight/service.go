package freight

import (
	"context"
	"encoding/json"
	"fmt"
	"freight_service/pkg/cache"
	"freight_service/pkg/logging"
)

type serviceUtils struct{}

func (c *serviceUtils) getCacheKeyByFreightRequest(freight *GetFreight) string {
	return fmt.Sprintf(`%d;%d;%d`, freight.FromCityId, freight.ToCityId, freight.ContainerTypeId)
}

type Service struct {
	repo   IRepository
	cache  cache.ICache
	logger logging.ILogger
	serviceUtils
}

func (c *Service) getFreights(ctx context.Context, freight GetFreight) ([]BaseFreight, error) {
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
func (c *Service) AddFreight(ctx context.Context, freight AddFreight) error {
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
func (c *Service) GetFreights(ctx context.Context, freight GetFreight) ([]BaseFreight, error) {
	return c.getFreights(ctx, freight)
}
func (c *Service) Update(ctx context.Context, id int, freight *AddFreight) error {
	go func() {
		j, err := json.Marshal(freight)
		if err != nil {
			return
		}
		c.logger.InfoLog(fmt.Sprintf(`freight with id %d was update: %v`, id, j))
	}()
	if err := c.repo.Update(ctx, id, freight); err != nil {
		return err
	}
	return c.cache.Del(ctx, c.getCacheKeyByFreightRequest(&GetFreight{
		FromCityId:      freight.FromCityId,
		ToCityId:        freight.ToCityId,
		ContainerTypeId: int64(freight.ContainerTypeId),
		Limit:           1,
	}))
}
func (c *Service) GetAll(ctx context.Context) ([]BaseFreight, error) {
	return c.repo.GetAll(ctx)
}
func (c *Service) Delete(ctx context.Context, id int) error {
	return c.repo.Delete(ctx, id)
}
func NewService(freightRepo IRepository, logger logging.ILogger, cache cache.ICache) *Service {
	return &Service{repo: freightRepo, logger: logger, cache: cache, serviceUtils: serviceUtils{}}
}
