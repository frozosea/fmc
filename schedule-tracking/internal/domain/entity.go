package domain

import (
	"time"
)

type BaseTrackReq struct {
	Number              string
	UserId              int64
	Country             string
	Time                string
	Emails              []string
	EmailMessageSubject string
}
type BaseAddOnTrackResponse struct {
	success     bool
	number      string
	nextRunTime time.Time
}
type AddOnTrackResponse struct {
	result         []*BaseAddOnTrackResponse
	alreadyOnTrack []string
}
type AddEmailRequest struct {
	numbers []string
	emails  []string
	userId  int64
}
type DeleteEmailFromTrack struct {
	number string
	email  string
	userId int64
}
type GetInfoAboutTrackResponse struct {
	number              string
	emails              []interface{}
	nextRunTime         time.Time
	emailMessageSubject string
}

type TrackingTask struct {
	Number              string
	UserId              int64
	Country             string
	Time                string
	Emails              []string
	IsContainer         bool
	EmailMessageSubject string
}
