package domain

import (
	"context"
	"fmt"
	"schedule-tracking/pkg/logging"
	"schedule-tracking/pkg/scheduler"
	"schedule-tracking/pkg/util"
)

type CannotFindEmailError struct{}

func (c *CannotFindEmailError) Error() string {
	return "cannot find email in job"
}

type NumberDoesntBelongThisUserError struct{}

func (n *NumberDoesntBelongThisUserError) Error() string {
	return "number does not belong to this user or cannot find job by your params"
}

type Service struct {
	logger            logging.ILogger
	cli               *UserClient
	taskManager       *scheduler.Manager
	saveResultDirPath string
	repository        IRepository
	*CustomTasks
}

func NewService(logger logging.ILogger, cli *UserClient, taskManager *scheduler.Manager, saveResultDirPath string, repository IRepository, customTasks *CustomTasks) *Service {
	return &Service{logger: logger, cli: cli, taskManager: taskManager, saveResultDirPath: saveResultDirPath, repository: repository, CustomTasks: customTasks}
}
func (p *Service) checkNumberInTaskTable(ctx context.Context, number string, userId int64) bool {
	task, err := p.repository.GetByNumber(ctx, number)
	if task.UserId != userId {
		return false
	}
	if err != nil {
		return false
	}
	return true
}
func (p *Service) addOneContainer(ctx context.Context, number, country, time string, userId int64, emails []string, emailSubject string) (*scheduler.Job, error) {
	if !p.cli.CheckNumberBelongUser(ctx, number, userId, true) {
		return &scheduler.Job{}, &NumberDoesntBelongThisUserError{}
	}
	task := p.GetTrackByContainerNumberTask(number, country, userId, emailSubject)
	job, err := p.taskManager.Add(context.Background(), number, task, time, util.ConvertArgsToInterface(emails)...)
	if err != nil {
		go p.logger.ExceptionLog(fmt.Sprintf(`add job failed: %s`, err.Error()))
		return job, err
	}
	if markErr := p.cli.MarkContainerOnTrack(ctx, userId, number); markErr != nil {
		p.logger.ExceptionLog(fmt.Sprintf(`mark on track container with Number %s failed: %s`, number, markErr.Error()))
		return job, markErr
	}
	go p.logger.InfoLog(fmt.Sprintf(`Number: %s, Time: %s, Emails: %v,userId: %d, IsContainer: true`, job.Id, time, emails, userId))
	return job, nil
}
func (p *Service) AddContainerNumbersOnTrack(ctx context.Context, req []*BaseTrackReq) (*AddOnTrackResponse, error) {
	var alreadyOnTrack []string
	var result []*BaseAddOnTrackResponse
	for _, v := range req {
		job, err := p.addOneContainer(ctx, v.Number, v.Number, v.Time, v.UserId, v.Emails, v.EmailMessageSubject)
		if err != nil {
			switch err.(type) {
			case *scheduler.AddJobError:
				alreadyOnTrack = append(alreadyOnTrack, v.Number)
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
	if addErr := p.repository.Add(ctx, req, true); addErr != nil {
		go func() {
			for _, v := range req {
				p.logger.ExceptionLog(fmt.Sprintf(`add containers with number: %v error: %s`, v.Number, addErr.Error()))
			}
		}()
	}
	return &AddOnTrackResponse{
		result:         result,
		alreadyOnTrack: alreadyOnTrack,
	}, nil
}
func (p *Service) addOneBillOnTrack(ctx context.Context, number, country, time string, userId int64, emails []string, emailSubject string) (*scheduler.Job, error) {
	if !p.cli.CheckNumberBelongUser(ctx, number, userId, false) {
		return &scheduler.Job{}, &NumberDoesntBelongThisUserError{}
	}
	task := p.GetTrackByBillNumberTask(number, country, userId, emailSubject)
	job, err := p.taskManager.Add(context.Background(), number, task, time, util.ConvertArgsToInterface(emails)...)
	if err != nil {
		go p.logger.ExceptionLog(fmt.Sprintf(`add job failed: %s`, err.Error()))
		return job, err
	}
	if addCntrErr := p.cli.MarkBillNoOnTrack(ctx, userId, number); addCntrErr != nil {
		p.logger.ExceptionLog(fmt.Sprintf(`mark bill on track with Number %s failed: %s`, number, addCntrErr.Error()))
		return job, addCntrErr
	}
	go p.logger.InfoLog(fmt.Sprintf(`number: %s, Time: %s, Emails: %v,userId: %d, IsContainer: false`, job.Id, time, emails, userId))
	return job, nil
}
func (p *Service) AddBillNumbersOnTrack(ctx context.Context, req []*BaseTrackReq) (*AddOnTrackResponse, error) {
	var alreadyOnTrack []string
	var result []*BaseAddOnTrackResponse
	for _, v := range req {
		job, err := p.addOneBillOnTrack(ctx, v.Number, v.Country, v.Time, v.UserId, v.Emails, v.EmailMessageSubject)
		if err != nil {
			switch err.(type) {
			case *scheduler.AddJobError:
				alreadyOnTrack = append(alreadyOnTrack, v.Number)
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
	addErr := p.repository.Add(ctx, req, false)
	if addErr != nil {
		go func() {
			for _, v := range req {
				p.logger.ExceptionLog(fmt.Sprintf(`add containers with number: %s error: %s`, v.Number, addErr.Error()))
			}
		}()
	}
	return &AddOnTrackResponse{
		result:         result,
		alreadyOnTrack: alreadyOnTrack,
	}, nil
}
func (p *Service) UpdateTrackingTime(ctx context.Context, numbers []string, newTime string, userId int64) ([]*BaseAddOnTrackResponse, error) {
	var response []*BaseAddOnTrackResponse
	for _, v := range numbers {
		if !p.checkNumberInTaskTable(ctx, v, userId) {
			return response, &NumberDoesntBelongThisUserError{}
		}
		job, err := p.taskManager.Reschedule(ctx, v, newTime)
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
	updErr := p.repository.UpdateTime(ctx, numbers, newTime)
	if updErr != nil {
		p.logger.ExceptionLog(fmt.Sprintf(`update tracking Time with Numbers: %v error: %s`, numbers, updErr.Error()))
	}
	return response, updErr
}
func (p *Service) AddEmailToTracking(ctx context.Context, req AddEmailRequest) error {
	for _, number := range req.numbers {
		if !p.checkNumberInTaskTable(ctx, number, req.userId) {
			return &NumberDoesntBelongThisUserError{}
		}
		job, err := p.taskManager.Get(ctx, number)
		if err != nil {
			return err
		}
		for _, email := range req.emails {
			job.Args = append(job.Args, email)
		}
		if err := p.taskManager.Modify(ctx, job.Id, job.Fn, job.Args...); err != nil {
			return err
		}
	}
	if addErr := p.repository.AddEmails(ctx, req.numbers, req.emails); addErr != nil {
		p.logger.ExceptionLog(fmt.Sprintf(`add Emails: %v to Numbers: %v error: %s`, req.emails, req.numbers, addErr.Error()))
		return addErr
	}
	return nil
}
func (p *Service) DeleteEmailFromTrack(ctx context.Context, req DeleteEmailFromTrack) error {
	if !p.checkNumberInTaskTable(ctx, req.number, req.userId) {
		return &NumberDoesntBelongThisUserError{}
	}
	job, err := p.taskManager.Get(ctx, req.number)
	if err != nil {
		return err
	}
	indexOfEmail := util.GetIndex(req.email, job.Args...)
	if indexOfEmail == -1 {
		return &CannotFindEmailError{}
	}
	job.Args = util.PopForInterfaces(job.Args, indexOfEmail)
	if delErr := p.repository.DeleteEmail(ctx, req.number, req.email); delErr != nil {
		p.logger.ExceptionLog(fmt.Sprintf(`delete email: %s from Number: %s error: %s`, req.email, req.number, delErr.Error()))
	}
	return p.taskManager.Modify(ctx, job.Id, job.Fn, job.Args...)
}
func (p *Service) DeleteFromTracking(ctx context.Context, userId int64, isContainer bool, numbers ...string) error {
	for _, v := range numbers {
		if !p.checkNumberInTaskTable(ctx, v, userId) {
			return &NumberDoesntBelongThisUserError{}
		}
		if err := p.taskManager.Remove(ctx, v); err != nil {
			return err
		}
		if isContainer {
			if markErr := p.cli.MarkContainerWasRemovedFromTrack(ctx, userId, v); markErr != nil {
				return markErr
			}
		} else {
			if markErr := p.cli.MarkBillNoWasRemovedFromTrack(ctx, userId, v); markErr != nil {
				return markErr
			}
		}
	}
	if delErr := p.repository.Delete(ctx, numbers); delErr != nil {
		if isContainer {
			p.logger.ExceptionLog(fmt.Sprintf(`delete from tracking containers with Numbers: %v error: %s`, numbers, delErr.Error()))
		} else {
			p.logger.ExceptionLog(fmt.Sprintf(`delete from tracking bills with Numbers: %v error: %s`, numbers, delErr.Error()))
		}
		return delErr
	}
	return nil
}
func (p *Service) GetInfoAboutTracking(ctx context.Context, number string, userId int64) (*GetInfoAboutTrackResponse, error) {
	if !p.checkNumberInTaskTable(ctx, number, userId) {
		return &GetInfoAboutTrackResponse{}, &NumberDoesntBelongThisUserError{}
	}
	job, err := p.taskManager.Get(ctx, number)
	if err != nil {
		return &GetInfoAboutTrackResponse{}, err
	}
	repoJob, err := p.repository.GetByNumber(ctx, number)
	if err != nil {
		return &GetInfoAboutTrackResponse{}, err
	}
	return &GetInfoAboutTrackResponse{
		number:              number,
		emails:              job.Args,
		nextRunTime:         job.NextRunTime,
		emailMessageSubject: repoJob.EmailMessageSubject,
	}, nil
}

func (p *Service) ChangeEmailMessageSubject(ctx context.Context, userId int64, number, newSubject string) error {
	job, err := p.taskManager.Get(ctx, number)
	if err != nil {
		return err
	}
	repoTask, err := p.repository.GetByNumber(ctx, number)
	if err != nil {
		return err
	}
	task := p.GetTrackByBillNumberTask(number, repoTask.Country, userId, newSubject)
	err = p.taskManager.Modify(context.Background(), number, task, job.Args...)
	if err != nil {
		return err
	}
	if err := p.repository.ChangeEmailMessageSubject(ctx, number, newSubject); err != nil {
		return err
	}
	return nil
}
