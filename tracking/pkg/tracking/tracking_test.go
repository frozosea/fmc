package tracking

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type containerTrackingMockUp struct {
	scac              string
	shouldReturnError bool
}

func newContainerTrackingMockUp(scac string, shouldReturnError bool) *containerTrackingMockUp {
	return &containerTrackingMockUp{scac: scac, shouldReturnError: shouldReturnError}
}

func (c *containerTrackingMockUp) Track(_ context.Context, number string) (*ContainerTrackingResponse, error) {
	if c.shouldReturnError {
		return nil, NewNumberNotFoundException(number)
	}
	return &ContainerTrackingResponse{
		Number:          number,
		Size:            "size",
		Scac:            c.scac,
		InfoAboutMoving: nil,
	}, nil
}

type TimeInspectorMockup struct {
}

func NewTimeInspectorMockup() *TimeInspectorMockup {
	return &TimeInspectorMockup{}
}

func (t *TimeInspectorMockup) CheckInfoAboutMovingExpires(date time.Time) (bool, error) {
	return true, nil
}

var inspector = NewTimeInspectorMockup()

func TestContainerTracker(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	testSuitTable := []struct {
		scac     string
		trackers map[string]IContainerTracker
	}{
		{
			scac: "COSU",
			trackers: map[string]IContainerTracker{
				"COSU": newContainerTrackingMockUp("COSU", false),
				"FESO": newContainerTrackingMockUp("FESO", true),
				"HALU": newContainerTrackingMockUp("HALU", true),
				"MAEU": newContainerTrackingMockUp("MEAU", true),
				"MSCU": newContainerTrackingMockUp("MSCU", true),
				"ONEY": newContainerTrackingMockUp("ONEY", true),
				"SITC": newContainerTrackingMockUp("SITC", true),
				"SKLU": newContainerTrackingMockUp("SKLU", true),
			},
		},
		{
			scac: "FESO",
			trackers: map[string]IContainerTracker{
				"COSU": newContainerTrackingMockUp("COSU", true),
				"FESO": newContainerTrackingMockUp("FESO", false),
				"HALU": newContainerTrackingMockUp("HALU", true),
				"MAEU": newContainerTrackingMockUp("MEAU", true),
				"MSCU": newContainerTrackingMockUp("MSCU", true),
				"ONEY": newContainerTrackingMockUp("ONEY", true),
				"SITC": newContainerTrackingMockUp("SITC", true),
				"SKLU": newContainerTrackingMockUp("SKLU", true),
			},
		},
		{
			scac: "HALU",
			trackers: map[string]IContainerTracker{
				"COSU": newContainerTrackingMockUp("COSU", true),
				"FESO": newContainerTrackingMockUp("FESO", true),
				"HALU": newContainerTrackingMockUp("HALU", false),
				"MAEU": newContainerTrackingMockUp("MAEU", true),
				"MSCU": newContainerTrackingMockUp("MSCU", true),
				"ONEY": newContainerTrackingMockUp("ONEY", true),
				"SITC": newContainerTrackingMockUp("SITC", true),
				"SKLU": newContainerTrackingMockUp("SKLU", true),
			},
		},
		{
			scac: "MAEU",
			trackers: map[string]IContainerTracker{
				"COSU": newContainerTrackingMockUp("COSU", true),
				"FESO": newContainerTrackingMockUp("FESO", true),
				"HALU": newContainerTrackingMockUp("HALU", true),
				"MAEU": newContainerTrackingMockUp("MAEU", false),
				"MSCU": newContainerTrackingMockUp("MSCU", true),
				"ONEY": newContainerTrackingMockUp("ONEY", true),
				"SITC": newContainerTrackingMockUp("SITC", true),
				"SKLU": newContainerTrackingMockUp("SKLU", true),
			},
		},
		{
			scac: "MSCU",
			trackers: map[string]IContainerTracker{
				"COSU": newContainerTrackingMockUp("COSU", true),
				"FESO": newContainerTrackingMockUp("FESO", true),
				"HALU": newContainerTrackingMockUp("HALU", true),
				"MAEU": newContainerTrackingMockUp("MAEU", true),
				"MSCU": newContainerTrackingMockUp("MSCU", false),
				"ONEY": newContainerTrackingMockUp("ONEY", true),
				"SITC": newContainerTrackingMockUp("SITC", true),
				"SKLU": newContainerTrackingMockUp("SKLU", true),
			},
		},
		{
			scac: "ONEY",
			trackers: map[string]IContainerTracker{
				"COSU": newContainerTrackingMockUp("COSU", true),
				"FESO": newContainerTrackingMockUp("FESO", true),
				"HALU": newContainerTrackingMockUp("HALU", true),
				"MAEU": newContainerTrackingMockUp("MAEU", true),
				"MSCU": newContainerTrackingMockUp("MSCU", true),
				"ONEY": newContainerTrackingMockUp("ONEY", false),
				"SITC": newContainerTrackingMockUp("SITC", true),
				"SKLU": newContainerTrackingMockUp("SKLU", true),
			},
		},
		{
			scac: "SITC",
			trackers: map[string]IContainerTracker{
				"COSU": newContainerTrackingMockUp("COSU", true),
				"FESO": newContainerTrackingMockUp("FESO", true),
				"HALU": newContainerTrackingMockUp("HALU", true),
				"MAEU": newContainerTrackingMockUp("MAEU", true),
				"MSCU": newContainerTrackingMockUp("MSCU", true),
				"ONEY": newContainerTrackingMockUp("ONEY", true),
				"SITC": newContainerTrackingMockUp("SITC", false),
				"SKLU": newContainerTrackingMockUp("SKLU", true),
			},
		},
		{
			scac: "SKLU",
			trackers: map[string]IContainerTracker{
				"COSU": newContainerTrackingMockUp("COSU", true),
				"FESO": newContainerTrackingMockUp("FESO", true),
				"HALU": newContainerTrackingMockUp("HALU", true),
				"MAEU": newContainerTrackingMockUp("MAEU", true),
				"MSCU": newContainerTrackingMockUp("MSCU", true),
				"ONEY": newContainerTrackingMockUp("ONEY", true),
				"SITC": newContainerTrackingMockUp("SITC", true),
				"SKLU": newContainerTrackingMockUp("SKLU", false),
			},
		},
	}

	ctx := context.Background()

	for _, suite := range testSuitTable {
		containerTracker := NewContainerTracker(suite.trackers, inspector)
		response, err := containerTracker.Track(ctx, "AUTO", "number")
		assert.NoError(t, err)
		assert.Equal(t, response.Scac, suite.scac)
	}
	for _, suite := range testSuitTable {
		containerTracker := NewContainerTracker(suite.trackers, inspector)
		response, err := containerTracker.Track(ctx, suite.scac, "number")
		assert.NoError(t, err)
		assert.Equal(t, response.Scac, suite.scac)
	}
}

