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

func NewController(repository IRepository, logger logging.ILogger, cache cache.ICache) *Controller {
	return &Controller{repository: repository, logger: logger, cache: cache}
}

func (c *Controller) AddContainerToAccount(ctx context.Context, userId int, containers []string) error {
	go func() {
		jsonRepr, err := json.Marshal(containers)
		if err != nil {
			return
		}
		c.logger.InfoLog(fmt.Sprintf(`add container to user-pb: %d, containers: %s`, userId, jsonRepr))

	}()
	if saveErr := c.repository.AddContainerToAccount(ctx, userId, containers); saveErr != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`add container to user-pb: %d failed with err: %s`, userId, saveErr.Error()))
		return saveErr
	}
	return nil
}
func (c *Controller) AddBillNumberToAccount(ctx context.Context, userId int, numbers []string) error {
	go func() {
		jsonRepr, err := json.Marshal(numbers)
		if err != nil {
			return
		}
		c.logger.InfoLog(fmt.Sprintf(`add numbers to user: %d, numbers: %s`, userId, jsonRepr))

	}()
	if addBillErr := c.repository.AddBillNumberToAccount(ctx, userId, numbers); addBillErr != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`add numbers: %v to user: %d err: %s`, numbers, userId, addBillErr.Error()))
		return addBillErr
	}
	return nil
}
func (c *Controller) DeleteContainersFromAccount(ctx context.Context, userId int, numberIds []int64) error {
	if deleteErr := c.repository.DeleteContainersFromAccount(ctx, userId, numberIds); deleteErr != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`delete containers: %v from user-pb: %d failed with err: %s`, numberIds, userId, deleteErr.Error()))
		return deleteErr
	}
	return nil
}
func (c *Controller) DeleteBillNumbersFromAccount(ctx context.Context, userId int, numberIds []int64) error {
	if delErr := c.repository.DeleteBillNumbersFromAccount(ctx, userId, numberIds); delErr != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`delete bill numbers: %v from user-pb: %d failed with err: %s`, numberIds, userId, delErr.Error()))
		return delErr
	}
	return nil
}
func (c *Controller) GetAllContainers(ctx context.Context, userId int) (*domain.AllContainersAndBillNumbers, error) {
	cacheCh := make(chan *domain.AllContainersAndBillNumbers)
	go func() {
		var containers *domain.AllContainersAndBillNumbers
		if getFromCacheError := c.cache.Get(ctx, fmt.Sprintf(`%d`, userId), &containers); getFromCacheError != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`get from cache failed: %s`, getFromCacheError.Error()))
		}
		cacheCh <- containers
	}()
	repoCh := make(chan *domain.AllContainersAndBillNumbers)
	go func() {
		result, getFromRepositoryErr := c.repository.GetAllContainersAndBillNumbers(ctx, userId)
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
		return repoRes, c.cache.Set(ctx, fmt.Sprintf(`%d`, userId), repoRes)
	}
}
