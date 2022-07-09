package scheduler

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type IJobExecutor interface {
	Run(job *Job) ShouldBeCancelled
}
type Executor struct {
	wg            sync.WaitGroup
	cancellations []context.CancelFunc
}

func (e *Executor) Run(job *Job) ShouldBeCancelled {
	ctx, cancel := context.WithCancel(job.Ctx)
	e.cancellations = append(e.cancellations, cancel)
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
			fmt.Println("ctx done")
			e.wg.Done()
			ticker.Stop()
			return true
		}
	}
}
func NewExecutor() *Executor {
	var cancellations []context.CancelFunc
	return &Executor{
		wg:            sync.WaitGroup{},
		cancellations: cancellations,
	}
}