type billNumberMockUp struct {
	scac              string
	shouldReturnError bool
}

func newBillNumberMockUp(scac string, shouldReturnError bool) *billNumberMockUp {
	return &billNumberMockUp{scac: scac, shouldReturnError: shouldReturnError}
}

func (b *billNumberMockUp) Track(_ context.Context, number string) (*BillNumberTrackingResponse, error) {
	if b.shouldReturnError {
		return nil, NewNumberNotFoundException(number)
	}
	return &BillNumberTrackingResponse{
		Number:          number,
		Scac:            b.scac,
		InfoAboutMoving: nil,
	}, nil
}

func TestBillTracker(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	testSuitTable := []struct {
		scac     string
		trackers map[string]IBillTracker
	}{
		{
			scac: "COSU",
			trackers: map[string]IBillTracker{
				"COSU": newBillNumberMockUp("COSU", false),
				"FESO": newBillNumberMockUp("FESO", true),
				"HALU": newBillNumberMockUp("HALU", true),
				"MAEU": newBillNumberMockUp("MEAU", true),
				"MSCU": newBillNumberMockUp("MSCU", true),
				"ONEY": newBillNumberMockUp("ONEY", true),
				"SITC": newBillNumberMockUp("SITC", true),
				"SKLU": newBillNumberMockUp("SKLU", true),
			},
		},
		{
			scac: "FESO",
			trackers: map[string]IBillTracker{
				"COSU": newBillNumberMockUp("COSU", true),
				"FESO": newBillNumberMockUp("FESO", false),
				"HALU": newBillNumberMockUp("HALU", true),
				"MAEU": newBillNumberMockUp("MEAU", true),
				"MSCU": newBillNumberMockUp("MSCU", true),
				"ONEY": newBillNumberMockUp("ONEY", true),
				"SITC": newBillNumberMockUp("SITC", true),
				"SKLU": newBillNumberMockUp("SKLU", true),
			},
		},
		{
			scac: "HALU",
			trackers: map[string]IBillTracker{
				"COSU": newBillNumberMockUp("COSU", true),
				"FESO": newBillNumberMockUp("FESO", true),
				"HALU": newBillNumberMockUp("HALU", false),
				"MAEU": newBillNumberMockUp("MAEU", true),
				"MSCU": newBillNumberMockUp("MSCU", true),
				"ONEY": newBillNumberMockUp("ONEY", true),
				"SITC": newBillNumberMockUp("SITC", true),
				"SKLU": newBillNumberMockUp("SKLU", true),
			},
		},
		{
			scac: "MAEU",
			trackers: map[string]IBillTracker{
				"COSU": newBillNumberMockUp("COSU", true),
				"FESO": newBillNumberMockUp("FESO", true),
				"HALU": newBillNumberMockUp("HALU", true),
				"MAEU": newBillNumberMockUp("MAEU", false),
				"MSCU": newBillNumberMockUp("MSCU", true),
				"ONEY": newBillNumberMockUp("ONEY", true),
				"SITC": newBillNumberMockUp("SITC", true),
				"SKLU": newBillNumberMockUp("SKLU", true),
			},
		},
		{
			scac: "MSCU",
			trackers: map[string]IBillTracker{
				"COSU": newBillNumberMockUp("COSU", true),
				"FESO": newBillNumberMockUp("FESO", true),
				"HALU": newBillNumberMockUp("HALU", true),
				"MAEU": newBillNumberMockUp("MAEU", true),
				"MSCU": newBillNumberMockUp("MSCU", false),
				"ONEY": newBillNumberMockUp("ONEY", true),
				"SITC": newBillNumberMockUp("SITC", true),
				"SKLU": newBillNumberMockUp("SKLU", true),
			},
		},
		{
			scac: "ONEY",
			trackers: map[string]IBillTracker{
				"COSU": newBillNumberMockUp("COSU", true),
				"FESO": newBillNumberMockUp("FESO", true),
				"HALU": newBillNumberMockUp("HALU", true),
				"MAEU": newBillNumberMockUp("MAEU", true),
				"MSCU": newBillNumberMockUp("MSCU", true),
				"ONEY": newBillNumberMockUp("ONEY", false),
				"SITC": newBillNumberMockUp("SITC", true),
				"SKLU": newBillNumberMockUp("SKLU", true),
			},
		},
		{
			scac: "SITC",
			trackers: map[string]IBillTracker{
				"COSU": newBillNumberMockUp("COSU", true),
				"FESO": newBillNumberMockUp("FESO", true),
				"HALU": newBillNumberMockUp("HALU", true),
				"MAEU": newBillNumberMockUp("MAEU", true),
				"MSCU": newBillNumberMockUp("MSCU", true),
				"ONEY": newBillNumberMockUp("ONEY", true),
				"SITC": newBillNumberMockUp("SITC", false),
				"SKLU": newBillNumberMockUp("SKLU", true),
			},
		},
		{
			scac: "SKLU",
			trackers: map[string]IBillTracker{
				"COSU": newBillNumberMockUp("COSU", true),
				"FESO": newBillNumberMockUp("FESO", true),
				"HALU": newBillNumberMockUp("HALU", true),
				"MAEU": newBillNumberMockUp("MAEU", true),
				"MSCU": newBillNumberMockUp("MSCU", true),
				"ONEY": newBillNumberMockUp("ONEY", true),
				"SITC": newBillNumberMockUp("SITC", true),
				"SKLU": newBillNumberMockUp("SKLU", false),
			},
		},
	}

	ctx := context.Background()

	for _, suite := range testSuitTable {
		containerTracker := NewBillNumberTracker(suite.trackers, inspector)
		response, err := containerTracker.Track(ctx, "AUTO", "number")
		assert.NoError(t, err)
		assert.Equal(t, response.Scac, suite.scac)
	}
	for _, suite := range testSuitTable {
		containerTracker := NewBillNumberTracker(suite.trackers, inspector)
		response, err := containerTracker.Track(ctx, suite.scac, "number")
		assert.NoError(t, err)
		assert.Equal(t, response.Scac, suite.scac)
	}
}
