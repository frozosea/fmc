package history

import (
	"context"
	"fmc-gateway/internal/tracking"
	pb "fmc-gateway/pkg/schedule-tracking-pb"
	"google.golang.org/grpc"
)

type IHistoryTasksClient interface {
	Get(ctx context.Context, userId int) (*TasksArchive, error)
}
type ScheduleTrackingTasksClient struct {
	cli       pb.ArchiveClient
	converter *scheduleTrackingTasksConverter
}

func NewScheduleTrackingTasksClient(conn *grpc.ClientConn) *ScheduleTrackingTasksClient {
	return &ScheduleTrackingTasksClient{cli: pb.NewArchiveClient(conn), converter: newScheduleTrackingTasksConverter()}
}

func (c *ScheduleTrackingTasksClient) Get(ctx context.Context, userId int) (*TasksArchive, error) {
	result, err := c.cli.GetArchive(ctx, c.converter.GetArchiveRequestConvert(userId))
	if err != nil {
		return nil, err
	}
	return c.converter.GetArchiveConvert(result), nil
}

type scheduleTrackingTasksConverter struct {
}

func newScheduleTrackingTasksConverter() *scheduleTrackingTasksConverter {
	return &scheduleTrackingTasksConverter{}
}

func (c *scheduleTrackingTasksConverter) infoAboutMovingConvert(r []*pb.InfoAboutMoving) []tracking.BaseInfoAboutMoving {
	var ar []tracking.BaseInfoAboutMoving
	for _, v := range r {
		ar = append(ar, tracking.BaseInfoAboutMoving{
			Time:          v.GetTime(),
			OperationName: v.GetOperationName(),
			Location:      v.GetLocation(),
			Vessel:        v.GetVessel(),
		})
	}
	return ar
}
func (c *scheduleTrackingTasksConverter) GetArchiveConvert(r *pb.GetAllBillsContainerResponse) *TasksArchive {
	var cntrs []*tracking.ContainerNumberResponse
	for _, cntr := range r.GetContainers() {
		cntrs = append(cntrs, &tracking.ContainerNumberResponse{
			Container:       cntr.GetContainer(),
			ContainerSize:   cntr.GetContainerSize(),
			Scac:            cntr.GetScac(),
			InfoAboutMoving: c.infoAboutMovingConvert(cntr.GetInfoAboutMoving()),
		})
	}
	var bills []*tracking.BillNumberResponse
	for _, bill := range r.GetBills() {
		bills = append(bills, &tracking.BillNumberResponse{
			BillNo:           bill.GetBillNo(),
			Scac:             bill.GetScac(),
			InfoAboutMoving:  c.infoAboutMovingConvert(bill.GetInfoAboutMoving()),
			EtaFinalDelivery: bill.GetEtaFinalDelivery(),
		})
	}
	return &TasksArchive{
		Containers: cntrs,
		Bills:      bills,
	}
}
func (c *scheduleTrackingTasksConverter) GetArchiveRequestConvert(userId int) *pb.GetArchiveRequest {
	return &pb.GetArchiveRequest{UserId: int64(userId)}
}
