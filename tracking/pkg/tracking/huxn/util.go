package huxn

func ReverseContainerTrackingEvents(events []*ContainerTrackingOneEvent) []*ContainerTrackingOneEvent {
	for i, j := 0, len(events)-1; i < j; i, j = i+1, j-1 {
		events[i], events[j] = events[j], events[i]
	}
	return events
}
