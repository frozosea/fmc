package schedule_tracking

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	excelwriter "user-api/internal/excel-writer"
	"user-api/internal/logging"
	"user-api/internal/mailing"
	"user-api/internal/scheduler"
	tracking2 "user-api/internal/tracking"
)

//const (
//	SUBJECT_OF_TRACKING_MAILING = ""
//)

type customTasks struct {
	client         *tracking2.Client
	arrivedChecker tracking2.IArrivedChecker
	logger         logging.ILogger
	writer         excelwriter.IWriter
	mailing        mailing.IMailing
	timeFormatter  ITimeFormatter
	repository     IRepository
}

func (c *customTasks) GetTrackByContainerNumberTask(number, country string, userId int64) scheduler.ITask {
	fn := func(ctx context.Context, emails ...interface{}) scheduler.ShouldBeCancelled {
		result, err := c.client.TrackByContainerNumber(ctx, tracking2.Track{
			Number:  number,
			Scac:    "AUTO",
			Country: country,
		})
		if err != nil {
			go c.logger.ExceptionLog(fmt.Sprintf(`track container with number %s failed: %s`, number, err.Error()))
			return true
		}
		if c.arrivedChecker.CheckContainerArrived(result) {
			go func() {
				if markErr := c.repository.AddMarkContainerWasArrived(ctx, result.Container, int(userId)); markErr != nil {
					c.logger.ExceptionLog(fmt.Sprintf(`mark container is arrived  with: %s err: %s `, result.Container, markErr.Error()))
				}
			}()
			return true
		}
		pathToFile, writeErr := c.writer.WriteContainerNo(result, c.timeFormatter.Convert)
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
				defer mu.Unlock()
				defer wg.Done()
				if err := c.mailing.SendWithFile(fmt.Sprintf(`%v`, v), fmt.Sprintf(`%s Tracking %s`, strings.ToUpper(result.Container), c.timeFormatter.Convert(time.Now())), pathToFile); err != nil {
					c.logger.ExceptionLog(fmt.Sprintf(`send mail to %s failed: %s`, v, err.Error()))
				}
			}()
			wg.Wait()
		}
		if err := os.Remove(pathToFile); err != nil {
			c.logger.ExceptionLog(fmt.Sprintf(`remove %s failed: %s`, pathToFile, err.Error()))
		}
		return false
	}
	return fn

}
func (c *customTasks) GetTrackByBillNumberTask(number, country string, userId int64) scheduler.ITask {
	return func(ctx context.Context, emails ...interface{}) scheduler.ShouldBeCancelled {
		result, err := c.client.TrackByBillNumber(ctx, &tracking2.Track{
			Number:  number,
			Scac:    "AUTO",
			Country: country,
		})
		if err != nil {
			go c.logger.ExceptionLog(fmt.Sprintf(`track container with number %s failed: %s`, number, err.Error()))
			return true
		}
		if c.arrivedChecker.CheckBillNoArrived(result) {
			go func() {
				if markErr := c.repository.AddMarkBillNoWasArrived(ctx, result.BillNo, int(userId)); markErr != nil {
					c.logger.ExceptionLog(fmt.Sprintf(`mark billNo is arrived  with: %s err: %s `, result.BillNo, markErr.Error()))
				}
			}()
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
				if err := c.mailing.SendWithFile(fmt.Sprintf(`%v`, v), fmt.Sprintf(`%s Tracking %s`, strings.ToUpper(result.BillNo), c.timeFormatter.Convert(time.Now())), pathToFile); err != nil {
					c.logger.ExceptionLog(fmt.Sprintf(`send mail to %s failed: %s`, v, err.Error()))
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
func NewCustomTasks(client *tracking2.Client, arrivedChecker tracking2.IArrivedChecker, logger logging.ILogger, writer excelwriter.IWriter, mailing mailing.IMailing, timeFormatter ITimeFormatter, repository IRepository) *customTasks {
	return &customTasks{client: client, arrivedChecker: arrivedChecker, logger: logger, writer: writer, mailing: mailing, timeFormatter: timeFormatter, repository: repository}
}
