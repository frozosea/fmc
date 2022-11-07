package freight

import (
	"context"
	"freight_service/internal/city"
	pb "freight_service/pkg/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Http struct {
	service *Service
}

func NewHttp(service *Service) *Http {
	return &Http{service: service}
}

// GetFreights
// @Summary get all freights
// @Security ApiKeyAuth
// @accept json
// @Tags         Freight
// @Param input body AddFreight true "body"
// @Success 200 {object} []BaseFreight
// @Failure 500 {object} BaseResponse
// @Router /freights [get]
func (h *Http) GetFreights(c *gin.Context) {
	result, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, result)
}

// AddFreight
// @Summary add new freight
// @Security ApiKeyAuth
// @accept json
// @Tags         Freight
// @Param input body AddFreight true "body"
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /freight [post]
func (h *Http) AddFreight(c *gin.Context) {
	ctx := c.Request.Context()
	var r AddFreight
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.service.AddFreight(ctx, r); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
}

// UpdateFreight
// @Summary update freight
// @Security ApiKeyAuth
// @accept json
// @Tags         Freight
// @Param input body UpdateFreight true "body"
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /freight [put]
func (h *Http) UpdateFreight(c *gin.Context) {
	ctx := c.Request.Context()
	var r *UpdateFreight
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.service.Update(ctx, r.Id, &r.AddFreight); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
}

// DeleteFreight
// @Summary delete freight
// @Security ApiKeyAuth
// @Tags         Freight
// @Param id  query     string     false  "id"
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /freight [delete]
func (h *Http) DeleteFreight(c *gin.Context) {
	ctx := c.Request.Context()
	var r *DeleteFreight
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.service.Delete(ctx, r.Id); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
}

type converter struct {
}

func (c *converter) cityToGrpc(city city.City) *pb.City {
	return &pb.City{
		City: &pb.LocEntity{
			Id:     int64(city.Id),
			RuName: city.RuFullName,
			EnName: city.EnFullName,
		},
		Country: &pb.LocEntity{
			Id:     int64(city.Country.Id),
			RuName: city.Country.RuFullName,
			EnName: city.Country.EnFullName,
		},
	}
}
func (c *converter) ToGrpc(freights []BaseFreight) *pb.GetFreightsResponseList {
	var outpurArr []*pb.GetFreightResponse
	for _, freight := range freights {
		outpurArr = append(outpurArr, freight.ToGrpc())
	}
	return &pb.GetFreightsResponseList{MultiResponse: outpurArr}
}
func (c *converter) fromGrpc(r *pb.GetFreightRequest) GetFreight {
	return GetFreight{
		FromCityId:      r.FromCityId,
		ToCityId:        r.ToCityId,
		ContainerTypeId: r.ContainerTypeId,
		Limit:           r.Limit,
	}
}

type Grpc struct {
	service *Service
	*converter
	pb.UnimplementedFreightServiceServer
}

func NewGrpc(service *Service) *Grpc {
	return &Grpc{service: service, converter: &converter{}, UnimplementedFreightServiceServer: pb.UnimplementedFreightServiceServer{}}
}
func (g *Grpc) GetFreights(ctx context.Context, r *pb.GetFreightRequest) (*pb.GetFreightsResponseList, error) {
	freights, err := g.service.getFreights(ctx, g.converter.fromGrpc(r))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return g.converter.ToGrpc(freights), nil
}
