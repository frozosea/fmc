package tracking

import (
	"context"
	"golang_tracking/pkg/tracking/util/time_inspector"
)

type ContainerTracker struct {
	trackers      map[string]IContainerTracker
	timeInspector time_inspector.ITimeInspector
}

func NewContainerTracker(trackers map[string]IContainerTracker, inspector time_inspector.ITimeInspector) *ContainerTracker {
	return &ContainerTracker{trackers: trackers, timeInspector: inspector}
}

func (c *ContainerTracker) Track(ctx context.Context, scac, number string) (*ContainerTrackingResponse, error) {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	ch := make(chan *ContainerTrackingResponse)
	counter := make(chan int, len(c.trackers))
	if scac == "AUTO" {
		index := 0
		for _, tracker := range c.trackers {
			go func(innerCtx context.Context, tracker IContainerTracker) {
				defer func() {
					index++
					counter <- index
				}()
				response, err := tracker.Track(innerCtx, number)
				if err != nil {
					return
				}
				if len(response.InfoAboutMoving) > 0 {
					if valid, err := c.timeInspector.CheckInfoAboutMovingExpires(response.InfoAboutMoving[len(response.InfoAboutMoving)-1].Time); valid && err == nil {
						ch <- response
						cancel()
						return
					} else {
						return
					}
				}
				ch <- response
				cancel()
			}(ctxWithCancel, tracker)
		}
	} else {
		tracker := c.trackers[scac]
		if tracker != nil {
			return tracker.Track(ctx, number)
		}
		return nil, NewNoScacException()
	}
	for {
		select {
		case <-ctx.Done():
			cancel()
			return nil, nil
		case <-ctxWithCancel.Done():
			return nil, NewNumberNotFoundException(number)
		case i := <-counter:
			if i == (len(c.trackers)) {
				cancel()
				return nil, NewNumberNotFoundException(number)
			}
		case data := <-ch:
			cancel()
			return data, nil
		}
	}
}
