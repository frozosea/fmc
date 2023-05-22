package domain

import "time"

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
	BillNumbers []*Container `json:"billNumbers"`
}
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisterUser struct {
	Email                  string `json:"email"`
	Username               string `json:"username"`
	Password               string `json:"password"`
	CompanyFullName        string `json:"companyFullName"`
	CompanyAbbreviatedName string `json:"CompanyAbbreviatedName"`
	INN                    string `json:"INN"`
	OGRN                   string `json:"OGRN"`
	LegalAddress           string `json:"legalAddress"`
	PostAddress            string `json:"postAddress"`
	WorkEmail              string `json:"workEmail"`
}
type UserWithId struct {
	Id int
	User
}

type Transaction struct {
	ID        int       `json:"ID,omitempty"`
	UserID    int       `json:"UserID,omitempty"`
	Value     float64   `json:"Value,omitempty"`
	Type      int       `json:"Type"`
	TimeStamp time.Time `json:"TimeStamp"`
}

type Balance struct {
	UserId       int            `json:"userId"`
	Value        float64        `json:"value"`
	Transactions []*Transaction `json:"transactions"`
}

type UserWithBalance struct {
	UserWithId
	Balance
}
