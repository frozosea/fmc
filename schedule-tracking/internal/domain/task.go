package domain

import (
	"context"
	"fmt"
	"github.com/frozosea/mailing"
	"os"
	"schedule-tracking/internal/archive"
	excelwriter "schedule-tracking/pkg/excel-writer"
	"schedule-tracking/pkg/logging"
	"schedule-tracking/pkg/scheduler"
	"schedule-tracking/pkg/tracking"
	"strings"
	"time"
)

type CustomTasks struct {
	trackingClient *tracking.Client
	userClient     *UserClient
	arrivedChecker tracking.IArrivedChecker
	logger         logging.ILogger
	writer         excelwriter.IWriter
	mailing        mailing.IMailing
	timeFormatter  ITimeFormatter
	repository     IRepository
	archiveService *archive.Service
	taskManager    *scheduler.Manager
}

func NewCustomTasks(trackingClient *tracking.Client, userClient *UserClient, arrivedChecker tracking.IArrivedChecker, logger logging.ILogger, writer excelwriter.IWriter, mailing mailing.IMailing, timeFormatter ITimeFormatter, repository IRepository, archiveService *archive.Service, taskManager *scheduler.Manager) *CustomTasks {
	return &CustomTasks{
		trackingClient: trackingClient,
		userClient:     userClient,
		arrivedChecker: arrivedChecker,
		logger:         logger, writer: writer,
		mailing:        mailing,
		timeFormatter:  timeFormatter,
		repository:     repository,
		archiveService: archiveService,
		taskManager:    taskManager,
	}
}

func (c *CustomTasks) GetTrackByContainerNumberTask(number string, emails []string, userId int64, emailSubject string) scheduler.ITask {
	fn := func(ctx context.Context) {
		var result tracking.ContainerNumberResponse
		var err error
		for i := 0; i < 3; i++ {
			result, err = c.trackingClient.TrackByContainerNumber(ctx, tracking.Track{
				Number:  number,
				Scac:    "AUTO",
				Country: "OTHER",
			})
			if err == nil {
				break
			} else {
				go c.logger.ExceptionLog(fmt.Sprintf(`track container with Number %s failed: %s`, number, err.Error()))
				if i == 2 {
					return
				}
			}
		}
		fmt.Println(result)
		if c.arrivedChecker.CheckContainerArrived(result) {
			if markErr := c.userClient.MarkContainerWasArrived(ctx, userId, number); markErr != nil {
				c.logger.ExceptionLog(fmt.Sprintf(`mark container is arrived  with: %s err: %s `, result.Container, markErr.Error()))
			}
			if delErr := c.repository.Delete(ctx, []string{number}); delErr != nil {
				c.logger.ExceptionLog(fmt.Sprintf(`delete from tracking containers with Numbers: %s error: %s`, number, delErr.Error()))
			}
			if err := c.userClient.MarkContainerWasRemovedFromTrack(ctx, userId, number); err != nil {
				c.logger.ExceptionLog(fmt.Sprintf(`mark container is removed from tracking containers with number: %s err: %s`, number, err.Error()))
			}
			if err := c.archiveService.AddByContainer(ctx, int(userId), &result); err != nil {
				c.logger.ExceptionLog(fmt.Sprintf(`add container to archive with number: %s err: %s`, number, err.Error()))
			}
			if err := c.taskManager.Remove(context.Background(), number); err != nil {
				c.logger.ExceptionLog(fmt.Sprintf(`remove number: %s from tracking exception: %s`, number, err.Error()))
				return
			}
			for i := 0; i < 3; i++ {
				if err := c.mailing.SendSimple(ctx, emails, fmt.Sprintf("%s was arrived", number), fmt.Sprintf("%s was arrived, and removed from schedule tracking. \nIf our system wrongly removed your cargo, please write your feedback in techical support on our website, we will do all our best. \nThanks for using us!", number), "text/plain"); err != nil {
					continue
				} else {
					break
				}
			}
		}
		pathToFile, writeErr := c.writer.WriteContainerNo(result, c.timeFormatter.Convert)
		if writeErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`write file failed: %s`, err.Error()))
			return
		}
		var subject string
		if emailSubject == " " || emailSubject == "" {
			subject = fmt.Sprintf(`%s Tracking %s`, strings.ToUpper(result.Container), c.timeFormatter.Convert(time.Now()))
		} else {
			subject = emailSubject
		}
		if sendErr := c.mailing.SendWithFile(ctx, emails, subject, pathToFile); sendErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`send mail to %s failed: %s`, emails, sendErr.Error()))
		}
		if removeErr := os.Remove(pathToFile); removeErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`remove %s failed: %s`, pathToFile, removeErr.Error()))
		}
	}
	return fn

}

