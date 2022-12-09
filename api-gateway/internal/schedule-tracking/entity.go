package schedule_tracking

import (
	pb "github.com/frozosea/fmc-pb/schedule-tracking"
	_ "gopkg.in/validator.v2"
)

type AddOnTrackRequest struct {
	Numbers             []string `json:"numbers" binding:"required"`
	Emails              []string `json:"emails" binding:"required"  validate:"min=3,max=500,regexp=^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$"`
	Time                string   `json:"time" binding:"required"  validate:"min=4,max=5,regexp=\d{1,2}:\d{2}"`
	EmailMessageSubject string   `json:"emailSubject,omitempty" binding:"required"  validate:""`
}

func (a *AddOnTrackRequest) ToGrpc(userId int) *pb.AddOnTrackRequest {
	return &pb.AddOnTrackRequest{
		UserId:              int64(userId),
		Numbers:             a.Numbers,
		Emails:              a.Emails,
		EmailMessageSubject: a.EmailMessageSubject,
		Time:                a.Time,
	}
}

func (a *AddOnTrackRequest) FromGrpc(r *pb.AddOnTrackRequest) *AddOnTrackRequest {
	return &AddOnTrackRequest{
		Numbers:             r.GetNumbers(),
		Emails:              r.GetEmails(),
		Time:                r.GetTime(),
		EmailMessageSubject: r.GetEmailMessageSubject(),
	}
}

type BaseAddOnTrackResponse struct {
	Success     bool   `json:"success"`
	Number      string `json:"number"`
	NextRunTime int64  `json:"nextRunTime"`
}

func (b *BaseAddOnTrackResponse) ToGrpc() *pb.BaseAddOnTrackResponse {
	return &pb.BaseAddOnTrackResponse{
		Success:     b.Success,
		Number:      b.Number,
		NextRunTime: b.NextRunTime,
	}
}
func (b *BaseAddOnTrackResponse) FromGrpc(r *pb.BaseAddOnTrackResponse) *BaseAddOnTrackResponse {
	return &BaseAddOnTrackResponse{
		Success:     r.GetSuccess(),
		Number:      r.GetNumber(),
		NextRunTime: r.GetNextRunTime(),
	}
}

type AddOnTrackResponse struct {
	Result         []BaseAddOnTrackResponse `json:"result"`
	AlreadyOnTrack []string                 `json:"alreadyOnTrack"`
}

func (a *AddOnTrackResponse) ToGrpc() *pb.AddOnTrackResponse {
	var result []*pb.BaseAddOnTrackResponse
	for _, v := range a.Result {
		result = append(result, v.ToGrpc())
	}
	return &pb.AddOnTrackResponse{
		BaseResponse:   result,
		AlreadyOnTrack: a.AlreadyOnTrack,
	}
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
	Number               string                `json:"number" validate:"min=10,max=28,regexp=[a-zA-Z]{3,}\d{5,}"`
	IsContainer          bool                  `json:"isContainer"`
	IsOnTrack            bool                  `json:"isOnTrack"`
	ScheduleTrackingInfo *ScheduleTrackingInfo `json:"scheduleTrackingInfo"`
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
type ScheduleTrackingInfo struct {
	Time    string   `json:"time"`
	Subject string   `json:"subject"`
	Emails  []string `json:"emails"`
}
