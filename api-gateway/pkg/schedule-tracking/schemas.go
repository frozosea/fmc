package schedule_tracking

import (
	_ "gopkg.in/validator.v2"
)

type AddOnTrackRequest struct {
	Number       string   `json:"number" binding:"required" "min=10,max=28,regexp=[a-zA-Z]{3,}\d{5,}"`
	Emails       []string `json:"emails" binding:"required"  validate:"min=3,max=500,regexp=^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$"`
	Time         string   `json:"time" binding:"required"  validate:"min=4,max=5,regexp=\d{1,2}:\d{2}"`
	EmailMessage string   `json:"email_subject" binding:"required"  validate:""`
}
type BaseAddOnTrackResponse struct {
	Success     bool   `json:"success"`
	Number      string `json:"number"`
	NextRunTime int64  `json:"nextRunTime"`
}
type AddOnTrackResponse struct {
	Result         []BaseAddOnTrackResponse `json:"result"`
	AlreadyOnTrack []string                 `json:"alreadyOnTrack"`
}
type UpdateTrackingTimeRequest struct {
	Numbers []string `json:"numbers" validate:"min=10,max=28,regexp=[a-zA-Z]{4}\d{7}"`
	NewTime string   `json:"newTime" validate:"min=4,max=5,regexp=\d{1,2}:\d{2}`
	userId  int64
}

type AddEmailRequest struct {
	Numbers []string `json:"numbers" validate:"min=10,max=28,regexp=[a-zA-Z]{3,}\d{5,}"`
	Emails  []string `json:"emails" validate:"min=3,max=500,regexp=^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$"`
	userId  int64
}

type DeleteEmailFromTrackRequest struct {
	Number string `json:"number" validate:"min=10,max=28,regexp=[a-zA-Z]{3,}\d{5,}"`
	Email  string `json:"email" validate:"min=3,max=500,regexp=^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$"`
	userId int64
}

type DeleteFromTrackRequest struct {
	Numbers []string `json:"numbers" validate:"min=10,max=28,regexp=[a-zA-Z]{3,}\d{5,}"`
	userId  int64
}

type GetInfoAboutTrackRequest struct {
	Number string `json:"number" validate:"min=10,max=28,regexp=[a-zA-Z]{3,}\d{5,}" form:"number"`
	userId int64
}
type GetInfoAboutTrackResponse struct {
	Number      string   `json:"number" validate:"min=10,max=28,regexp=[a-zA-Z]{3,}\d{5,}"`
	Emails      []string `json:"emails" validate:"min=3,max=500,regexp=^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$"`
	NextRunTime int64    `json:"nextRunTime"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type TimeZoneResponse struct {
	TimeZone string `json:"timeZone"`
}

type ChangeEmailMessageSubjectRequest struct {
	Number     string
	userId     int64
	NewSubject string
}
