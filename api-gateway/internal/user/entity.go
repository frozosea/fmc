package user

type Container struct {
	Id        int64  `json:"id" binding:"required"`
	Number    string `json:"number" binding:"required"`
	IsOnTrack bool   `json:"IsOnTrack" binding:"required"`
}
type AllContainersAndBillNumbers struct {
	Containers  []*Container `json:"containers"`
	BillNumbers []*Container `json:"billNumbers"`
}
