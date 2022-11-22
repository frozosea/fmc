package domain

import (
	"time"
)

type BaseTrackReq struct {
	Numbers             []string
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
	Number               string                `json:"number" validate:"min=10,max=28,regexp=[a-zA-Z]{3,}\d{5,}"`
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
	Country             string
	Time                string
	Emails              []string
	IsContainer         bool
	EmailMessageSubject string
}
