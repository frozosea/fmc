package user

import (
	"context"
	"database/sql"
	"user-api/pkg/domain"
)

type IRepository interface {
	AddContainerToAccount(ctx context.Context, userId int, containers []string) error
	AddBillNumberToAccount(ctx context.Context, userId int, containers []string) error
	DeleteContainersFromAccount(ctx context.Context, userId int, numberIds []int64) error
	DeleteBillNumbersFromAccount(ctx context.Context, userId int, numberIds []int64) error
	GetAllContainersAndBillNumbers(ctx context.Context, userId int) (*domain.AllContainersAndBillNumbers, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddContainerToAccount(ctx context.Context, userId int, containers []string) error {
	for _, v := range containers {
		_, err := r.db.ExecContext(ctx, `INSERT INTO "containers" (number,user_id,is_on_track,is_arrived) VALUES($1,$2,false,false)`, v, userId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *Repository) AddBillNumberToAccount(ctx context.Context, userId int, containers []string) error {
	for _, v := range containers {
		_, err := r.db.ExecContext(ctx, `INSERT INTO "bill_numbers" (number,user_id,is_on_track,is_arrived) VALUES($1,$2,false,false)`, v, userId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) DeleteContainersFromAccount(ctx context.Context, userId int, numberIds []int64) error {
	for _, v := range numberIds {
		_, err := r.db.ExecContext(ctx, `DELETE FROM "containers" AS c WHERE c.user_id = $1 AND c.id = $2`, userId, v)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *Repository) DeleteBillNumbersFromAccount(ctx context.Context, userId int, numberIds []int64) error {
	for _, v := range numberIds {
		_, err := r.db.ExecContext(ctx, `DELETE FROM "bill_numbers" AS c WHERE c.user_id = $1 AND c.id = $2`, userId, v)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *Repository) getAllContainers(ctx context.Context, userId int) ([]*domain.Container, error) {
	var containers []*domain.Container
	containerRows, err := r.db.QueryContext(ctx, `SELECT DISTINCT ON (c.number)  c.id,c.number,c.is_on_track FROM "containers" AS c WHERE c.user_id = $1 AND c.is_arrived = false`, userId)
	if err != nil {
		return containers, nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(containerRows)
	for containerRows.Next() {
		var container domain.Container
		if scanErr := containerRows.Scan(&container.Id, &container.Number, &container.IsOnTrack); scanErr != nil {
			return containers, scanErr
		}
		containers = append(containers, &container)
	}
	return containers, nil
}
func (r *Repository) getAllBillNumbers(ctx context.Context, userId int) ([]*domain.Container, error) {
	var containers []*domain.Container
	containerRows, err := r.db.QueryContext(ctx, `SELECT DISTINCT ON (c.number) c.id,c.number,c.is_on_track FROM "bill_numbers" AS c WHERE c.user_id = $1 AND c.is_arrived = false`, userId)
	if err != nil {
		return containers, err
	}
	defer containerRows.Close()
	for containerRows.Next() {
		var container domain.Container
		if scanErr := containerRows.Scan(&container.Id, &container.Number, &container.IsOnTrack); scanErr != nil {
			return containers, scanErr
		}
		containers = append(containers, &container)
	}
	return containers, nil
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
