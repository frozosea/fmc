package user

type AddContainers struct {
	Numbers []string `json:"numbers"`
}

type DeleteNumbers struct {
	Numbers []int64 `json:"numberIds"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
