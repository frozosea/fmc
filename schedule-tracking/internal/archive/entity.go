package archive

import "schedule-tracking/pkg/tracking"

type AllArchive struct {
	containers []*tracking.ContainerNumberResponse
	bills      []*tracking.BillNumberResponse
}
