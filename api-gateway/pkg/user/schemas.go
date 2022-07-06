package user

type AddContainers struct {
	Numbers []string `json:"numbers" validate:"min=10,max=28,regexp=[a-zA-Z]{3,}\d{5,}"`
}

type DeleteNumbers struct {
	numberIds []int64 `json:"numbers" validate:"min=10,max=28,regexp=[a-zA-Z]{3,}\d{5,}"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}