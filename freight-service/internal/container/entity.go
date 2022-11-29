package container

import pb "github.com/frozosea/fmc-proto/freight"

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
