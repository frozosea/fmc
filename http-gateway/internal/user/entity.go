package user

type ScheduleTrackingInfoObject struct {
	Time    string   `json:"time"`
	Emails  []string `json:"emails"`
	Subject string   `json:"subject"`
}
type Container struct {
	Number               string                      `json:"number" binding:"required"`
	IsOnTrack            bool                        `json:"isOnTrack" binding:"required"`
	IsContainer          bool                        `json:"isContainer"`
	ScheduleTrackingInfo *ScheduleTrackingInfoObject `json:"scheduleTrackingInfo"`
}
type AllContainersAndBillNumbers struct {
	Containers  []*Container `json:"containers"`
	BillNumbers []*Container `json:"billNumbers"`
}
