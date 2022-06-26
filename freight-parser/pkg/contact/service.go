package contact

import (
	"context"
	pb "fmc-newest/internal/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type converter struct{}

func (c *converter) convertGrpcMessageToAddContactStruct(addContact *pb.AddContactRequest) *BaseContact {
	return &BaseContact{
		Url:         addContact.Url,
		Email:       addContact.Email,
		AgentName:   addContact.AgentName,
		PhoneNumber: addContact.PhoneNumber,
	}
}

func (c *converter) convertControllerResponseToGrpcMessage(repoContacts []*Contact) *pb.GetAllContactsResponse {
	var allContacts []*pb.Contact
	for _, c := range repoContacts {
		oneContact := pb.Contact{
			Id:          int64(c.Id),
			Url:         c.Url,
			PhoneNumber: c.PhoneNumber,
			AgentName:   c.AgentName,
			Email:       c.Email,
		}
		allContacts = append(allContacts, &oneContact)
	}
	return &pb.GetAllContactsResponse{Contact: allContacts}
}

type service struct {
	controller IController
	converter
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
