package domain

import (
	"context"
	"fmt"
	"schedule-tracking/internal/logging"
	"schedule-tracking/internal/scheduler"
	"schedule-tracking/internal/util"
)

type CannotFindEmailError struct{}

func (c *CannotFindEmailError) Error() string {
	return "cannot find email in job"
}

type NumberDoesntBelongThisUserError struct{}

func (n *NumberDoesntBelongThisUserError) Error() string {
	return "number does not belong to this user or cannot find job by your params"
}

type Controller struct {
	logger            logging.ILogger
	cli               *UserClient
	taskManager       *scheduler.Manager
	saveResultDirPath string
	repository        IRepository
	*CustomTasks
}

func NewController(logger logging.ILogger, cli *UserClient, taskManager *scheduler.Manager, saveResultDirPath string, repository IRepository, customTasks *CustomTasks) *Controller {
	return &Controller{logger: logger, cli: cli, taskManager: taskManager, saveResultDirPath: saveResultDirPath, repository: repository, CustomTasks: customTasks}
}
func (c *Controller) checkNumberBelongUser(ctx context.Context, number string, userId int64) bool {
	task, err := c.repository.GetByNumber(ctx, number)
	if task.UserId != userId {
		return false
	}
	if err != nil {
		return false
	}
	return true
}
func (c *Controller) addOneContainer(ctx context.Context, number, country, time string, userId int64, emails []string) (*scheduler.Job, error) {
	task := c.GetTrackByContainerNumberTask(number, country, userId)
	job, err := c.taskManager.Add(context.Background(), number, task, time, util.ConvertArgsToInterface(emails)...)
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`add job failed: %s`, err.Error()))
		return job, err
	}
	if markErr := c.cli.MarkContainerOnTrack(ctx, userId, number); markErr != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`mark on track container with Number %s failed: %s`, number, markErr.Error()))
		return job, markErr
	}
	go c.logger.InfoLog(fmt.Sprintf(`Number: %s, Time: %s, Emails: %v,UserId: %d, IsContainer: true`, job.Id, time, emails, userId))
	return job, nil
}
func (c *Controller) AddContainerNumbersOnTrack(ctx context.Context, req TrackByContainerNoReq) (*AddOnTrackResponse, error) {
	var alreadyOnTrack []string
	var result []*BaseAddOnTrackResponse
	for _, v := range req.Numbers {
		job, err := c.addOneContainer(ctx, v, req.Country, req.Time, req.UserId, req.Emails)
		if err != nil {
			switch err.(type) {
			case *scheduler.AddJobError:
				alreadyOnTrack = append(alreadyOnTrack, v)
				continue
			default:
				return &AddOnTrackResponse{
					result:         result,
					alreadyOnTrack: alreadyOnTrack,
				}, err
			}
		}
		result = append(result, &BaseAddOnTrackResponse{
			success:     true,
			number:      job.Id,
			nextRunTime: job.NextRunTime,
		})
	}
	addErr := c.repository.Add(ctx, &req.BaseTrackReq, true)
	if addErr != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`add containers with Numbers: %v error: %s`, req.Numbers, addErr.Error()))
	}
	return &AddOnTrackResponse{
		result:         result,
		alreadyOnTrack: alreadyOnTrack,
	}, addErr
}
func (c *Controller) addOneBillOnTrack(ctx context.Context, number, country, time string, userId int64, emails []string) (*scheduler.Job, error) {
	task := c.GetTrackByBillNumberTask(number, country, userId)
	job, err := c.taskManager.Add(context.Background(), number, task, time, util.ConvertArgsToInterface(emails)...)
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`add job failed: %s`, err.Error()))
		return job, err
	}
	if addCntrErr := c.cli.MarkBillNoOnTrack(ctx, userId, number); addCntrErr != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`mark bill on track with Number %s failed: %s`, number, addCntrErr.Error()))
		return job, addCntrErr
	}
	go c.logger.InfoLog(fmt.Sprintf(`Number: %s, Time: %s, Emails: %v,UserId: %d, IsContainer: false`, job.Id, time, emails, userId))
	return job, nil
}
func (c *Controller) AddBillNumbersOnTrack(ctx context.Context, req TrackByBillNoReq) (*AddOnTrackResponse, error) {
	var alreadyOnTrack []string
	var result []*BaseAddOnTrackResponse
	for _, v := range req.Numbers {
		job, err := c.addOneBillOnTrack(ctx, v, req.Country, req.Time, req.UserId, req.Emails)
		if err != nil {
			switch err.(type) {
			case *scheduler.AddJobError:
				alreadyOnTrack = append(alreadyOnTrack, v)
				continue
			default:
				return &AddOnTrackResponse{
					result:         result,
					alreadyOnTrack: alreadyOnTrack,
				}, err
			}
		}
		result = append(result, &BaseAddOnTrackResponse{
			success:     true,
			number:      job.Id,
			nextRunTime: job.NextRunTime,
		})
	}
	addErr := c.repository.Add(ctx, &req.BaseTrackReq, false)
	if addErr != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`add containers with Numbers: %v error: %s`, req.Numbers, addErr.Error()))
	}
	return &AddOnTrackResponse{
		result:         result,
		alreadyOnTrack: alreadyOnTrack,
	}, nil
}
func (c *Controller) UpdateTrackingTime(ctx context.Context, numbers []string, newTime string, userId int64) ([]*BaseAddOnTrackResponse, error) {
	var response []*BaseAddOnTrackResponse
	for _, v := range numbers {
		if !c.checkNumberBelongUser(ctx, v, userId) {
			return response, &NumberDoesntBelongThisUserError{}
		}
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
	updErr := c.repository.UpdateTime(ctx, numbers, newTime)
	if updErr != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`update tracking Time with Numbers: %v error: %s`, numbers, updErr.Error()))
	}
	return response, updErr
}
func (c *Controller) AddEmailToTracking(ctx context.Context, req AddEmailRequest) error {
	for _, number := range req.numbers {
		if !c.checkNumberBelongUser(ctx, number, req.userId) {
			return &NumberDoesntBelongThisUserError{}
		}
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
	if addErr := c.repository.AddEmails(ctx, req.numbers, req.emails); addErr != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`add Emails: %v to Numbers: %v error: %s`, req.emails, req.numbers, addErr.Error()))
		return addErr
	}
	return nil
}
func (c *Controller) DeleteEmailFromTrack(ctx context.Context, req DeleteEmailFromTrack) error {
	if !c.checkNumberBelongUser(ctx, req.number, req.userId) {
		return &NumberDoesntBelongThisUserError{}
	}
	job, err := c.taskManager.Get(ctx, req.number)
	if err != nil {
		return err
	}
	indexOfEmail := util.GetIndex(req.email, job.Args...)
	if indexOfEmail == -1 {
		return &CannotFindEmailError{}
	}
	job.Args = util.PopForInterfaces(job.Args, indexOfEmail)
	if delErr := c.repository.DeleteEmail(ctx, req.number, req.email); delErr != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`delete email: %s from Number: %s error: %s`, req.email, req.number, delErr.Error()))
	}
	return c.taskManager.Modify(ctx, job.Id, job.Fn, job.Args...)
}
func (c *Controller) DeleteFromTracking(ctx context.Context, userId int64, isContainer bool, numbers ...string) error {
	for _, v := range numbers {
		if !c.checkNumberBelongUser(ctx, v, userId) {
			return &NumberDoesntBelongThisUserError{}
		}
		if err := c.taskManager.Remove(ctx, v); err != nil {
			return err
		}
		if isContainer {
			if markErr := c.cli.MarkContainerWasRemovedFromTrack(ctx, userId, v); markErr != nil {
				return markErr
			}
		} else {
			if markErr := c.cli.MarkBillNoWasRemovedFromTrack(ctx, userId, v); markErr != nil {
				return markErr
			}
		}
	}
	if delErr := c.repository.Delete(ctx, numbers); delErr != nil {
		if isContainer {
			c.logger.ExceptionLog(fmt.Sprintf(`delete from tracking containers with Numbers: %v error: %s`, numbers, delErr.Error()))
		} else {
			c.logger.ExceptionLog(fmt.Sprintf(`delete from tracking bills with Numbers: %v error: %s`, numbers, delErr.Error()))
		}
		return delErr
	}
	return nil
}
func (c *Controller) GetInfoAboutTracking(ctx context.Context, number string, userId int64) (*GetInfoAboutTrackResponse, error) {
	if !c.checkNumberBelongUser(ctx, number, userId) {
		return &GetInfoAboutTrackResponse{}, &NumberDoesntBelongThisUserError{}
	}
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
