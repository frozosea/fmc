package scac

import "errors"

type Service struct {
	containerLines []*WithFullName
	billLines      []*WithFullName
}

func NewService(containerLines []*WithFullName, billLines []*WithFullName) *Service {
	return &Service{containerLines: containerLines, billLines: billLines}
}

func (s *Service) GetContainerLines() (*WithFullNameList, error) {
	if len(s.containerLines) == 0 {
		return nil, errors.New("cannot resolve container lines")
	}
	return &WithFullNameList{list: s.containerLines}, nil
}
func (s *Service) GetBillLines() (*WithFullNameList, error) {
	if len(s.billLines) == 0 {
		return nil, errors.New("cannot resolve bill lines")
	}
	return &WithFullNameList{list: s.billLines}, nil
}
