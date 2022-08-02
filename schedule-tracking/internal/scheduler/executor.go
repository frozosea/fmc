package scheduler

import (
	"context"
	"sync"
	"time"
)

type IJobExecutor interface {
	Run(job *Job) ShouldBeCancelled
	Remove(taskId string) error
}
type Executor struct {
	wg            *sync.WaitGroup
	cancellations map[string]context.CancelFunc
}

func (e *Executor) Run(job *Job) ShouldBeCancelled {
	ctx, cancel := context.WithCancel(job.Ctx)
	e.cancellations[job.Id] = cancel
	e.wg.Add(1)
	return e.process(ctx, job.Fn, job.Interval, job.Args...)
}
func (e *Executor) process(ctx context.Context, task ITask, interval time.Duration, jobArgs ...interface{}) ShouldBeCancelled {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			if shouldBeCancel := task(ctx, jobArgs...); shouldBeCancel {
				return true
			}
		case <-ctx.Done():
			e.wg.Done()
			ticker.Stop()
			return true
		}
	}
}
func (e *Executor) Remove(taskId string) error {
	for jobId, cancel := range e.cancellations {
		if jobId == taskId {
			cancel()
		}
	}
	return nil
}
func NewExecutor() *Executor {
	return &Executor{
		wg:            &sync.WaitGroup{},
		cancellations: make(map[string]context.CancelFunc),
	}
}
