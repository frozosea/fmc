package domain

type Container struct {
	Number               string                      `json:"number"`
	IsOnTrack            bool                        `json:"IsOnTrack"`
	IsContainer          bool                        `json:"isContainer"`
	ScheduleTrackingInfo *ScheduleTrackingInfoObject `json:"scheduleTrackingInfo"`
}
type ScheduleTrackingInfoObject struct {
	Emails  []string `json:"emails"`
	Subject string   `json:"subject"`
	Time    string   `json:"time"`
}
type AllContainersAndBillNumbers struct {
	Containers  []*Container `json:"containers"`
	BillNumbers []*Container `json:"bill_numbers"`
}
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisterUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserWithId struct {
	Id int
	User
}
