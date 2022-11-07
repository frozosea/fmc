package company

import (
	"context"
	"encoding/json"
	"fmt"
	"freight_service/pkg/cache"
	"freight_service/pkg/logging"
)

const cacheKey = "allContacts"

type Service struct {
	repo   IRepository
	cache  cache.ICache
	logger logging.ILogger
}

func (c *Service) GetAllContacts(ctx context.Context) ([]*Company, error) {
	cacheCh := make(chan []*Company)
	go func() {
		var result []*Company
		if cacheErr := c.cache.Get(ctx, cacheKey, &result); cacheErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`get all contacts from cache error: %s`, cacheErr.Error()))
		}
		cacheCh <- result
	}()
	repoCh := make(chan []*Company)
	go func() {
		repoRes, repoErr := c.repo.GetAll(ctx)
		if repoErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`get all contacts from Repository error: %s`, repoErr.Error()))
		}
		repoCh <- repoRes
	}()
	for {
		select {
		case result := <-cacheCh:
			return result, nil
		case repoResult := <-repoCh:
			jsonRepresent, jsonMarshErr := json.Marshal(repoResult)
			if jsonMarshErr != nil {
				return repoResult, nil
			}
			return repoResult, c.cache.Set(ctx, cacheKey, jsonRepresent)
		default:
			continue
		}
	}
}
func (c *Service) AddContact(ctx context.Context, contact BaseCompany) error {
	addErr := c.repo.Add(ctx, contact)
	marshResp, marshErr := json.Marshal(contact)
	if marshErr != nil {
		return marshErr
	}
	c.logger.InfoLog(fmt.Sprintf(`add new company: %s`, marshResp))
	if addErr != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`add new company error: %s`, addErr.Error()))
		return addErr
	}
	return c.cache.Del(ctx, cacheKey)
}
func (c *Service) Update(ctx context.Context, id int, contact *BaseCompany) error {
	go func() {
		j, err := json.Marshal(contact)
		if err != nil {
			return
		}
		c.logger.InfoLog(fmt.Sprintf(`company with id %d was update: %v`, id, j))
	}()
	if err := c.repo.Update(ctx, id, contact); err != nil {
		return err
	}
	return c.cache.Del(ctx, cacheKey)
}
func (c *Service) Delete(ctx context.Context, id int) error {
	return c.repo.Delete(ctx, id)
}

func NewService(repo IRepository, logger logging.ILogger, cache cache.ICache) *Service {
	return &Service{
		repo:   repo,
		cache:  cache,
		logger: logger,
	}
}
