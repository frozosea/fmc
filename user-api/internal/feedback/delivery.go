package feedback

import (
	"context"
	pb "github.com/frozosea/fmc-proto/user"
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

func (h *Http) GetByEmail(c *gin.Context) {
	var r struct {
		Email string `json:"email" form:"email"`
	}
	if err := c.ShouldBindQuery(&r); err != nil {
		c.JSON(402, gin.H{"success": false, "error": "bad request"})
		return
	}
	fbs, err := h.service.GetByEmail(c.Request.Context(), r.Email)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, fbs)
}
func (h *Http) GetAll(c *gin.Context) {
	fbs, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, fbs)
}

type Grpc struct {
	service *Service
	pb.UnimplementedUserFeedbackServer
}

func NewGrpc(service *Service) *Grpc {
	return &Grpc{service: service, UnimplementedUserFeedbackServer: pb.UnimplementedUserFeedbackServer{}}
}

func (g *Grpc) AddFeedback(ctx context.Context, r *pb.AddFeedbackRequest) (*emptypb.Empty, error) {
	if err := g.service.Add(ctx, FromGrpc(r)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
