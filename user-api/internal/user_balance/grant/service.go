package grant

import (
	"context"
	"fmt"
	"user-api/pkg/logging"
)

type IService interface {
	AddStartGrant(ctx context.Context, userId int64) error
}

type Service struct {
	repo            IRepository
	logger          logging.ILogger
	grantStartValue float64
}

func NewService(repo IRepository, logger logging.ILogger, grantStartValue float64) *Service {
	return &Service{repo: repo, logger: logger, grantStartValue: grantStartValue}
}

func (s *Service) AddStartGrant(ctx context.Context, userId int64) error {
	if err := s.repo.Add(ctx, userId, s.grantStartValue); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`add start grant to user: %d err: %s`, userId, err.Error()))
		return err
	}
	go s.logger.InfoLog(fmt.Sprintf(`added start grant to user: %d value: %d`, userId, s.grantStartValue))
	return nil
}
