package domain

import (
	"time"
)

type BaseTrackReq struct {
	Numbers             []string
	UserId              int64
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

type GetInfoAboutTrackResponse struct {
	Number               string                `json:"number"`
	IsContainer          bool                  `json:"isContainer"`
	IsOnTrack            bool                  `json:"isOnTrack"`
	ScheduleTrackingInfo *ScheduleTrackingInfo `json:"scheduleTrackingInfo"`
}

type ScheduleTrackingInfo struct {
	Time    string   `json:"time"`
	Subject string   `json:"subject"`
	Emails  []string `json:"emails"`
}

type TrackingTask struct {
	Number              string
	UserId              int64
	Time                string
	Emails              []string
	IsContainer         bool
	EmailMessageSubject string
}
