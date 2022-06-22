package freight

import (
	"context"
	"fmc-newest/internal/logging"
	"fmc-newest/pkg/cache"
	"fmc-newest/pkg/domain"
	"fmt"
)

type IController interface {
	GetFreights(ctx context.Context, freight domain.GetFreight) ([]domain.BaseFreight, error)
}

type Controller struct {
	freightRepository   IFreightRepository
	cityRepository      ICityRepository
	contactRepository   IContactRepository
	containerRepository IContainerRepository
	cache               cache.ICache
	logger              logging.ILogger
}

func (s *Controller) getFreightsFromRepo(ctx context.Context, freight domain.GetFreight) ([]domain.BaseFreight, error) {
	result, err := s.freightRepository.Get(ctx, freight)
	if err != nil {
		fmt.Println(err.Error())
		s.logger.FatalLog(fmt.Sprintf(`get information from database error: %s`, err.Error()))
	}
	go s.logger.InfoLog(fmt.Sprintf(`get info from database has result: %v`, result))
	return result, err
}

func (s *Controller) getFreights(ctx context.Context, freight domain.GetFreight) ([]domain.BaseFreight, error) {
	return s.getFreightsFromRepo(ctx, freight)
}
func (s *Controller) GetFreights(ctx context.Context, freight domain.GetFreight) ([]domain.BaseFreight, error) {
	return s.getFreights(ctx, freight)
}

func NewController(freightRepo IFreightRepository, cityRepository ICityRepository, contactRepo IContactRepository, containerRepo IContainerRepository, logger logging.ILogger) *Controller {
	return &Controller{freightRepository: freightRepo, cityRepository: cityRepository, contactRepository: contactRepo, containerRepository: containerRepo, logger: logger}
}
