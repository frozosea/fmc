package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"user-api/pkg/domain"
)

type RespositoryMoch struct {
}

func (r *RespositoryMoch) AddContainerToAccount(ctx context.Context, userId int, containers []string) error {
	return nil
}

func (r *RespositoryMoch) AddBillNumberToAccount(ctx context.Context, userId int, containers []string) error {
	return nil

}

func (r *RespositoryMoch) DeleteContainersFromAccount(ctx context.Context, userId int, numberIds []int64) error {
	return nil

}

func (r *RespositoryMoch) DeleteBillNumbersFromAccount(ctx context.Context, userId int, numberIds []int64) error {
	return nil

}
func (r *RespositoryMoch) GetAllContainersAndBillNumbers(ctx context.Context, userId int) (*domain.AllContainersAndBillNumbers, error) {
	time.Sleep(time.Second * 2)
	var containers []*domain.Container
	for i := 0; i < 5; i++ {
		containers = append(containers, &domain.Container{
			Number:    fmt.Sprintf("FESO%d", i*1000),
			IsOnTrack: false,
		})
	}
	var billNumbers []*domain.Container
	for i := 0; i < 7; i++ {
		billNumbers = append(billNumbers, &domain.Container{
			Number:    fmt.Sprintf("SKLU%d", i*100000),
			IsOnTrack: true,
		})
	}
	return &domain.AllContainersAndBillNumbers{Containers: containers, BillNumbers: billNumbers}, nil
}

func getFromCache() *domain.AllContainersAndBillNumbers {
	var containers []*domain.Container
	for i := 0; i < 10; i++ {
		containers = append(containers, &domain.Container{
			Number:    "FESO2219270",
			IsOnTrack: false,
		})
	}
	var billNumbers []*domain.Container
	for i := 0; i < 7; i++ {
		billNumbers = append(billNumbers, &domain.Container{
			Number:    "SNKO101220501450",
			IsOnTrack: true,
		})
	}
	return &domain.AllContainersAndBillNumbers{Containers: containers, BillNumbers: billNumbers}
}

type cacheMoch struct{}

func (c *cacheMoch) Get(ctx context.Context, key string, dest interface{}) error {
	j, err := json.Marshal(getFromCache())
	if err != nil {
		return err
	}
	return json.Unmarshal(j, &dest)
}
func (c *cacheMoch) Set(ctx context.Context, key string, value interface{}) error {
	return nil
}
func (c *cacheMoch) Del(ctx context.Context, key string) error {
	return nil
}

type loggerMoch struct{}

func (l *loggerMoch) InfoLog(logString string) {}

func (l *loggerMoch) ExceptionLog(logString string) {}

func (l *loggerMoch) WarningLog(logString string) {}

func (l *loggerMoch) FatalLog(logString string) {}

func TestControllerWithMochs(t *testing.T) {
	c := NewProvider(&RespositoryMoch{}, &loggerMoch{}, &cacheMoch{})
	ctx := context.Background()
	res, err := c.GetAllContainers(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, getFromCache(), res)
}
