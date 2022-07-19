package domain

import (
	"time"
)

type BaseTrackReq struct {
	numbers []string
	userId  int64
	country string
	time    string
	emails  []string
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
