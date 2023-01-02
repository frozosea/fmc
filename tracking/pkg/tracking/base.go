package tracking

import (
	"context"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
)

type IContainerTracker interface {
	Track(ctx context.Context, number string) (*ContainerTrackingResponse, error)
}

type IBillTracker interface {
	Track(ctx context.Context, number string) (*BillNumberTrackingResponse, error)
}

type BaseConstructorArgumentsForTracker struct {
	Request            requests.IHttp
	UserAgentGenerator requests.IUserAgentGenerator
	Datetime           datetime.IDatetime
}
