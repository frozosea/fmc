package scheduler

import (
	"context"
	"log"
	"sync"
	"time"
)

type IJobExecutor interface {
	Run(job *Job)
	Remove(taskId string) error
}
type Executor struct {
	wg            *sync.WaitGroup
	cancellations map[string]context.CancelFunc
	jobStore      IJobStore
	logger        *log.Logger
	timeParser    ITimeParser
}

func (e *Executor) Run(job *Job) {
	ctx, cancel := context.WithCancel(job.Ctx)
	job.Ctx = ctx
	e.cancellations[job.Id] = cancel
	e.wg.Add(1)
	if shouldBeCancel := e.process(job); shouldBeCancel {
		if err := e.Remove(job.Id); err != nil {
			e.logger.Printf(`remove job with id: %s err: %s`, job.Id, err.Error())
		}
		e.logger.Printf(`job with id %s was removed`, job.Id)
	}
}
func (e *Executor) process(job *Job) ShouldBeCancelled {
	ticker := time.NewTicker(job.Interval)
	for {
		select {
		case <-ticker.C:
			e.logger.Printf(`job with id : %s now run`, job.Id)
			if shouldBeCancelled := job.Fn(job.Ctx, job.Args...); shouldBeCancelled {
				return shouldBeCancelled
			}
			nextInterval, err := e.timeParser.Parse(job.Time)
			if err != nil {
				continue
			}
			ticker.Reset(nextInterval)
			job.NextRunTime = time.Now().Add(nextInterval)
			continue
		case <-job.Ctx.Done():
			e.logger.Printf(`job with id:%s ctx done`, job.Id)
			e.wg.Done()
			ticker.Stop()
			if err := e.jobStore.Remove(job.Ctx, job.Id); err != nil {
				e.logger.Printf(`remove task with id: %s from jobstore error: %s`, job.Id, err)
				return true
			}
			return true
		default:
			continue
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
func NewExecutor(jobStore IJobStore, timeParser ITimeParser, logger *log.Logger) *Executor {
	return &Executor{
		wg:            &sync.WaitGroup{},
		cancellations: make(map[string]context.CancelFunc),
		jobStore:      jobStore,
		logger:        logger,
		timeParser:    timeParser,
	}
}
