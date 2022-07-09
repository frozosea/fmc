package contact

import (
	"context"
	pb "fmc-newest/pkg/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type adapter struct{}

func (a *adapter) convertGrpcMessageToAddContactStruct(addContact *pb.AddContactRequest) *BaseContact {
	return &BaseContact{
		Url:         addContact.Url,
		Email:       addContact.Email,
		AgentName:   addContact.AgentName,
		PhoneNumber: addContact.PhoneNumber,
	}
}

func (a *adapter) convertControllerResponseToGrpcMessage(repoContacts []*Contact) *pb.GetAllContactsResponse {
	var allContacts []*pb.Contact
	for _, v := range repoContacts {
		oneContact := pb.Contact{
			Id:          int64(v.Id),
			Url:         v.Url,
			PhoneNumber: v.PhoneNumber,
			AgentName:   v.AgentName,
			Email:       v.Email,
		}
		allContacts = append(allContacts, &oneContact)
	}
	return &pb.GetAllContactsResponse{Contact: allContacts}
}

type service struct {
	controller IController
	adapter
	pb.UnimplementedContactServiceServer
}

func (s *service) AddContact(ctx context.Context, addContact *pb.AddContactRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.controller.AddContact(ctx, *s.convertGrpcMessageToAddContactStruct(addContact))
}
func (s *service) GetAllContacts(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllContactsResponse, error) {
	result, getContactsErr := s.controller.GetAllContacts(ctx)
	return s.convertControllerResponseToGrpcMessage(result), getContactsErr
}
func NewService(controller IController) *service {
	return &service{controller: controller}
}
