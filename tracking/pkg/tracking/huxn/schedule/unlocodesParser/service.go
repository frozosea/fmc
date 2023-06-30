package unlocodesParser

import (
	"context"
	"golang_tracking/pkg/tracking/util/requests"
)

type IService interface {
	GetUnlocodes(ctx context.Context, portLoadUnlocode string) ([]string, error)
}

type Service struct {
	request *Request
	parser  *Parser
	data    map[string][]string
}

func NewService(http requests.IHttp) *Service {
	return &Service{request: NewRequest(http), parser: NewParser(), data: map[string][]string{
		"RUVYP": {"CNTJN", "RUVVO"},
		"CNTAO": {"CNTAC", "CNTJN", "KRPUS", "RUVVO", "RUVYP"},
		"CNTAC": {"CNTJN", "KRPUS", "RUVVO", "RUVYP"},
		"KRPUS": {"CNTJN", "RUVVO", "RUVYP"},
		"CNTJN": {"CNTAO", "KRPUS", "RUVVO"},
		"RUVVO": {"CNTAC", "CNTAO", "CNTJN"},
	}}
}

func (s *Service) GetUnlocodes(_ context.Context, portLoadUnlocode string) ([]string, error) {
	return s.data[portLoadUnlocode], nil
}
