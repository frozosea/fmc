package user

type AddContainers struct {
	Numbers []string `json:"numbers"`
}

type DeleteNumbers struct {
	Numbers []string `json:"numbers"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
type AddFeedback struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}
