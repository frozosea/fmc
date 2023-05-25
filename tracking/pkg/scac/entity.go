package scac

import pb "github.com/frozosea/fmc-pb/v2/tracking"

type WithFullName struct {
	Scac     string `json:"scac"`
	Fullname string `json:"fullname"`
}

func (w *WithFullName) ToGRPC() *pb.Scac {
	return &pb.Scac{
		Scac:     w.Scac,
		Fullname: w.Fullname,
	}
}

type WithFullNameList struct {
	list []*WithFullName
}

func (w *WithFullNameList) ToGRPC() *pb.GetAllScacResponse {
	var allLines []*pb.Scac
	for _, v := range w.list {
		allLines = append(allLines, v.ToGRPC())
	}
	return &pb.GetAllScacResponse{Data: allLines}
}
