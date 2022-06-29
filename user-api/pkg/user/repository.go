package user

import (
	"context"
	"database/sql"
	"user-api/pkg/domain"
)

type IRepository interface {
	AddContainerToAccount(ctx context.Context, userId int, containers []domain.Container) error
	DeleteContainersFromAccount(ctx context.Context, userId int, containers []domain.Container) error
	GetAllContainers(ctx context.Context, userId int) ([]*domain.Container, error)
}

type Repository struct {
	db *sql.DB
}

func (r *Repository) AddContainerToAccount(ctx context.Context, userId int, containers []domain.Container) error {
	for _, v := range containers {
		_, err := r.db.ExecContext(ctx, `INSERT INTO "containers" (number,user_id) VALUES($1,$2)`, userId, v.Number)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) DeleteContainersFromAccount(ctx context.Context, userId int, containers []domain.Container) error {
	for _, v := range containers {
		_, err := r.db.ExecContext(ctx, `DELETE FROM "containers" AS c WHERE c.user_id = $1 AND c.number = $2`, userId, v.Number)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) GetAllContainers(ctx context.Context, userId int) ([]*domain.Container, error) {
	var containers []*domain.Container
	rows, err := r.db.QueryContext(ctx, `SELECT * FROM "containers" AS c WHERE c.user_id = $1`, userId)
	if err != nil {
		return containers, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		container := new(domain.Container)
		if scanErr := rows.Scan(&container.Number); scanErr != nil {
			return containers, scanErr
		}
		containers = append(containers, container)
	}
	return containers, nil
}
