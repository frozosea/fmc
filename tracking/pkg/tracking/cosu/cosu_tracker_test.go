package cosu

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang_tracking/pkg/tracking/util/datetime"
	"os"
	"testing"
	"time"
)

type RequestMockUp struct{}

func NewRequestMockUp() *RequestMockUp {
	return &RequestMockUp{}
}

func (RequestMockUp) GetInfoAboutMovingRawResponse(_ context.Context, _ string) (*ApiResponseSchema, error) {
	var s *ApiResponseSchema
	infoAboutMovingResponseExample, err := os.ReadFile("test_data/exampleApiResponse.json")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(infoAboutMovingResponseExample, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func (RequestMockUp) GetEtaRawResponse(_ context.Context, _ string) (*EtaApiResponseSchema, error) {
	var s *EtaApiResponseSchema
	etaExampleApiResp, err := os.ReadFile("test_data/etaExampleApiResponse.json")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(etaExampleApiResp, &s); err != nil {
		return nil, err
	}
	return s, nil
}

var requestMockUp = NewRequestMockUp()
var ctx = context.Background()

func TestUnitEtaAndPodParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	etaParser := NewEtaParser(datetime.NewDatetime())

	infoAboutMovingResponse, err := requestMockUp.GetInfoAboutMovingRawResponse(ctx, "")
	assert.NoError(t, err)
	podParser := NewPodParser()
	pod := podParser.get(infoAboutMovingResponse)
	assert.Equal(t, "", pod)

	etaResponse, err := requestMockUp.GetEtaRawResponse(ctx, "")
	assert.NoError(t, err)
	event, err := etaParser.get(etaResponse, pod)
	assert.NoError(t, err)
	//2022-05-22 23:00
	expectedTime := time.Date(2022, 05, 22, 23, 00, 0, 0, time.UTC)
	assert.Equal(t, expectedTime, event.Time)
	assert.Equal(t, "", event.Vessel)
	assert.Equal(t, "ETA", event.OperationName)
	assert.Equal(t, pod, event.Location)
}

func TestContainerSizeParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	const expectedContainerSize = "40HQ"

	containerSizeParser := NewContainerSizeParser()

	infoAboutMovingResponse, err := requestMockUp.GetInfoAboutMovingRawResponse(ctx, "")
	assert.NoError(t, err)

	assert.Equal(t, expectedContainerSize, containerSizeParser.get(infoAboutMovingResponse))

}

func TestInfoAboutMovingParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	parser := NewInfoAboutMovingParser(datetime.NewDatetime())

	infoAboutMovingResponse, err := requestMockUp.GetInfoAboutMovingRawResponse(ctx, "")
	assert.NoError(t, err)

	infoAboutMoving := parser.get(infoAboutMovingResponse)

	const expectedInfoAboutMovingLen = 7
	assert.Equal(t, expectedInfoAboutMovingLen, len(infoAboutMoving))

}
