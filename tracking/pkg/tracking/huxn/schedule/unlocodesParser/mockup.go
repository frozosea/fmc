package unlocodesParser

import "context"

type ServiceMockup struct {
}

func NewServiceMockup() *ServiceMockup {
	return &ServiceMockup{}
}
func (s *ServiceMockup) GetUnlocodes(_ context.Context, portLoadUnlocode string) ([]string, error) {
	return []string{portLoadUnlocode}, nil
}
