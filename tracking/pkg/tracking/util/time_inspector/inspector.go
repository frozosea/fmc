package time_inspector

import (
	"errors"
	"math"
	"time"
)

type ITimeInspector interface {
	CheckInfoAboutMovingExpires(date time.Time) (bool, error)
}

type TimeInspector struct {
}

func New() *TimeInspector {
	return &TimeInspector{}
}
func (t *TimeInspector) checkDateExpires(date time.Time) bool {
	t1 := time.Now()
	timeDiff := t1.Sub(date)
	monthDiff := math.Round(timeDiff.Hours() * 0.001389)
	if monthDiff > 3 {
		return false
	}
	return true
}
func (t *TimeInspector) CheckInfoAboutMovingExpires(date time.Time) (bool, error) {
	t1 := time.Now()
	timeDiff := t1.Sub(date)
	monthDiff := math.Round(timeDiff.Hours() * 0.001389)
	if monthDiff > 3 {
		return false, errors.New("date is bigger then 3 month")
	}
	return true, nil
}
