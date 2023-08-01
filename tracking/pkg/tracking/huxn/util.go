package huxn

import (
	"golang_tracking/pkg/tracking"
	"strings"
)

func ReverseTrackingEvents(events []*OneEvent) []*OneEvent {
	for i, j := 0, len(events)-1; i < j; i, j = i+1, j-1 {
		events[i], events[j] = events[j], events[i]
	}
	return events
}

func checkNumberArrived(infoAboutMoving []*tracking.Event) bool {
	for _, item := range infoAboutMoving {
		if strings.EqualFold(item.OperationName, "Discharge from vessel full") ||
			strings.EqualFold(item.OperationName, "Discharge from vessel empty") {
			return true
		}
	}
	return false
}
