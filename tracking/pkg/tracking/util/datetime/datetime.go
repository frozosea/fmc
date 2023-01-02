package datetime

import (
	"github.com/archsh/timefmt"
	"time"
)

type IDatetime interface {
	Strftime(t time.Time, format string) (string, error)
	Strptime(value string, format string) (time.Time, error)
}

func (d *Datetime) Strftime(t time.Time, format string) (string, error) {
	return timefmt.Strftime(t, format)
}

func (d *Datetime) Strptime(value string, format string) (time.Time, error) {
	return timefmt.Strptime(value, format)
}

type Datetime struct {
}

func NewDatetime() *Datetime {
	return &Datetime{}
}
