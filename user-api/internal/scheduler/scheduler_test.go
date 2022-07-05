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
	}
	p := NewTimeParser()
	for _, v := range testTable {
		_, err := p.ParseHourMinuteString(v.timeStr)
		if v.valid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestTimeParser(t *testing.T) {
	//testTable := []struct {
	//	TimeStr          string
	//	expectedDuration time.Duration
	//}{
	//	{TimeStr: "14:00", expectedDuration: time.Hour*24 + time.Duration(14)*time.Hour},
	//	{TimeStr: "15:00", expectedDuration: time.Hour*24 + time.Duration(15)*time.Hour},
	//	{TimeStr: "16:00", expectedDuration: time.Hour*24 + time.Duration(16)*time.Hour},
	//	{TimeStr: "17:00", expectedDuration: time.Hour*24 + time.Duration(17)*time.Hour},
	//	{TimeStr: "18:00", expectedDuration: time.Hour*24 + time.Duration(18)*time.Hour},
	//	{TimeStr: "19:00", expectedDuration: time.Hour*24 + time.Duration(19)*time.Hour},
	//	{TimeStr: "20:00", expectedDuration: time.Hour*24 + time.Duration(20)*time.Hour},
	//	{TimeStr: "9:00", expectedDuration: time.Hour*24 + time.Duration(9)*time.Hour},
	//	{TimeStr: "8:00", expectedDuration: time.Hour*24 + time.Duration(8)*time.Hour},
	//	{TimeStr: "7:00", expectedDuration: time.Hour*24 + time.Duration(7)*time.Hour},
	//	{TimeStr: "9:05", expectedDuration: time.Hour*24 + time.Duration(9)*time.Hour + (time.Minute * time.Duration(5)) - (time.Duration(time.Now().Minute()) * time.Minute)},
	//	//{TimeStr: "21:00", expectedDuration: time.Hour*24 + time.Duration(21)*time.Hour},
	//	//{TimeStr: "21:00", expectedDuration: time.Hour*24 + time.Duration(14)*time.Hour},
	//	//{TimeStr: "21:00", expectedDuration: time.Hour*24 + time.Duration(14)*time.Hour},
	//	//{TimeStr: "21:00", expectedDuration: time.Hour*24 + time.Duration(14)*time.Hour},
	//	//{TimeStr: "21:00", expectedDuration: time.Hour*24 + time.Duration(14)*time.Hour},
	//}
	p := NewTimeParser()
	_, err := p.ParseHourMinuteString("asfasfasf")
	assert.Error(t, err)
	//for _, v := range testTable {
	//	d, err := timeParser.ParseHourMinuteString(v.TimeStr)
	//	assert.NoError(t, err)
	//	assert.Equal(t, v.expectedDuration, d)
	//}
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
	job, err := store.Save(ctx, TASK_ID, fn, time.Second*1, []interface{}{"asdasd", "asdasd", "asdadas"})
	assert.NoError(t, err)
	assert.Equal(t, job.Id, TASK_ID)
	assert.Equal(t, job.Args, []interface{}{"asdasd", "asdasd", "asdadas"})
	getJob, err := store.Get(ctx, TASK_ID)
	assert.NoError(t, err)
	assert.Equal(t, job, getJob)
}