func (c *CustomTasks) GetTrackByBillNumberTask(number string, emails []string, userId int64, emailSubject string) scheduler.ITask {
	return func(ctx context.Context) {
		var result tracking.BillNumberResponse
		var err error
		for i := 0; i < 3; i++ {
			result, err = c.trackingClient.TrackByBillNumber(ctx, &tracking.Track{
				Number:  number,
				Scac:    "AUTO",
				Country: "OTHER",
			})
			if err == nil {
				break
			} else {
				c.logger.ExceptionLog(fmt.Sprintf(`track container with Number %s failed: %s`, number, err.Error()))
				if i == 2 {
					return
				}
			}
		}
		if c.arrivedChecker.CheckBillNoArrived(result) {
			if markErr := c.userClient.MarkBillNoWasArrived(ctx, userId, number); markErr != nil {
				c.logger.ExceptionLog(fmt.Sprintf(`mark bill no is arrived  with: %s err: %s `, result.BillNo, markErr.Error()))
			}
			if delErr := c.repository.Delete(ctx, []string{number}); delErr != nil {
				c.logger.ExceptionLog(fmt.Sprintf(`delete from tracking containers with Numbers: %s error: %s`, number, delErr.Error()))
			}
			if err := c.userClient.MarkBillNoWasRemovedFromTrack(ctx, userId, number); err != nil {
				c.logger.ExceptionLog(fmt.Sprintf(`mark bill number is removed from tracking containers with number: %s err: %s`, number, err.Error()))
			}
			if err := c.archiveService.AddByBill(ctx, int(userId), &result); err != nil {
				c.logger.ExceptionLog(fmt.Sprintf(`add bill to archive with number: %s err: %s`, number, err.Error()))
			}
			if err := c.taskManager.Remove(context.Background(), number); err != nil {
				c.logger.ExceptionLog(fmt.Sprintf(`remove number: %s from tracking exception: %s`, number, err.Error()))
				return
			}
			for i := 0; i < 3; i++ {
				if err := c.mailing.SendSimple(ctx, emails, fmt.Sprintf("%s was arrived", number), fmt.Sprintf("%s was arrived, and removed from schedule tracking. \nIf our system wrongly removed your cargo, please write your feedback in techical support on our website, we will do all our best. \nThanks for using us!", number), "text/plain"); err != nil {
					continue
				} else {
					break
				}
			}
		}
		pathToFile, writeErr := c.writer.WriteBillNo(result, c.timeFormatter.Convert)
		if writeErr != nil {
			c.logger.FatalLog(fmt.Sprintf(`write file failed: %s`, err.Error()))
			return
		}
		var subject string
		if emailSubject == " " || emailSubject == "" {
			subject = fmt.Sprintf(`%s Tracking %s`, strings.ToUpper(result.BillNo), c.timeFormatter.Convert(time.Now()))
		} else {
			subject = emailSubject
		}
		if sendErr := c.mailing.SendWithFile(ctx, emails, subject, pathToFile); sendErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`send mail to %s failed: %s`, emails, sendErr.Error()))
		}
		if err := os.Remove(pathToFile); err != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`remove %s failed: %s`, pathToFile, err.Error()))
		}
	}
}
