package company

import (
	"context"
	pb "github.com/frozosea/fmc-pb/freight"

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

// GetAll
// @Summary get all contacts
// @Security ApiKeyAuth
// @accept json
// @Tags         Company
// @Success 200 {object} []Company
// @Failure 500 {object} BaseResponse
// @Router /companies [get]
func (h *Http) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := h.service.GetAllContacts(ctx)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, result)
}

// Add
// @Summary add new company
// @Security ApiKeyAuth
// @accept json
// @Tags         Company
// @Param input body BaseCompany true "body"
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /company [post]
func (h *Http) Add(c *gin.Context) {
	ctx := c.Request.Context()
	var r BaseCompany
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.service.AddContact(ctx, r); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": false, "error": nil})

}

// UpdateCompany
// @Summary update company
// @Security ApiKeyAuth
// @accept json
// @Tags         Company
// @Param input body UpdateCompany true "body"
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /company [put]
func (h *Http) UpdateCompany(c *gin.Context) {
	ctx := c.Request.Context()
	var r UpdateCompany
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(401, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	if err := h.service.Update(ctx, r.Id, &r.BaseCompany); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": false, "error": nil})

}

// Delete
// @Summary delete company by id
// @Security ApiKeyAuth
// @accept json
// @Tags         Company
// @Param id  query int  true "body"
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /company [delete]
func (h *Http) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	var r struct {
		Id int `json:"id" form:"id"`
	}
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

type converter struct{}

func (a *converter) ToGrpc(repoContacts []*Company) *pb.GetAllCompainesResponse {
	var allContacts []*pb.Company
	for _, v := range repoContacts {
		oneContact := pb.Company{
			Id:          int64(v.Id),
			Url:         v.Url,
			PhoneNumber: v.PhoneNumber,
			AgentName:   v.Name,
			Email:       v.Email,
		}
		allContacts = append(allContacts, &oneContact)
	}
	return &pb.GetAllCompainesResponse{Contact: allContacts}
}

type Grpc struct {
	provider *Service
	*converter
	pb.UnimplementedCompanyServiceServer
}

func (s *Grpc) GetAllContacts(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllCompainesResponse, error) {
	result, getContactsErr := s.provider.GetAllContacts(ctx)
	if getContactsErr != nil {
		return &pb.GetAllCompainesResponse{}, status.Error(codes.Internal, getContactsErr.Error())
	}
	return s.ToGrpc(result), nil
}
func NewGrpc(provider *Service) *Grpc {
	return &Grpc{provider: provider, converter: &converter{}, UnimplementedCompanyServiceServer: pb.UnimplementedCompanyServiceServer{}}
}
