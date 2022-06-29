package user

import (
	"context"
	"encoding/json"
	"fmt"
	"user-api/internal/cache"
	"user-api/internal/logging"
	"user-api/pkg/domain"
)

type Controller struct {
	repository IRepository
	logger     logging.ILogger
	cache      cache.ICache
}

func (c *Controller) AddContainerToAccount(ctx context.Context, userId int, containers []domain.Container) {
	jsonRepr, err := json.Marshal(containers)
	if err != nil {
		c.logger.InfoLog(fmt.Sprintf(`add container to user: %d, containers: %s`, userId, jsonRepr))
	}
	go func() {
		if saveErr := c.repository.AddContainerToAccount(ctx, userId, containers); saveErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`add container to user: %d failed with err: %s`, userId, saveErr.Error()))
			return
		}
	}()
	return
}
func (c *Controller) DeleteContainersFromAccount(ctx context.Context, userId int, containers []domain.Container) {
	jsonRepr, err := json.Marshal(containers)
	if err != nil {
		c.logger.InfoLog(fmt.Sprintf(`delete containers to user: %d, containers: %s`, userId, jsonRepr))
	}
	go func() {
		if deleteErr := c.repository.DeleteContainersFromAccount(ctx, userId, containers); deleteErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`delete containers: %s from user: %d failed with err: %s`, jsonRepr, userId, deleteErr.Error()))
		}
	}()
	return
}

//	GetAllContainers(ctx context.Context, userId int) ([]domain.Container, error)
func (c *Controller) GetAllContainers(ctx context.Context, userId int) ([]*domain.Container, error) {
	cacheCh := make(chan []*domain.Container)
	go func() {
		var containers []*domain.Container
		if getFromCacheError := c.cache.Get(ctx, fmt.Sprintf(`%d`, userId), &containers); getFromCacheError != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`get from cache failed: %s`, getFromCacheError.Error()))
		}
		cacheCh <- containers
	}()
	repoCh := make(chan []*domain.Container)
	go func() {
		result, getFromRepositoryErr := c.repository.GetAllContainers(ctx, userId)
		if getFromRepositoryErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`get from repo failed: %s`, getFromRepositoryErr.Error()))
		}
		repoCh <- result
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case cacheResult := <-cacheCh:
		return cacheResult, nil
	case repoRes := <-repoCh:
		jsonRepr, _ := json.Marshal(repoRes)
		return repoRes, c.cache.Set(ctx, fmt.Sprintf(`%d`, userId), jsonRepr)
	}
}
