package archive

import (
	"context"
	"encoding/json"
	"fmt"
	"schedule-tracking/pkg/logging"
	"schedule-tracking/pkg/tracking"
)

type Service struct {
	logger     logging.ILogger
	repository IRepository
}

func NewService(logger logging.ILogger, repository IRepository) *Service {
	return &Service{logger: logger, repository: repository}
}

func (s *Service) GetAll(ctx context.Context, userId int) (*AllArchive, error) {
	return s.repository.GetAll(ctx, userId)
}
func (s *Service) AddByContainer(ctx context.Context, userId int, info *tracking.ContainerNumberResponse) error {
	if err := s.repository.AddByContainer(ctx, userId, info); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`add new container into archive user id: %d error: %s`, userId, err.Error()))
		return err
	}
	go func() {
		j, err := json.Marshal(info)
		if err != nil {
			return
		}
		s.logger.InfoLog(fmt.Sprintf(`add new container into archive user id: %d info: %v`, userId, j))
	}()
	return nil
}
func (s *Service) AddByBill(ctx context.Context, userId int, info *tracking.BillNumberResponse) error {
	if err := s.repository.AddByBill(ctx, userId, info); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`add new bill into archive user id: %d error: %s`, userId, err.Error()))
		return err
	}
	go func() {
		j, err := json.Marshal(info)
		if err != nil {
			return
		}
		s.logger.InfoLog(fmt.Sprintf(`add new bill into archive user id: %d info: %v`, userId, j))
	}()
	return nil
}
