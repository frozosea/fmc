package scheduler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeParserValidator(t *testing.T) {
	testTable := []struct {
		timeStr string
		valid   bool
	}{
		{"14:07", true},
		{"asdasdads", false},
		{"1:02", true},
		{"13:1", false},
		{"13:1,", false},
		{"1,3:1", false},
		{"13,:1", false},
		{"13:1", false},
		{"13:1", false},
		{"1,3:1", false},
		{"13:01", true},
		{"13:01", true},
		{"8:14", true},
		{"00:14", true},
	}
	p := NewTimeParser()
	for _, v := range testTable {
		_, err := p.Parse(v.timeStr)
		if v.valid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestJobStores(t *testing.T) {
	const TASK_ID = "testTaskId"
	fn := func(ctx context.Context, args ...interface{}) ShouldBeCancelled {
		for _, v := range args {
			assert.NotEmpty(t, v)
		}
		return true
	}
	store := NewMemoryJobStore()
	ctx := context.Background()
	job, err := store.Save(ctx, TASK_ID, fn, time.Second*1, []interface{}{"asdasd", "asdasd", "asdadas"}, "14:20")
	assert.NoError(t, err)
	assert.Equal(t, job.Id, TASK_ID)
	assert.Equal(t, job.Args, []interface{}{"asdasd", "asdasd", "asdadas"})
	getJob, err := store.Get(ctx, TASK_ID)
	assert.NoError(t, err)
	assert.Equal(t, job, getJob)
}
