package scheduler

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type TimeParseError struct{}

func (m *TimeParseError) Error() string {
	return "cannot parse time, format should be 14:20 (24 hours: 60 minutes)"
}

type ITimeParser interface {
	ParseHourMinuteString(s string) (time.Duration, error)
}

type TimeParser struct{}

func (t *TimeParser) validate(s string) error {
	match, err := regexp.MatchString(`\d{1,2}:\d{2}`, s)
	if !match {
		return &TimeParseError{}
	}
	if err != nil {
		return &TimeParseError{}
	}
	return nil
}
func (t *TimeParser) scanTimeString(s string) int {
	var timeInt int
	if _, err := fmt.Sscanf(s, `%d`, &timeInt); err != nil {
		return -1
	}
	return timeInt
}
func (t *TimeParser) getHoursToTick(strHours string) (time.Duration, error) {
	hours := t.scanTimeString(strHours)
	if hours == -1 {
		return time.Second, &TimeParseError{}
	}
	return (time.Hour * time.Duration(hours)) - (time.Duration(time.Now().Hour()) * time.Hour), nil
}
func (t *TimeParser) getMinutesToTick(strMinutes string) (time.Duration, error) {
	minutes := t.scanTimeString(strMinutes)
	if minutes == -1 {
		return time.Second, &TimeParseError{}

	}
	return (time.Minute * time.Duration(minutes)) - (time.Duration(time.Now().Minute()) * time.Minute), nil
}

func (t *TimeParser) ParseHourMinuteString(s string) (time.Duration, error) {
	if validateErr := t.validate(s); validateErr != nil {
		return time.Second, validateErr
	}
	splitTime := strings.Split(s, ":")
	strHours := splitTime[0]
	strMinutes := splitTime[1]
	hourDuration, getHoursToTickErr := t.getHoursToTick(strHours)
	if getHoursToTickErr != nil {
		return time.Second, getHoursToTickErr
	}
	minuteDuration, getMinutesToTickErr := t.getMinutesToTick(strMinutes)
	if getMinutesToTickErr != nil {
		return time.Second, getMinutesToTickErr
	}
	outputDurationToTick := hourDuration + minuteDuration
	return time.Hour*24 + outputDurationToTick, nil
}

func NewTimeParser() ITimeParser {
	return &TimeParser{}
}
