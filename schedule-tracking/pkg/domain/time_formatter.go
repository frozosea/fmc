package domain

import "time"

type ITimeFormatter interface {
	Convert(timestamp time.Time) string
}

type TimeFormatter struct {
	timeFormat string
}

func (t *TimeFormatter) Convert(timestamp time.Time) string {
	return timestamp.Format(t.timeFormat)
}
func NewTimeFormatter(timeFormat string) *TimeFormatter {
	return &TimeFormatter{timeFormat: timeFormat}
}
