package login_provider

import (
	"context"
	scheduler "github.com/frozosea/scheduler/pkg"
	"time"
)

type TaskGenerator struct {
	store    *Store
	provider *Provider
}

func NewTaskGenerator(store *Store, provider *Provider) *TaskGenerator {
	return &TaskGenerator{store: store, provider: provider}
}

func (t *TaskGenerator) Generate() scheduler.ITask {
	return func(ctx context.Context, args ...interface{}) scheduler.ShouldBeCancelled {
		apiResponse, err := t.provider.Login(ctx)
		if err != nil {
			return true
		}
		t.store.SetAuthToken(apiResponse.AccessToken)
		return false
	}
}

type TaskManager struct {
	duration  time.Duration
	generator *TaskGenerator
	manager   *scheduler.Manager
}

func NewTaskManager(duration time.Duration, generator *TaskGenerator, manager *scheduler.Manager) *TaskManager {
	return &TaskManager{duration: duration, generator: generator, manager: manager}
}

func (t *TaskManager) Run() error {
	ctx := context.Background()
	task := t.generator.Generate()
	task(ctx)
	const taskId = "sitcLoginProviderTask"
	if _, err := t.manager.AddWithDuration(ctx, taskId, task, t.duration); err != nil {
		return err
	}
	return nil
}
