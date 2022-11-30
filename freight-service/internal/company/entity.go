package company

import pb "github.com/frozosea/fmc-pb/freight"

type BaseCompany struct {
	Url         string
	Email       string
	Name        string
	PhoneNumber string
}

type Company struct {
	Id int
	*BaseCompany
}

func (c *Company) ToGrpc() *pb.Company {
	return &pb.Company{
		Id:          int64(c.Id),
		Url:         c.Url,
		PhoneNumber: c.PhoneNumber,
		AgentName:   c.Name,
		Email:       c.Email,
	}
}

type UpdateCompany struct {
	Id int `json:"id"`
	BaseCompany
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
