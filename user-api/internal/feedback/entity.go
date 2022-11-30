package feedback

import (
	pb "github.com/frozosea/fmc-pb/user"
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
