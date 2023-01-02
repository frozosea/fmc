package tracking

func ReverseSliceWithEvents(events []*Event) []*Event {
	for i, j := 0, len(events)-1; i < j; i, j = i+1, j-1 {
		events[i], events[j] = events[j], events[i]
	}
	return events
}

func StringSliceWithStep(array []string, startIndex, lastIndex, step int) []string {
	var slice []string
	for i := startIndex; i <= lastIndex; i += step {
		slice = append(slice, array[i])
	}
	return slice
}
