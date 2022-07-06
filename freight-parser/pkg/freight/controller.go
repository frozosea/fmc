package freight

import (
	"context"
	"fmc-newest/internal/cache"
	"fmc-newest/internal/logging"
	"fmt"
)

type IController interface {
	GetFreights(ctx context.Context, freight GetFreight) ([]BaseFreight, error)
}

type controllerUtils struct{}

func (c *controllerUtils) getCacheKeyByFreightRequest(freight *GetFreight) string {
	return fmt.Sprintf(`%d;%d;%s`, freight.FromCityId, freight.ToCityId, freight.ContainerType)
}

type controller struct {
	freightRepository IRepository
	cache             cache.ICache
	logger            logging.ILogger
	controllerUtils
}

func (s *controller) getFreightsFromRepo(ctx context.Context, freight GetFreight) ([]BaseFreight, error) {
	result, err := s.freightRepository.Get(ctx, freight)
	if err != nil {
		s.logger.FatalLog(fmt.Sprintf(`get information from database error: %s`, err.Error()))
	}
	go s.logger.InfoLog(fmt.Sprintf(`get info from database has result: %v`, result))
	return result, err
}

func (s *controller) getFreights(ctx context.Context, freight GetFreight) ([]BaseFreight, error) {
	cacheCh := make(chan []BaseFreight)
	go func() {
		var cacheFreights []BaseFreight
		chErr := s.cache.Get(ctx, s.getCacheKeyByFreightRequest(&freight), &cacheFreights)
		if chErr != nil {
			s.logger.FatalLog(fmt.Sprintf(`getFromCacheError: %s`, chErr.Error()))
		}
		cacheCh <- cacheFreights
	}()
	repoCh := make(chan []BaseFreight)
	go func() {
		result, err := s.getFreightsFromRepo(ctx, freight)
		if err != nil {
			s.logger.FatalLog(fmt.Sprintf(`getFreightsFromRepo err: %s`, err.Error()))
		}
		repoCh <- result
	}()
	select {
	case res := <-repoCh:
		return res, s.cache.Set(ctx, s.getCacheKeyByFreightRequest(&freight), res)
	case result := <-cacheCh:
		return result, nil
	}
}
func (s *controller) GetFreights(ctx context.Context, freight GetFreight) ([]BaseFreight, error) {
	return s.getFreights(ctx, freight)
}

func NewController(freightRepo IRepository, logger logging.ILogger) *controller {
	return &controller{freightRepository: freightRepo, logger: logger}
}
