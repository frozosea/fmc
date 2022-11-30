package container

import (
	"context"
	pb "github.com/frozosea/fmc-proto/freight"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Http struct {
	service *Service
}

func NewHttp(service *Service) *Http {
	return &Http{service: service}
}

// GetAllContainers
// @Summary get all containers
// @Security ApiKeyAuth
// @Tags         Containers
// @Success 200 {object} []Container
// @Failure 500 {object} BaseResponse
// @Router /containers [get]
func (h *Http) GetAllContainers(c *gin.Context) {
	ctx := c.Request.Context()
	r, err := h.service.GetAll(ctx)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, r)
}

// AddContainer
// @Summary add new container
// @Security ApiKeyAuth
// @Tags         Containers
// @Param containerType  query string false  "containerType"
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /container [post]
func (h *Http) AddContainer(c *gin.Context) {
	ctx := c.Request.Context()
	var r *struct {
		ContainerType string `json:"containerType" form:"containerType"`
	}
	if err := c.ShouldBindQuery(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.service.Add(ctx, r.ContainerType); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
}

// DeleteContainer
// @Summary update
// @Security ApiKeyAuth
// @Tags         Containers
// @Param id  query string false  "id"
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /container [delete]
func (h *Http) DeleteContainer(c *gin.Context) {
	ctx := c.Request.Context()
	var r *struct {
		Id int `json:"id" form:"id"`
	}
	if err := c.ShouldBindQuery(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.service.Delete(ctx, r.Id); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
}

// UpdateContainer
// @Summary update
// @Security ApiKeyAuth
// @Tags         Containers
// @Param input body Container true "body"
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /container [put]
func (h *Http) UpdateContainer(c *gin.Context) {
	ctx := c.Request.Context()
	var r *Container
	if err := c.ShouldBindQuery(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.service.Update(ctx, r.Id, r.Type); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
}

type Grpc struct {
	service   *Service
	converter *converter
	pb.UnimplementedContainersServiceServer
}

func NewGrpc(service *Service) *Grpc {
	return &Grpc{service: service, converter: newConverter(), UnimplementedContainersServiceServer: pb.UnimplementedContainersServiceServer{}}
}

func (g *Grpc) GetAllContainers(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllContainersResponse, error) {
	allContainer, err := g.service.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return g.converter.convertArrayToGrpc(allContainer), nil
}

type converter struct {
}

func newConverter() *converter {
	return &converter{}
}

func (c *converter) convertArrayToGrpc(containers []*Container) *pb.GetAllContainersResponse {
	var ar []*pb.Container
	for _, v := range containers {
		ar = append(ar, &pb.Container{
			ContainerType:   v.Type,
			ContainerTypeId: int64(v.Id),
		})
	}
	return &pb.GetAllContainersResponse{Containers: ar}
}
