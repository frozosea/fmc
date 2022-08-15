package domain

import (
	"context"
	"fmt"
	"os"
	excelwriter "schedule-tracking/pkg/excel-writer"
	"schedule-tracking/pkg/logging"
	"schedule-tracking/pkg/mailing"
	"schedule-tracking/pkg/scheduler"
	"schedule-tracking/pkg/tracking"
	"strings"
	"sync"
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
}

func NewCustomTasks(trackingClient *tracking.Client, userClient *UserClient, arrivedChecker tracking.IArrivedChecker, logger logging.ILogger, writer excelwriter.IWriter, mailing mailing.IMailing, timeFormatter ITimeFormatter, repository IRepository) *CustomTasks {
	return &CustomTasks{trackingClient: trackingClient, userClient: userClient, arrivedChecker: arrivedChecker, logger: logger, writer: writer, mailing: mailing, timeFormatter: timeFormatter, repository: repository}
}

func (c *CustomTasks) GetTrackByContainerNumberTask(number, country string, userId int64, emailSubject string) scheduler.ITask {
	fn := func(ctx context.Context, emails ...interface{}) scheduler.ShouldBeCancelled {
		var result tracking.ContainerNumberResponse
		var err error
		for i := 0; i < 3; i++ {
			result, err = c.trackingClient.TrackByContainerNumber(ctx, tracking.Track{
				Number:  number,
				Scac:    "AUTO",
				Country: country,
			})
			if err == nil {
				break
			} else {
				go c.logger.ExceptionLog(fmt.Sprintf(`track container with Number %s failed: %s`, number, err.Error()))
				if i == 2 {
					return false
				}
			}
		}
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
			return true
		}
		pathToFile, writeErr := c.writer.WriteContainerNo(result, c.timeFormatter.Convert)
		if writeErr != nil {
			go c.logger.ExceptionLog(fmt.Sprintf(`write file failed: %s`, err.Error()))
			return false
		}
		for _, v := range emails {
			var wg sync.WaitGroup
			var mu sync.Mutex
			wg.Add(1)
			go func() {
				mu.Lock()
				defer mu.Unlock()
				defer wg.Done()
				var subject string
				if emailSubject == " " || emailSubject == "" {
					subject = fmt.Sprintf(`%s Tracking %s`, strings.ToUpper(result.Container), c.timeFormatter.Convert(time.Now()))
				} else {
					subject = emailSubject
				}
				if sendErr := c.mailing.SendWithFile(fmt.Sprintf(`%v`, v), subject, pathToFile); sendErr != nil {
					c.logger.ExceptionLog(fmt.Sprintf(`send mail to %s failed: %s`, v, sendErr.Error()))
				}
			}()
			wg.Wait()
		}
		if removeErr := os.Remove(pathToFile); removeErr != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`remove %s failed: %s`, pathToFile, removeErr.Error()))
		}
		return false
	}
	return fn

}
func (c *CustomTasks) GetTrackByBillNumberTask(number, country string, userId int64, emailSubject string) scheduler.ITask {
	return func(ctx context.Context, emails ...interface{}) scheduler.ShouldBeCancelled {
		var result tracking.BillNumberResponse
		var err error
		for i := 0; i < 3; i++ {
			result, err = c.trackingClient.TrackByBillNumber(ctx, &tracking.Track{
				Number:  number,
				Scac:    "AUTO",
				Country: country,
			})
			if err == nil {
				break
			} else {
				go c.logger.ExceptionLog(fmt.Sprintf(`track container with Number %s failed: %s`, number, err.Error()))
				if i == 2 {
					return false
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
			return true
		}
		pathToFile, writeErr := c.writer.WriteBillNo(result, c.timeFormatter.Convert)
		if writeErr != nil {
			go c.logger.FatalLog(fmt.Sprintf(`write file failed: %s`, err.Error()))
			return false
		}
		for _, v := range emails {
			var wg sync.WaitGroup
			var mu sync.Mutex
			wg.Add(1)
			go func() {
				mu.Lock()
				defer wg.Done()
				defer mu.Unlock()
				var subject string
				if emailSubject == " " || emailSubject == "" {
					subject = fmt.Sprintf(`%s Tracking %s`, strings.ToUpper(result.BillNo), c.timeFormatter.Convert(time.Now()))
				} else {
					subject = emailSubject
				}
				if sendErr := c.mailing.SendWithFile(fmt.Sprintf(`%v`, v), subject, pathToFile); sendErr != nil {
					c.logger.ExceptionLog(fmt.Sprintf(`send mail to %s failed: %s`, v, sendErr.Error()))
				}
			}()
			wg.Wait()
		}
		if err := os.Remove(pathToFile); err != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`remove %s failed: %s`, pathToFile, err.Error()))
		}
		return false
	}
}
