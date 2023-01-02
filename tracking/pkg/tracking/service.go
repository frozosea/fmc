package tracking

import (
	"context"
	"fmt"
	"golang_tracking/pkg/cache"
	"golang_tracking/pkg/logging"
	"golang_tracking/pkg/tracking/util/scac_accessory"
)

type ContainerTrackingService struct {
	containerTracker *ContainerTracker
	scacRepo         scac_accessory.IRepository
	logger           logging.ILogger
	cache            cache.ICache
}

func NewContainerTrackingService(containerTracker *ContainerTracker, scacRepo scac_accessory.IRepository, logger logging.ILogger, cache cache.ICache) *ContainerTrackingService {
	return &ContainerTrackingService{containerTracker: containerTracker, scacRepo: scacRepo, logger: logger, cache: cache}
}

func (c *ContainerTrackingService) trackByContainerNumber(ctx context.Context, scac, number string) (*ContainerTrackingResponse, error) {
	response, err := c.containerTracker.Track(ctx, scac, number)
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`%s container was not found`, number))
		return nil, err
	}
	go func() {
		ctx := context.Background()
		if err := c.scacRepo.Add(ctx, response.Scac, number); err != nil {
			fmt.Println(err)
			c.logger.ExceptionLog(fmt.Sprintf(`add scac: %s to repo for container number: %s`, response.Scac, number))
		}
		c.logger.InfoLog(fmt.Sprintf(`%s container was found with scac %s`, number, response.Scac))
		if err := c.cache.Set(ctx, number, response); err != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`cache number: %s error: %s`, number, err.Error()))
		}
	}()
	return response, nil
}
func (c *ContainerTrackingService) Track(ctx context.Context, scac, number string) (*ContainerTrackingResponse, error) {
	var r *ContainerTrackingResponse
	if err := c.cache.Get(ctx, number, &r); err == nil {
		fmt.Println(r)
		return r, nil
	}
	if scac == "AUTO" {
		scacFromDb, err := c.scacRepo.Get(ctx, number)
		if err == nil && scacFromDb != "" {
			fmt.Println(scacFromDb)
			return c.trackByContainerNumber(ctx, scacFromDb, number)
		}
		return c.trackByContainerNumber(ctx, scac, number)
	}
	return c.trackByContainerNumber(ctx, scac, number)
}

type BillTrackingService struct {
	billTracker *BillTracker
	scacRepo    scac_accessory.IRepository
	logger      logging.ILogger
	cache       cache.ICache
}

func NewBillTrackingService(billTracker *BillTracker, scacRepo scac_accessory.IRepository, logger logging.ILogger, cache cache.ICache) *BillTrackingService {
	return &BillTrackingService{billTracker: billTracker, scacRepo: scacRepo, logger: logger, cache: cache}
}

func (s *BillTrackingService) trackByBillNumber(ctx context.Context, scac, number string) (*BillNumberTrackingResponse, error) {
	response, err := s.billTracker.Track(ctx, scac, number)
	if err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`%s bill was not found`, number))
		return nil, err
	}
	go func() {
		ctx := context.Background()
		if err := s.scacRepo.Add(ctx, response.Scac, number); err != nil {
			fmt.Println(err)
			s.logger.ExceptionLog(fmt.Sprintf(`add scac: %s to repo for bill number: %s`, response.Scac, number))
		}
		s.logger.InfoLog(fmt.Sprintf(`%s bill was found with scac %s`, number, response.Scac))
		if err := s.cache.Set(ctx, number, response); err != nil {
			s.logger.ExceptionLog(fmt.Sprintf(`cache number: %s error: %s`, number, err.Error()))
		}
	}()
	return response, err
}

func (s *BillTrackingService) Track(ctx context.Context, scac, number string) (*BillNumberTrackingResponse, error) {
	var r *BillNumberTrackingResponse
	if err := s.cache.Get(ctx, number, &r); err == nil {
		fmt.Println(r)
		return r, nil
	}
	if scac == "AUTO" {
		scacFromDb, err := s.scacRepo.Get(ctx, number)
		if err == nil && scacFromDb != "" {
			return s.trackByBillNumber(ctx, scacFromDb, number)
		}
		return s.trackByBillNumber(ctx, scac, number)
	}
	return s.trackByBillNumber(ctx, scac, number)
}
