package user

import (
	"context"
	"database/sql"
	pb "github.com/frozosea/fmc-pb/schedule-tracking"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user-api/internal/domain"
	"user-api/pkg/util"
)

type NoTaskError struct {
}

func (e *NoTaskError) Error() string {
	return "no task with this id!"
}

type IScheduleTrackingInfoRepository interface {
	GetInfo(ctx context.Context, number, token string) (*domain.ScheduleTrackingInfoObject, error)
}

type ScheduleTrackingInfoRepository struct {
	cli pb.ScheduleTrackingClient
}

func NewScheduleTrackingInfoRepository(conn *grpc.ClientConn) *ScheduleTrackingInfoRepository {
	return &ScheduleTrackingInfoRepository{cli: pb.NewScheduleTrackingClient(conn)}
}

func (r *ScheduleTrackingInfoRepository) GetInfo(ctx context.Context, number, token string) (*domain.ScheduleTrackingInfoObject, error) {
	response, err := r.cli.GetInfoAboutTrack(ctx, &pb.GetInfoAboutTrackRequest{
		Number: number,
	}, util.GenerateGRPCAuthHeader(token))
	if err != nil {
		statusCode := status.Convert(err).Code()
		switch statusCode {
		case codes.NotFound:
			return nil, &NoTaskError{}
		case codes.PermissionDenied:
			return nil, &NoTaskError{}
		default:
			return nil, err
		}
	}
	s := response.GetScheduleTrackingInfo()
	return &domain.ScheduleTrackingInfoObject{
		Emails:  s.GetEmails(),
		Subject: s.GetSubject(),
		Time:    s.GetTime(),
	}, nil
}

type IRepository interface {
	AddContainerToAccount(ctx context.Context, userId int, containers []string) error
	AddBillNumberToAccount(ctx context.Context, userId int, containers []string) error
	DeleteContainersFromAccount(ctx context.Context, userId int, numbers []string) error
	DeleteBillNumbersFromAccount(ctx context.Context, userId int, numbers []string) error
	GetAllContainersAndBillNumbers(ctx context.Context, userId int) (*domain.AllContainersAndBillNumbers, error)
}

type Repository struct {
	scheduleTrackingInfoRepository IScheduleTrackingInfoRepository
	db                             *sql.DB
}

func NewRepository(db *sql.DB, scheduleTrackingInfoRepository IScheduleTrackingInfoRepository) *Repository {
	return &Repository{db: db, scheduleTrackingInfoRepository: scheduleTrackingInfoRepository}
}
func (r *Repository) checkContainerOrBillExists(ctx context.Context, userId int, isContainer bool, number string) bool {
	if isContainer {
		row := r.db.QueryRowContext(ctx, `SELECT c.number FROM "containers" AS c WHERE c.user_id = $1 AND c.number = $2`, userId, number)
		var s sql.NullString
		if scanErr := row.Scan(&s); scanErr != nil {
			return false
		}
		if s.String != "" || s.Valid {
			return true
		}

	} else {
		row := r.db.QueryRowContext(ctx, `SELECT c.number FROM "bill_numbers" AS c WHERE c.user_id = $1 AND c.number = $2`, userId, number)
		var s sql.NullString
		if scanErr := row.Scan(&s); scanErr != nil {
			return false
		}
		if s.String != "" || s.Valid {
			return true
		}
	}
	return false
}
func (r *Repository) AddContainerToAccount(ctx context.Context, userId int, containers []string) error {
	for _, v := range containers {
		if !r.checkContainerOrBillExists(ctx, userId, true, v) {
			_, err := r.db.ExecContext(ctx, `INSERT INTO "containers" (number,user_id,is_on_track,is_arrived) VALUES($1,$2,false,false)`, v, userId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (r *Repository) AddBillNumberToAccount(ctx context.Context, userId int, containers []string) error {
	for _, v := range containers {
		if !r.checkContainerOrBillExists(ctx, userId, false, v) {
			_, err := r.db.ExecContext(ctx, `INSERT INTO "bill_numbers" (number,user_id,is_on_track,is_arrived) VALUES($1,$2,false,false)`, v, userId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Repository) DeleteContainersFromAccount(ctx context.Context, userId int, numbers []string) error {
	for _, v := range numbers {
		_, err := r.db.ExecContext(ctx, `DELETE FROM "containers" AS c WHERE c.user_id = $1 AND c.number = $2`, userId, v)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *Repository) DeleteBillNumbersFromAccount(ctx context.Context, userId int, numbers []string) error {
	for _, v := range numbers {
		_, err := r.db.ExecContext(ctx, `DELETE FROM "bill_numbers" AS c WHERE c.user_id = $1 AND c.number = $2`, userId, v)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *Repository) getAllContainers(ctx context.Context, userId int) ([]*domain.Container, error) {
	var containers []*domain.Container
	rows, err := r.db.QueryContext(ctx, `SELECT DISTINCT ON (c.number)  c.number,c.is_on_track FROM "containers" AS c WHERE c.user_id = $1 AND c.is_arrived = false`, userId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		var container domain.Container
		container.IsContainer = true
		if err := rows.Scan(&container.Number, &container.IsOnTrack); err != nil {
			return containers, err
		}
		if container.IsOnTrack {
			scheduleTrackingInfo, err := r.scheduleTrackingInfoRepository.GetInfo(ctx, container.Number, userId)
			if err != nil {
				switch err.(type) {
				case *NoTaskError:
					container.ScheduleTrackingInfo = nil
					container.IsOnTrack = false
				default:
					return containers, err

				}
			} else {
				container.ScheduleTrackingInfo = scheduleTrackingInfo
				container.IsOnTrack = true
			}
		}
		containers = append(containers, &container)
	}
	return containers, nil
}
func (r *Repository) getAllBillNumbers(ctx context.Context, userId int) ([]*domain.Container, error) {
	var bills []*domain.Container
	rows, err := r.db.QueryContext(ctx, `SELECT DISTINCT ON (c.number) c.number,c.is_on_track FROM "bill_numbers" AS c WHERE c.user_id = $1 AND c.is_arrived = false`, userId)
	if err != nil {
		return bills, err
	}
	defer rows.Close()
	for rows.Next() {
		var bill domain.Container
		bill.IsContainer = false
		if err := rows.Scan(&bill.Number, &bill.IsOnTrack); err != nil {
			return bills, err
		}
		if bill.IsOnTrack {
			scheduleTrackingInfo, err := r.scheduleTrackingInfoRepository.GetInfo(ctx, bill.Number, userId)
			if err != nil {
				switch err.(type) {
				case *NoTaskError:
					bill.ScheduleTrackingInfo = nil
					bill.IsOnTrack = false
				default:
					return bills, err

				}
			} else {
				bill.ScheduleTrackingInfo = scheduleTrackingInfo
				bill.IsOnTrack = true
			}
		}
		bills = append(bills, &bill)
	}
	return bills, nil
}
func (r *Repository) GetAllContainersAndBillNumbers(ctx context.Context, userId int) (*domain.AllContainersAndBillNumbers, error) {
	var allBillNumbersAndContainers domain.AllContainersAndBillNumbers
	containers, err := r.getAllContainers(ctx, userId)
	if err != nil {
		return &allBillNumbersAndContainers, err
	}
	billNumbers, billErr := r.getAllBillNumbers(ctx, userId)
	if billErr != nil {
		return &allBillNumbersAndContainers, billErr
	}
	allBillNumbersAndContainers.Containers = containers
	allBillNumbersAndContainers.BillNumbers = billNumbers
	return &allBillNumbersAndContainers, nil
}
