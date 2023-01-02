package tracking

import (
	"context"
)

type BillTracker struct {
	trackers map[string]IBillTracker
}

func NewBillNumberTracker(trackers map[string]IBillTracker) *BillTracker {
	return &BillTracker{trackers: trackers}
}

func (c *BillTracker) Track(ctx context.Context, scac, number string) (*BillNumberTrackingResponse, error) {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	ch := make(chan *BillNumberTrackingResponse)
	counter := make(chan int, len(c.trackers))
	if scac == "AUTO" {
		index := 0
		for _, tracker := range c.trackers {
			go func(c context.Context, tracker IBillTracker) {
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
