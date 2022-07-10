package line

import (
	"bytes"
	"context"
	pb "fmc-newest/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
)

type converter struct{}

func (c *converter) convertControllerResponseToGrpcMessage(result []*Line) *pb.GetAllLinesResponse {
	var allGrpcLines []*pb.Line
	for _, v := range result {
		oneLine := pb.Line{
			LineId:    int64(v.Id),
			Scac:      v.Scac,
			LineName:  v.FullName,
			LineImage: v.ImageUrl,
		}
		allGrpcLines = append(allGrpcLines, &oneLine)
	}
	return &pb.GetAllLinesResponse{Lines: allGrpcLines}
}

type Service struct {
	controller IController
	converter  converter
	pb.UnimplementedLineServiceServer
}

func NewService(controller IController) *Service {
	return &Service{controller: controller, UnimplementedLineServiceServer: pb.UnimplementedLineServiceServer{}, converter: converter{}}
}

func (s *Service) AddLine(stream pb.LineService_AddLineServer) error {
	req, readStreamErr := stream.Recv()
	if readStreamErr != nil {
		return readStreamErr
	}
	imageData := bytes.Buffer{}
	for {
		ctxErr := contextError(stream.Context())
		if ctxErr != nil {
			return ctxErr
		}
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		_, writeByteErr := imageData.Write(req.GetImage())
		if writeByteErr != nil {
			return writeByteErr
		}
	}
	line := WithByteImage{
		BaseLine: BaseLine{
			Scac:     req.GetScac(),
			FullName: req.GetFullName(),
		},
		Image: bytes.NewReader(imageData.Bytes()),
	}
	return s.controller.AddLine(stream.Context(), line)
}
func (s *Service) GetAllLines(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllLinesResponse, error) {
	result, err := s.controller.GetAllLines(ctx)
	if err != nil {
		var allGrpcLines *pb.GetAllLinesResponse
		return allGrpcLines, err
	}
	return s.converter.convertControllerResponseToGrpcMessage(result), nil
}
func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}
