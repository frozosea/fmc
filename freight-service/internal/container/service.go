package container

import (
	"context"
	"fmt"
	"freight_service/pkg/logging"
)

type Service struct {
	logger logging.ILogger
	repo   IRepository
}

func NewService(logger logging.ILogger, repo IRepository) *Service {
	return &Service{logger: logger, repo: repo}
}

func (s *Service) GetAll(ctx context.Context) ([]*Container, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) Add(ctx context.Context, containerTypeFullName string) error {
	if err := s.repo.Add(ctx, containerTypeFullName); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf("add new container with full name error: %s", err.Error()))
		return err
	}
	go s.logger.InfoLog(fmt.Sprintf("success add new container with full name: %s", containerTypeFullName))
	return nil
}
func (s *Service) Update(ctx context.Context, containerId int, newContainerTypeFullName string) error {
	if err := s.repo.Update(ctx, containerId, newContainerTypeFullName); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf("update container with id %d error: %s", containerId, err.Error()))
		return err
	}
	go s.logger.InfoLog(fmt.Sprintf("success update container with id: %d new container full name: %s", containerId, newContainerTypeFullName))
	return nil
}
func (s *Service) Delete(ctx context.Context, containerId int) error {
	if err := s.repo.Delete(ctx, containerId); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf("delete container with id %d error: %s", containerId, err.Error()))
		return err
	}
	go s.logger.InfoLog(fmt.Sprintf("success delete container with id: %d", containerId))
	return nil
}
