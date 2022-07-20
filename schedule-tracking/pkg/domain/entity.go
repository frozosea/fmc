package domain

import (
	"time"
)

type BaseTrackReq struct {
	Numbers []string
	UserId  int64
	Country string
	Time    string
	Emails  []string
}

type TrackByBillNoReq struct {
	BaseTrackReq
}
type TrackByContainerNoReq struct {
	BaseTrackReq
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
}
type DeleteEmailFromTrack struct {
	number string
	email  string
}
type GetInfoAboutTrackResponse struct {
	number      string
	emails      []interface{}
	nextRunTime time.Time
}

type TrackingTask struct {
	Number      string
	UserId      int64
	Country     string
	Time        string
	Emails      []string
	IsContainer bool
}
