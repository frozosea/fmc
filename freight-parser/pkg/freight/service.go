package freight

import (
	"fmt"
	"sort"
	"sync"
)

type Service struct {
	repo   IRepository
	logger ILogger
}

func (s *Service) getFreights(freight GetFreight) ([]BaseFreight, error) {
	var wg sync.WaitGroup
	freightResultCh := make(chan []BaseFreight)
	wg.Add(1)
	go func() {
		defer wg.Done()
		result, err := s.repo.GetFrieght(freight)
		if err != nil {
			fmt.Println(err.Error())
			s.logger.FatalLog(fmt.Sprintf(`get information from database error: %s`, err.Error()))
		}
		go s.logger.InfoLog(fmt.Sprintf(`get info from database has result: %v`, result))
		freightResultCh <- result
	}()
	var freightResult = <-freightResultCh
	wg.Wait()
	return freightResult, nil
}
func (s *Service) GetBestFreights(freight GetFreight) ([]BaseFreight, error) {
	allFreights, err := s.getFreights(freight)
	if err != nil {
		s.logger.FatalLog(fmt.Sprintf(`get all freights by %v was exception: %s`, freight, err.Error()))
		return allFreights, err
	}
	sort.Slice(allFreights, func(i, j int) bool {
		return allFreights[i].UsdPrice < allFreights[j].UsdPrice
	})
	fmt.Println(allFreights)
	return allFreights, nil
}

func (s *Service) GetFreights(freight GetFreight) ([]BaseFreight, error) {
	return s.getFreights(freight)
}

func NewService(repo IRepository, logger ILogger) *Service {
	return &Service{repo: repo, logger: logger}
}
