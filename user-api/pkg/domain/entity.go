package domain

type Container struct {
	Id        int64  `json:"id"`
	Number    string `json:"number"`
	IsOnTrack bool   `json:"is_on_track"`
}
type AllContainersAndBillNumbers struct {
	Containers  []*Container `json:"containers"`
	BillNumbers []*Container `json:"bill_numbers"`
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserWithId struct {
	Id int
	User
}
