package domain

import (
	"context"
	"fmt"
	"schedule-tracking/pkg/logging"
	"schedule-tracking/pkg/scheduler"
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

func (s *Service) checkNumberInTaskTable(ctx context.Context, number string, userId int64) bool {
	task, err := s.repository.GetByNumber(ctx, number)
	if task.UserId != userId || err != nil {
		return false
	}
	return true
}

func (s *Service) addOneContainer(ctx context.Context, number, time string, userId int64, emails []string, emailSubject string) (*scheduler.Job, error) {
	if !s.cli.CheckNumberBelongUser(ctx, number, userId, true) {
		return nil, &NumberDoesntBelongThisUserError{}
	}
	task := s.GetTrackByContainerNumberTask(number, emails, userId, emailSubject)
	job, err := s.taskManager.Add(context.Background(), number, task, time)
	if err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`add job failed: %s`, err.Error()))
		return nil, err
	}
	if markErr := s.cli.MarkContainerOnTrack(ctx, userId, number); markErr != nil {
		s.logger.ExceptionLog(fmt.Sprintf(`mark on track container with Number %s failed: %s`, number, markErr.Error()))
		return nil, markErr
	}
	go s.logger.InfoLog(fmt.Sprintf(`Number: %s, Time: %s, Emails: %v,userId: %d, IsContainer: true`, job.Id, time, emails, userId))
	return job, nil
}

func (s *Service) AddContainerNumbersOnTrack(ctx context.Context, req *BaseTrackReq) (*AddOnTrackResponse, error) {
	var alreadyOnTrack []string
	var result []*BaseAddOnTrackResponse
	for _, v := range req.Numbers {
		job, err := s.addOneContainer(ctx, v, req.Time, req.UserId, req.Emails, req.EmailMessageSubject)
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
	if addErr := s.repository.Add(ctx, req, true); addErr != nil {
		go func() {
			for _, v := range req.Numbers {
				s.logger.ExceptionLog(fmt.Sprintf(`add containers with number: %v error: %s`, v, addErr.Error()))
			}
		}()
	}
	return &AddOnTrackResponse{
		result:         result,
		alreadyOnTrack: alreadyOnTrack,
	}, nil
}

func (s *Service) addOneBillOnTrack(ctx context.Context, number, time string, userId int64, emails []string, emailSubject string) (*scheduler.Job, error) {
	if !s.cli.CheckNumberBelongUser(ctx, number, userId, false) {
		return nil, &NumberDoesntBelongThisUserError{}
	}
	task := s.GetTrackByBillNumberTask(number, emails, userId, emailSubject)
	job, err := s.taskManager.Add(context.Background(), number, task, time)
	if err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`add job failed: %s`, err.Error()))
		return nil, err
	}
	if markErr := s.cli.MarkBillNoOnTrack(ctx, userId, number); markErr != nil {
		s.logger.ExceptionLog(fmt.Sprintf(`mark on track container with Number %s failed: %s`, number, markErr.Error()))
		return nil, markErr
	}
	go s.logger.InfoLog(fmt.Sprintf(`Number: %s, Time: %s, Emails: %v,userId: %d, IsContainer: true`, job.Id, time, emails, userId))
	return job, nil
}

func (s *Service) AddBillNumbersOnTrack(ctx context.Context, req *BaseTrackReq) (*AddOnTrackResponse, error) {
	var alreadyOnTrack []string
	var result []*BaseAddOnTrackResponse
	for _, v := range req.Numbers {
		job, err := s.addOneBillOnTrack(ctx, v, req.Time, req.UserId, req.Emails, req.EmailMessageSubject)
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
	addErr := s.repository.Add(ctx, req, false)
	if addErr != nil {
		go func() {
			for _, v := range req.Numbers {
				s.logger.ExceptionLog(fmt.Sprintf(`add containers with number: %s error: %s`, v, addErr.Error()))
			}
		}()
	}
	return &AddOnTrackResponse{
		result:         result,
		alreadyOnTrack: alreadyOnTrack,
	}, nil
}

func (s *Service) DeleteFromTracking(ctx context.Context, userId int64, isContainer bool, numbers []string) error {
	for _, v := range numbers {
		if !s.checkNumberInTaskTable(ctx, v, userId) {
			return &NumberDoesntBelongThisUserError{}
		}
		if err := s.taskManager.Remove(ctx, v); err != nil {
			return err
		}
		if isContainer {
			if markErr := s.cli.MarkContainerWasRemovedFromTrack(ctx, userId, v); markErr != nil {
				return markErr
			}
		} else {
			if markErr := s.cli.MarkBillNoWasRemovedFromTrack(ctx, userId, v); markErr != nil {
				return markErr
			}
		}
	}
	if delErr := s.repository.Delete(ctx, numbers); delErr != nil {
		if isContainer {
			s.logger.ExceptionLog(fmt.Sprintf(`delete from tracking containers with Numbers: %v error: %s`, numbers, delErr.Error()))
		} else {
			s.logger.ExceptionLog(fmt.Sprintf(`delete from tracking bills with Numbers: %v error: %s`, numbers, delErr.Error()))
		}
		return delErr
	}
	return nil
}

func (s *Service) GetInfoAboutTracking(ctx context.Context, number string, userId int64) (*GetInfoAboutTrackResponse, error) {
	if !s.checkNumberInTaskTable(ctx, number, userId) {
		return &GetInfoAboutTrackResponse{}, &NumberDoesntBelongThisUserError{}
	}
	repoJob, err := s.repository.GetByNumber(ctx, number)
	if err != nil {
		return &GetInfoAboutTrackResponse{
			Number:               repoJob.Number,
			IsContainer:          repoJob.IsContainer,
			IsOnTrack:            false,
			ScheduleTrackingInfo: &ScheduleTrackingInfo{},
		}, nil
	}
	_, err = s.taskManager.Get(ctx, number)
	if err != nil {
		return &GetInfoAboutTrackResponse{
			Number:               repoJob.Number,
			IsContainer:          repoJob.IsContainer,
			IsOnTrack:            false,
			ScheduleTrackingInfo: &ScheduleTrackingInfo{},
		}, err
	}
	return &GetInfoAboutTrackResponse{
		Number:      repoJob.Number,
		IsContainer: repoJob.IsContainer,
		IsOnTrack:   true,
		ScheduleTrackingInfo: &ScheduleTrackingInfo{
			Time:    repoJob.Time,
			Subject: repoJob.EmailMessageSubject,
			Emails:  repoJob.Emails,
		},
	}, nil
}

func (s *Service) Update(ctx context.Context, r *BaseTrackReq, isContainer bool) error {
	if err := s.DeleteFromTracking(ctx, r.UserId, isContainer, r.Numbers); err != nil {
		return err
	}
	if isContainer {
		if _, err := s.AddContainerNumbersOnTrack(ctx, r); err != nil {
			return err
		}

	} else {
		if _, err := s.AddBillNumbersOnTrack(ctx, r); err != nil {
			return err
		}
	}
	return nil

}
