package line

import (
	"bytes"
	"context"
	"fmc-newest/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
)

type converter struct{}

func (c *converter) convertControllerResponseToGrpcMessage(result []*Line) *___.GetAllLinesResponse {
	var allGrpcLines []*___.Line
	for _, v := range result {
		oneLine := ___.Line{
			LineId:    int64(v.Id),
			Scac:      v.Scac,
			LineName:  v.FullName,
			LineImage: v.ImageUrl,
		}
		allGrpcLines = append(allGrpcLines, &oneLine)
	}
	return &___.GetAllLinesResponse{Lines: allGrpcLines}
}

type service struct {
	controller IController
	converter  converter
	___.UnimplementedLineServiceServer
}

func (s *service) AddLine(stream ___.LineService_AddLineServer) error {
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
func (s *service) GetAllLines(ctx context.Context, _ *emptypb.Empty) (*___.GetAllLinesResponse, error) {
	result, err := s.controller.GetAllLines(ctx)
	if err != nil {
		var allGrpcLines *___.GetAllLinesResponse
		return allGrpcLines, err
	}
	return s.converter.convertControllerResponseToGrpcMessage(result), nil
}
func newService(controller IController) *service {
	return &service{controller: controller, converter: converter{}}
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
