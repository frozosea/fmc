package user

import (
	"context"
	"fmt"
	"user-api/internal/domain"
	"user-api/pkg/cache"
	"user-api/pkg/logging"
)

type Provider struct {
	repository IRepository
	logger     logging.ILogger
	cache      cache.ICache
}

func NewProvider(repository IRepository, logger logging.ILogger, cache cache.ICache) *Provider {
	return &Provider{repository: repository, logger: logger, cache: cache}
}

func (p *Provider) AddContainerToAccount(ctx context.Context, userId int, containers []string) error {
	go p.logger.InfoLog(fmt.Sprintf(`add numbers to user: %d, numbers: %v`, userId, containers))
	if saveErr := p.repository.AddContainerToAccount(ctx, userId, containers); saveErr != nil {
		go p.logger.ExceptionLog(fmt.Sprintf(`add container to user-pb: %d failed with err: %s`, userId, saveErr.Error()))
		return saveErr
	}
	return p.cache.Del(ctx, fmt.Sprintf(`%d`, userId))
}
func (p *Provider) AddBillNumberToAccount(ctx context.Context, userId int, numbers []string) error {
	go p.logger.InfoLog(fmt.Sprintf(`add numbers to user: %d, numbers: %v`, &userId, &numbers))
	if addBillErr := p.repository.AddBillNumberToAccount(ctx, userId, numbers); addBillErr != nil {
		go p.logger.ExceptionLog(fmt.Sprintf(`add numbers: %v to user: %d err: %s`, numbers, userId, addBillErr.Error()))
		return addBillErr
	}
	return p.cache.Del(ctx, fmt.Sprintf(`%d`, userId))
}
func (p *Provider) DeleteContainersFromAccount(ctx context.Context, userId int, numberIds []int64) error {
	if deleteErr := p.repository.DeleteContainersFromAccount(ctx, userId, numberIds); deleteErr != nil {
		go p.logger.ExceptionLog(fmt.Sprintf(`delete containers: %v from user-pb: %d failed with err: %s`, numberIds, userId, deleteErr.Error()))
		return deleteErr
	}
	return p.cache.Del(ctx, fmt.Sprintf(`%d`, userId))
}
func (p *Provider) DeleteBillNumbersFromAccount(ctx context.Context, userId int, numberIds []int64) error {
	if delErr := p.repository.DeleteBillNumbersFromAccount(ctx, userId, numberIds); delErr != nil {
		//go p.logger.ExceptionLog(fmt.Sprintf(`delete bill numbers: %v from user-pb: %d failed with err: %s`, numberIds, userId, delErr.Error()))
		return delErr
	}
	return p.cache.Del(ctx, fmt.Sprintf(`%d`, userId))
}
func (p *Provider) GetAllContainers(ctx context.Context, userId int) (*domain.AllContainersAndBillNumbers, error) {
	cacheCh := make(chan *domain.AllContainersAndBillNumbers)
	go func() {
		var containers domain.AllContainersAndBillNumbers
		if getFromCacheError := p.cache.Get(ctx, fmt.Sprintf(`%d`, userId), &containers); getFromCacheError != nil {
			return
		}
		cacheCh <- &containers
	}()
	repoCh := make(chan *domain.AllContainersAndBillNumbers)
	go func() {
		result, getFromRepositoryErr := p.repository.GetAllContainersAndBillNumbers(ctx, userId)
		if getFromRepositoryErr != nil {
			return
		}
		repoCh <- result
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case cacheResult := <-cacheCh:
		return cacheResult, nil
	case repoRes := <-repoCh:
		return repoRes, p.cache.Set(ctx, fmt.Sprintf(`%d`, userId), repoRes)
	}
}
