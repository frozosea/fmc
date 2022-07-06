package scheduler

import (
	"context"
	"time"
)

type AddJobError struct{}

func (a *AddJobError) Error() string {
	return "job with this id already exists."
}

type GetJobError struct{}

func (g *GetJobError) Error() string {
	return "job with this id already exists."
}

type IJobStore interface {
	Save(ctx context.Context, taskId string, task ITask, interval time.Duration, args []interface{}) (*Job, error)
	Get(ctx context.Context, taskId string) (*Job, error)
	Reschedule(ctx context.Context, taskId string, interval time.Duration) (*Job, error)
	Remove(ctx context.Context, taskId string) error
	RemoveAll(ctx context.Context) error
}
type MemoryJobStore struct {
	jobs map[string]*Job
}

func (m *MemoryJobStore) Save(ctx context.Context, taskId string, task ITask, interval time.Duration, args []interface{}) (*Job, error) {
	job := Job{
		Id:          taskId,
		Fn:          task,
		NextRunTime: time.Now().Add(interval),
		Args:        args,
		Interval:    interval,
		Ctx:         ctx,
	}
	getJob := m.jobs[taskId]
	if getJob != nil {
		return getJob, &AddJobError{}
	}
	m.jobs[taskId] = &job
	return &job, nil
}
func (m *MemoryJobStore) Get(_ context.Context, taskId string) (*Job, error) {
	return m.checkTask(taskId)
}
func (m *MemoryJobStore) checkTask(taskId string) (*Job, error) {
	job := m.jobs[taskId]
	if job == nil {
		return nil, &GetJobError{}
	}
	return job, nil
}
func (m *MemoryJobStore) Reschedule(_ context.Context, taskId string, interval time.Duration) (*Job, error) {
	job, err := m.checkTask(taskId)
	if err != nil {
		return job, err
	}
	modifiedJob := Job{
		Id:          taskId,
		Fn:          job.Fn,
		NextRunTime: time.Now().Add(interval),
		Args:        job.Args,
		Interval:    interval,
		Ctx:         job.Ctx,
	}
	m.jobs[taskId] = &modifiedJob
	return &modifiedJob, nil
}
func (m *MemoryJobStore) Remove(_ context.Context, taskId string) error {
	_, err := m.checkTask(taskId)
	if err != nil {
		return err
	}
	delete(m.jobs, taskId)
	return nil
}
func (m *MemoryJobStore) RemoveAll(_ context.Context) error {
	for k := range m.jobs {
		delete(m.jobs, k)
	}
	return nil
}
func NewMemoryJobStore() *MemoryJobStore {
	return &MemoryJobStore{jobs: make(map[string]*Job)}
}
