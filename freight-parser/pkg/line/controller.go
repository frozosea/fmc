package line

import (
	"context"
	"encoding/json"
	"fmc-newest/internal/cache"
	"fmc-newest/internal/logging"
	"fmt"
)

type IController interface {
	AddLine(ctx context.Context, lineObj WithByteImage) error
	GetAllLines(ctx context.Context) ([]*Line, error)
}
type controller struct {
	repo        IRepository
	logger      logging.ILogger
	cache       cache.ICache
	fileStorage IFileStorage
}

func (s *controller) AddLine(ctx context.Context, lineObj WithByteImage) error {
	go func() {
		jsonRepr, err := json.Marshal(lineObj)
		if err != nil {
			return
		}
		s.logger.InfoLog(fmt.Sprintf(`line was add: %s`, jsonRepr))
	}()
	stringUrl, uploadImgErr := s.fileStorage.UploadImage(ctx, lineObj.Image)
	if uploadImgErr != nil {
		s.logger.ExceptionLog(fmt.Sprintf(`upload image error: %s`, uploadImgErr.Error()))
		return uploadImgErr
	}
	readyRepoObj := AddRepoLine{BaseLine{
		Scac:     lineObj.Scac,
		FullName: lineObj.FullName,
	}, stringUrl}
	return s.repo.Add(ctx, readyRepoObj)
}
func (s *controller) GetAllLines(ctx context.Context) ([]*Line, error) {
	const cacheKey = "allLines"
	cacheCh := make(chan []*Line)
	go func() {
		var lines []*Line
		s.cache.Get(ctx, cacheKey, &lines)
		cacheCh <- lines
	}()
	repoCh := make(chan []*Line)
	go func() {
		result, repoErr := s.repo.GetAll(ctx)
		if repoErr != nil {
			s.logger.ExceptionLog(fmt.Sprintf(`get all lines from repo error: %s`, repoErr.Error()))
		}
		repoCh <- result
	}()
	select {
	case cacheResult := <-cacheCh:
		return cacheResult, nil
	case result := <-repoCh:
		jsonRepr, err := json.Marshal(result)
		if err != nil {
			return result, nil
		}
		return result, s.cache.Set(ctx, cacheKey, jsonRepr)
	}

}

func NewController(repo IRepository, logger logging.ILogger, cache cache.ICache) *controller {
	return &controller{repo: repo, logger: logger, cache: cache}
}
