package login_provider

import (
	"context"
	scheduler "github.com/frozosea/scheduler/pkg"
	log "github.com/sirupsen/logrus"
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
	return func(ctx context.Context) {
		apiResponse, err := t.provider.Login(ctx)
		if err != nil {
			log.Println(err)
			return
		}
		t.store.SetAuthToken(apiResponse.AccessToken)
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
	task := t.generator.Generate()
	const taskId = "sitcLoginProviderTask"
	ctx := context.Background()
	task(ctx)
	if _, err := t.manager.AddWithDuration(ctx, taskId, task, t.duration); err != nil {
		return err
	}
	return nil
}
