package container

import pb "freight_service/pkg/proto"

type Container struct {
	Id   int
	Type string
}

func (c *Container) ToGrpc() *pb.Container {
	return &pb.Container{
		ContainerType:   c.Type,
		ContainerTypeId: int64(c.Id),
	}
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
