package feedback

import (
	pb "user-api/pkg/proto"
)

type Feedback struct {
	Email   string
	Message string
}

func FromGrpc(r *pb.AddFeedbackRequest) *Feedback {
	return &Feedback{
		Email:   r.GetEmail(),
		Message: r.GetMessage(),
	}
}
