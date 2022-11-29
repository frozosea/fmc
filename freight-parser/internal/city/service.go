package city

import (
	"context"
	"encoding/json"
	"fmt"
	"freight_service/pkg/cache"
	"freight_service/pkg/logging"
)

const (
	CitiesRedisKey    = "allCities"
	CountriesCacheKey = "allCountries"
)

type Service struct {
	repo   IRepository
	logger logging.ILogger
	cache  cache.ICache
}

func NewService(repo *Repository, logger logging.ILogger, cache cache.ICache) *Service {
	return &Service{repo: repo, logger: logger, cache: cache}
}
func (c *Service) AddCity(ctx context.Context, city *CountryWithId) error {
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
func (c *Service) GetAllCities(ctx context.Context) ([]*City, error) {
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
		res, err := c.repo.GetAllCities(ctx)
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

func (c *Service) GetAllCountries(ctx context.Context) ([]*Country, error) {
	cacheCh := make(chan []*Country, 1)
	repoCh := make(chan []*Country, 1)
	errCh := make(chan error, 1)
	go func() {
		var res []*Country
		if err := c.cache.Get(ctx, CountriesCacheKey, &res); err != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`get all countries error: %s`, err.Error()))
			close(cacheCh)
			return
		}
		cacheCh <- res
	}()
	go func() {
		result, err := c.repo.GetAllCountries(ctx)
		if err != nil {
			close(repoCh)
			errCh <- err
			return
		}
		repoCh <- result
	}()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-repoCh:
			return res, nil
		case err := <-errCh:
			return nil, err
		case res := <-cacheCh:
			return res, nil
		default:
			continue
		}
	}
}

func (c *Service) UpdateCity(ctx context.Context, id int, city *CountryWithId) error {
	if err := c.repo.UpdateCity(ctx, id, city); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`update city with id %d error: %s`, id, err.Error()))
		return err
	}
	go func() {
		j, err := json.Marshal(city)
		if err != nil {
			return
		}
		c.logger.InfoLog(fmt.Sprintf(`update city with id %d : %v`, id, j))
	}()
	return nil
}

func (c *Service) UpdateCountry(ctx context.Context, id int, city *BaseEntity) error {
	if err := c.repo.UpdateCountry(ctx, id, city); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`update country with id %d error: %s`, id, err.Error()))
		return err
	}
	go func() {
		j, err := json.Marshal(city)
		if err != nil {
			return
		}
		c.logger.InfoLog(fmt.Sprintf(`update country with id %d : %v`, id, j))
	}()
	return nil
}
func (c *Service) AddCountry(ctx context.Context, country BaseEntity) error {
	return c.repo.AddCountry(ctx, country)
}
func (c *Service) DeleteCountry(ctx context.Context, id int) error {
	return c.repo.DeleteCountry(ctx, id)
}

func (c *Service) DeleteCity(ctx context.Context, id int) error {
	return c.repo.DeleteCity(ctx, id)
}
