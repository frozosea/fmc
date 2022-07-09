package schedule_tracking

import (
	"context"
	"fmt"
	"user-api/internal/cache"
	"user-api/internal/logging"
	"user-api/internal/scheduler"
	"user-api/internal/util"
)

type CannotFindEmailError struct{}

func (c *CannotFindEmailError) Error() string {
	return "cannot find email in job"
}

type Controller struct {
	logger            logging.ILogger
	repository        IRepository
	taskManager       *scheduler.Manager
	cache             cache.ICache
	saveResultDirPath string
	*customTasks
}

func (c *Controller) addOneContainer(ctx context.Context, number, country, time string, userId int64, emails []string) (*scheduler.Job, error) {
	task := c.GetTrackByContainerNumberTask(number, country, userId)
	job, err := c.taskManager.Add(context.Background(), number, task, time, util.ConvertArgsToInterface(emails)...)
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`add job failed: %s`, err.Error()))
		return job, err
	}
	if addMarkErr := c.repository.AddMarkContainerOnTrack(ctx, number, int(userId)); addMarkErr != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`add mark container is on track failed: %s`, err.Error()))
		return job, addMarkErr
	}
	go c.logger.InfoLog(fmt.Sprintf(`add job success: %s, nextRunTime: %s`, job.Id, job.NextRunTime.String()))
	return job, nil
}
func (c *Controller) AddContainerNumbersOnTrack(ctx context.Context, req TrackByContainerNoReq) (*AddOnTrackResponse, error) {
	var alreadyOnTrack []string
	var result []*BaseAddOnTrackResponse
	for _, v := range req.numbers {
		job, err := c.addOneContainer(ctx, v, req.country, req.time, req.userId, req.emails)
		switch err.(type) {
		case *scheduler.AddJobError:
			alreadyOnTrack = append(alreadyOnTrack, v)
		}
		result = append(result, &BaseAddOnTrackResponse{
			success:     true,
			number:      job.Id,
			nextRunTime: job.NextRunTime,
		})
	}
	return &AddOnTrackResponse{
		result:         result,
		alreadyOnTrack: alreadyOnTrack,
	}, nil
}
func (c *Controller) addOneBillOnTrack(ctx context.Context, number, country, time string, userId int64, emails []string) (*scheduler.Job, error) {
	task := c.GetTrackByBillNumberTask(number, country, userId)
	job, err := c.taskManager.Add(context.Background(), number, task, time, util.ConvertArgsToInterface(emails)...)
	if err != nil {
		//go c.logger.ExceptionLog(fmt.Sprintf(`add job failed: %s`, err.Error()))
		return job, err
	}
	if err := c.repository.AddMarkBillNoOnTrack(ctx, number, int(userId)); err != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`add mark container is on track failed: %s`, err.Error()))
		return job, err
	}
	go c.logger.InfoLog(fmt.Sprintf(`add job success: %s, nextRunTime: %s`, job.Id, job.NextRunTime.String()))
	return job, nil
}
func (c *Controller) AddBillNumbersOnTrack(ctx context.Context, req TrackByBillNoReq) (*AddOnTrackResponse, error) {
	var alreadyOnTrack []string
	var result []*BaseAddOnTrackResponse
	for _, v := range req.numbers {
		job, err := c.addOneBillOnTrack(ctx, v, req.country, req.time, req.userId, req.emails)
		switch err.(type) {
		case *scheduler.AddJobError:
			alreadyOnTrack = append(alreadyOnTrack, v)
		}
		result = append(result, &BaseAddOnTrackResponse{
			success:     true,
			number:      job.Id,
			nextRunTime: job.NextRunTime,
		})
	}
	return &AddOnTrackResponse{
		result:         result,
		alreadyOnTrack: alreadyOnTrack,
	}, nil
}
func (c *Controller) UpdateTrackingTime(ctx context.Context, numbers []string, newTime string) ([]*BaseAddOnTrackResponse, error) {
	var response []*BaseAddOnTrackResponse
	for _, v := range numbers {
		job, err := c.taskManager.Reschedule(ctx, v, newTime)
		if err != nil {
			return response, err
		}
		oneStruct := &BaseAddOnTrackResponse{
			success:     true,
			number:      v,
			nextRunTime: job.NextRunTime,
		}
		response = append(response, oneStruct)
	}
	return response, nil
}
func (c *Controller) AddEmailToTracking(ctx context.Context, req AddEmailRequest) error {
	for _, number := range req.numbers {
		job, err := c.taskManager.Get(ctx, number)
		if err != nil {
			return err
		}
		for _, email := range req.emails {
			job.Args = append(job.Args, email)
		}
		if err := c.taskManager.Modify(ctx, job.Id, job.Fn, job.Args...); err != nil {
			return err
		}
	}
	return nil
}
func (c *Controller) DeleteEmailFromTrack(ctx context.Context, req DeleteEmailFromTrack) error {
	job, err := c.taskManager.Get(ctx, req.number)
	if err != nil {
		return err
	}
	indexOfEmail := util.GetIndex(req.email, job.Args...)
	if indexOfEmail == -1 {
		return &CannotFindEmailError{}
	}
	job.Args = util.PopForInterfaces(job.Args, indexOfEmail)
	return c.taskManager.Modify(ctx, job.Id, job.Fn, job.Args...)
}
func (c *Controller) DeleteFromTracking(ctx context.Context, userId int, isContainer bool, number ...string) error {
	for _, v := range number {
		if err := c.taskManager.Remove(ctx, v); err != nil {
			return err
		}
		if isContainer {
			if err := c.repository.AddMarkContainerWasRemovedFromTrack(ctx, v, userId); err != nil {
				return err
			}
			if deleteFromCacheErr := c.cache.Del(ctx, fmt.Sprintf(`%d`, userId)); deleteFromCacheErr != nil {
				return deleteFromCacheErr
			}
		} else {
			if err := c.repository.AddMarkBillNoWasRemovedFromTrack(ctx, v, userId); err != nil {
				return err
			}
			if deleteFromCacheErr := c.cache.Del(ctx, fmt.Sprintf(`%d`, userId)); deleteFromCacheErr != nil {
				return deleteFromCacheErr
			}
		}
	}
	return nil
}
func (c *Controller) GetInfoAboutTracking(ctx context.Context, number string) (*GetInfoAboutTrackResponse, error) {
	job, err := c.taskManager.Get(ctx, number)
	if err != nil {
		return &GetInfoAboutTrackResponse{}, err
	}
	return &GetInfoAboutTrackResponse{
		number:      number,
		emails:      job.Args,
		nextRunTime: job.NextRunTime,
	}, nil
}
func NewController(logger logging.ILogger, repository IRepository, taskManager *scheduler.Manager, taskGetter *customTasks, cache cache.ICache) *Controller {
	return &Controller{logger: logger, repository: repository, taskManager: taskManager, customTasks: taskGetter, cache: cache}
}
