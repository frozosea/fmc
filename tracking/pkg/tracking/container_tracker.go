package tracking

import (
	"context"
)

type ContainerTracker struct {
	trackers map[string]IContainerTracker
}

func NewContainerTracker(trackers map[string]IContainerTracker) *ContainerTracker {
	return &ContainerTracker{trackers: trackers}
}

func (c *ContainerTracker) Track(ctx context.Context, scac, number string) (*ContainerTrackingResponse, error) {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	ch := make(chan *ContainerTrackingResponse)
	counter := make(chan int, len(c.trackers))
	if scac == "AUTO" {
		index := 0
		for _, tracker := range c.trackers {
			go func(c context.Context, tracker IContainerTracker) {
				defer func() {
					index++
					counter <- index
				}()
				response, err := tracker.Track(c, number)
				if err != nil {
					return
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
