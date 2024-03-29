package scheduler

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

type TimeParseError struct{}

func (m *TimeParseError) Error() string {
	return "cannot parse time, format should be 14:20 (24 hours: 60 minutes)"
}

type ITimeParser interface {
	Parse(s string) (time.Duration, error)
}

type TimeParser struct {
	timezone string
}

func NewTimeParser(timezone string) *TimeParser {
	return &TimeParser{timezone: timezone}
}

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

func (t TimeParser) Parse(s string) (time.Duration, error) {
	if validateErr := t.validate(s); validateErr != nil {
		return time.Second, validateErr
	}
	now := time.Now()
	splitTime := strings.Split(s, ":")
	strHour, strMinute := splitTime[0], splitTime[1]
	hour, err := strconv.Atoi(strHour)
	if err != nil {
		return time.Second, err
	}
	minute, err := strconv.Atoi(strMinute)
	if err != nil {
		return time.Second, err
	}
	var day int
	if hour >= now.Hour() && minute > now.Minute() {
		day = now.Day()
	} else {
		day = now.Day() + 1
	}
	tz, err := time.LoadLocation(t.timezone)
	if err != nil {
		tz = time.Local
	}
	date := time.Date(now.Year(), now.Month(), day, hour, minute, 0, 0, tz)
	return date.Sub(now), nil
}
