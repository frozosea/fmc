package city

import (
	"context"
	pb "github.com/frozosea/fmc-pb/freight"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Http struct {
	provider *Service
}

func NewHttp(provider *Service) *Http {
	return &Http{provider: provider}
}

// AddCity
// @Summary Add city by params
// @Security ApiKeyAuth
// @Description Add city by params
// @accept json
// @Param input body CountryWithId true "body"
// @Tags         City
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /city [post]
func (h *Http) AddCity(c *gin.Context) {
	ctx := c.Request.Context()
	var r *CountryWithId
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.provider.AddCity(ctx, r); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
}

// GetAllCities
// @Summary get all cities
// @Security ApiKeyAuth
// @Description get all cities
// @accept json
// @Tags         City
// @Success 200 {object} []City
// @Failure 500 {object} BaseResponse
// @Router /cities [get]
func (h *Http) GetAllCities(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := h.provider.GetAllCities(ctx)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, result)
}

// GetAllCountries
// @Summary get all countries
// @Security ApiKeyAuth
// @Description get all countries
// @accept json
// @Tags         Country
// @Success 200 {object} []Country
// @Failure 500 {object} BaseResponse
// @Router /countries [get]
func (h *Http) GetAllCountries(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := h.provider.GetAllCountries(ctx)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, result)
}

// UpdateCity
// @Summary Add city by params
// @Security ApiKeyAuth
// @Description Add city by params
// @accept json
// @Tags         City
// @Param input body UpdateCity true "body"
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /city [put]
func (h *Http) UpdateCity(c *gin.Context) {
	ctx := c.Request.Context()
	var r UpdateCity
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.provider.UpdateCity(ctx, r.Id, &r.CountryWithId); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
}

// UpdateCountry
// @Summary Add city by params
// @Security ApiKeyAuth
// @Description Add city by params
// @accept json
// @Tags         Country
// @Param input body Country true "body"
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /country [put]
func (h *Http) UpdateCountry(c *gin.Context) {
	ctx := c.Request.Context()
	var r Country
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.provider.UpdateCountry(ctx, r.Id, &r.BaseEntity); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
}

// AddCountry
// @Summary Add country by params
// @Security ApiKeyAuth
// @Description Add country by params
// @accept json
// @Param input body BaseEntity true "body"
// @Tags         City
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /country [post]
func (h *Http) AddCountry(c *gin.Context) {
	ctx := c.Request.Context()
	var r BaseEntity
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.provider.AddCountry(ctx, r); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
}
func (h *Http) deleteWithId(c *gin.Context, fn func(ctx context.Context, id int) error) {
	ctx := c.Request.Context()
	var r Id
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := fn(ctx, r.Id); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
}

// DeleteCountry
// @Summary Add country by params
// @Security ApiKeyAuth
// @Description Add country by params
// @accept json
// @Param 		id  query int  true "body"
// @Tags         Country
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /country [delete]
func (h *Http) DeleteCountry(c *gin.Context) {
	h.deleteWithId(c, func(ctx context.Context, id int) error {
		return h.provider.DeleteCountry(ctx, id)
	})
}

// DeleteCity
// @Summary Add country by params
// @Security ApiKeyAuth
// @Description Add country by params
// @accept json
// @Param id  query int  true "body"
// @Tags         Country
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /city [delete]
func (h *Http) DeleteCity(c *gin.Context) {
	h.deleteWithId(c, func(ctx context.Context, id int) error {
		return h.provider.DeleteCity(ctx, id)
	})
}

type converter struct {
}

func newConverter() *converter {
	return &converter{}
}

func (c *converter) convertResponseToGrpcResponse(cities []*City) *pb.GetAllCitiesResponse {
	var outputCitiesArray []*pb.City
	for _, city := range cities {
		oneGrpcCity := pb.City{
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
		outputCitiesArray = append(outputCitiesArray, &oneGrpcCity)
	}
	return &pb.GetAllCitiesResponse{Cities: outputCitiesArray}
}

type Grpc struct {
	provider *Service
	pb.UnimplementedCityServiceServer
	converter *converter
}

func NewGrpc(provider *Service) *Grpc {
	return &Grpc{provider: provider, UnimplementedCityServiceServer: pb.UnimplementedCityServiceServer{}, converter: newConverter()}
}

func (s *Grpc) GetAllCities(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllCitiesResponse, error) {
	result, err := s.provider.GetAllCities(ctx)
	if err != nil {
		return &pb.GetAllCitiesResponse{}, status.Error(codes.Internal, err.Error())
	}
	return s.converter.convertResponseToGrpcResponse(result), err
}
