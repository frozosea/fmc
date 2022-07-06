package scheduler

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"
)

type ShouldBeCancelled bool
type ITask func(ctx context.Context, args ...interface{}) ShouldBeCancelled

type Job struct {
	Id          string
	Fn          ITask
	NextRunTime time.Time
	Args        []interface{}
	Interval    time.Duration
	Ctx         context.Context
}
type Manager struct {
	executor   IJobExecutor
	jobstore   IJobStore
	timeParser ITimeParser
	baseLogger log.Logger
}

func (m *Manager) Start() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
}
func (m *Manager) Add(ctx context.Context, taskId string, task ITask, timeStr string, taskArgs ...interface{}) (*Job, error) {
	taskTime, err := m.timeParser.ParseHourMinuteString(timeStr)
	if err != nil {
		var job *Job
		return job, err
	}
	job, err := m.jobstore.Save(ctx, taskId, task, taskTime, taskArgs)
	if err != nil {
		m.baseLogger.Println(fmt.Sprintf(`add task with id: %s err: %s`, taskId, err.Error()))
		return job, err
	}
	go func() {
		shouldBeCancel := m.executor.Run(job)
		if shouldBeCancel {
			if removeErr := m.jobstore.Remove(ctx, taskId); removeErr != nil {
				m.baseLogger.Println(fmt.Sprintf(`remove task with id: %s`, taskId))
			}
			job.Ctx.Done()
		}
		job.NextRunTime = time.Now().Add(job.Interval)
	}()
	return job, err
}
func (m *Manager) AddWithDuration(ctx context.Context, taskId string, task ITask, interval time.Duration, taskArgs ...interface{}) (*Job, error) {
	job, err := m.jobstore.Save(ctx, taskId, task, interval, taskArgs)
	fmt.Println(job.Args)
	if err != nil {
		m.baseLogger.Println(fmt.Sprintf(`add task with id: %s err: %s`, taskId, err.Error()))
		return job, err
	}
	go func() {
		shouldBeCancel := m.executor.Run(job)
		if shouldBeCancel {
			if removeErr := m.jobstore.Remove(ctx, taskId); removeErr != nil {
				m.baseLogger.Println(fmt.Sprintf(`remove task with id: %s`, taskId))
			}
		}
	}()
	return job, err
}
func (m *Manager) Get(ctx context.Context, taskId string) (*Job, error) {
	return m.jobstore.Get(ctx, taskId)
}
func (m *Manager) Reschedule(ctx context.Context, taskId string, timeStr string) (*Job, error) {
	taskTime, err := m.timeParser.ParseHourMinuteString(timeStr)
	if err != nil {
		var job *Job
		return job, err
	}
	return m.jobstore.Reschedule(ctx, taskId, taskTime)
}
func (m *Manager) RescheduleWithDuration(ctx context.Context, taskId string, newInterval time.Duration) (*Job, error) {
	return m.jobstore.Reschedule(ctx, taskId, newInterval)
}
func (m *Manager) Remove(ctx context.Context, taskId string) error {
	return m.jobstore.Remove(ctx, taskId)
}
func (m *Manager) RemoveAll(ctx context.Context) error {
	return m.jobstore.RemoveAll(ctx)
}
func (m *Manager) Modify(ctx context.Context, taskId string, task ITask, args ...interface{}) error {
	job, err := m.jobstore.Get(ctx, taskId)
	if err != nil {
		return err
	}
	job.Fn = task
	job.Args = args
	if err := m.jobstore.Remove(ctx, taskId); err != nil {
		return err
	}
	_, err = m.jobstore.Save(job.Ctx, job.Id, job.Fn, job.Interval, job.Args)
	if err != nil {
		return err
	}
	return nil
}

func NewDefaultScheduler() *Manager {
	return &Manager{executor: NewExecutor(), jobstore: NewMemoryJobStore(), baseLogger: *log.New(os.Stdout, "log", 1), timeParser: NewTimeParser()}

}
func NewDefaultSchedulerWithCustomLogger(out io.Writer) *Manager {
	return &Manager{executor: NewExecutor(), jobstore: NewMemoryJobStore(), baseLogger: *log.New(out, "log", 1), timeParser: NewTimeParser()}

}