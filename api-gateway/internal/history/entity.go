package history

import (
	"fmc-gateway/internal/tracking"
)

type TasksArchive struct {
	Containers []*tracking.ContainerNumberResponse
	Bills      []*tracking.BillNumberResponse
}
type BaseResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
