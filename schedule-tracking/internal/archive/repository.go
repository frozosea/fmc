package archive

import (
	"context"
	"database/sql"
	"encoding/json"
	"schedule-tracking/pkg/tracking"
)

type IRepository interface {
	GetContainers(ctx context.Context, userId int) ([]*tracking.ContainerNumberResponse, error)
	GetBills(ctx context.Context, userId int) ([]*tracking.BillNumberResponse, error)
	GetAll(ctx context.Context, userId int) (*AllArchive, error)
	AddByContainer(ctx context.Context, userId int, info *tracking.ContainerNumberResponse) error
	AddByBill(ctx context.Context, userId int, info *tracking.BillNumberResponse) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetContainers(ctx context.Context, userId int) ([]*tracking.ContainerNumberResponse, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT co.response FROM "container_archive" AS co WHERE co.id = $1`, userId)
	if err != nil {
		return nil, err
	}
	var allContainers []*tracking.ContainerNumberResponse
	for rows.Next() {
		var rawJson string
		if err := rows.Scan(&rawJson); err != nil {
			return nil, err
		}
		var response *tracking.ContainerNumberResponse
		if err := json.Unmarshal([]byte(rawJson), &response); err != nil {
			return nil, err
		}
		allContainers = append(allContainers, response)
	}
	return allContainers, nil
}

func (r *Repository) GetBills(ctx context.Context, userId int) ([]*tracking.BillNumberResponse, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT co.response FROM "bill_archive" AS co WHERE co.id = $1`, userId)
	if err != nil {
		return nil, err
	}
	var allBills []*tracking.BillNumberResponse
	for rows.Next() {
		var rawJson string
		if err := rows.Scan(&rawJson); err != nil {
			return nil, err
		}
		var response *tracking.BillNumberResponse
		if err := json.Unmarshal([]byte(rawJson), &response); err != nil {
			return nil, err
		}
		allBills = append(allBills, response)
	}
	return allBills, nil
}

func (r *Repository) GetAll(ctx context.Context, userId int) (*AllArchive, error) {
	bills, err := r.GetBills(ctx, userId)
	if err != nil {
		return nil, err
	}
	containers, err := r.GetContainers(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &AllArchive{
		containers: containers,
		bills:      bills,
	}, nil
}

func (r *Repository) AddByContainer(ctx context.Context, userId int, info *tracking.ContainerNumberResponse) error {
	j, err := json.Marshal(info)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, `INSERT INTO "container_archive" (user_id,response) VALUES ($1,$2)`, userId, string(j))
	return err
}

func (r *Repository) AddByBill(ctx context.Context, userId int, info *tracking.BillNumberResponse) error {
	j, err := json.Marshal(info)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, `INSERT INTO "bill_archive" (user_id,response) VALUES ($1,$2)`, userId, string(j))
	return err
}
