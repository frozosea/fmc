package contact

import (
	"context"
	"encoding/json"
	"fmc-newest/internal/cache"
	"fmc-newest/internal/logging"
	"fmt"
)

const cacheKey = "allContacts"

type IController interface {
	AddContact(ctx context.Context, contact BaseContact) error
	GetAllContacts(ctx context.Context) ([]*Contact, error)
}

type controller struct {
	repo   IRepository
	cache  cache.ICache
	logger logging.ILogger
}

func (c *controller) GetAllContacts(ctx context.Context) ([]*Contact, error) {
	cacheCh := make(chan []*Contact)
	go func() {
		var result []*Contact
		if cacheErr := c.cache.Get(ctx, cacheKey, &result); cacheErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`get all contacts from cache error: %s`, cacheErr.Error()))
		}
		cacheCh <- result
	}()
	repoCh := make(chan []*Contact)
	go func() {
		repoRes, repoErr := c.repo.GetAll(ctx)
		if repoErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`get all contacts from repository error: %s`, repoErr.Error()))
		}
		repoCh <- repoRes
	}()
	select {
	case result := <-cacheCh:
		return result, nil
	case repoResult := <-repoCh:
		jsonRepresent, jsonMarshErr := json.Marshal(repoResult)
		if jsonMarshErr != nil {
			return repoResult, nil
		}
		return repoResult, c.cache.Set(ctx, cacheKey, jsonRepresent)
	}
}
func (c *controller) AddContact(ctx context.Context, contact BaseContact) error {
	addErr := c.repo.Add(ctx, contact)
	marshResp, marshErr := json.Marshal(contact)
	if marshErr != nil {
		return marshErr
	}
	c.logger.InfoLog(fmt.Sprintf(`add new contact: %s`, marshResp))
	if addErr != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`add new contact error: %s`, addErr.Error()))
		return addErr
	}
	return c.cache.Del(ctx, cacheKey)
}
func NewController(repo IRepository, logger logging.ILogger, cache cache.ICache) *controller {
	return &controller{
		repo:   repo,
		cache:  cache,
		logger: logger,
	}
}
