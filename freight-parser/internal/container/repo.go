package container

import (
	"context"
	"database/sql"
)

type IRepository interface {
	GetAll(ctx context.Context) ([]*Container, error)
	Add(ctx context.Context, containerType string) error
	Update(ctx context.Context, containerId int, newContainerType string) error
	Delete(ctx context.Context, containerId int) error
}
type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context) ([]*Container, error) {
	var ar []*Container
	rows, err := r.db.QueryContext(ctx, `SELECT * FROM "containers"`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		container := new(Container)
		if err := rows.Scan(&container.Id, container.Type); err != nil {
			return nil, err
		}
		ar = append(ar, container)
	}
	return ar, nil
}
func (r *Repository) wrapExecContext(ctx context.Context, sqlString string, args ...interface{}) error {
	_, err := r.db.ExecContext(ctx, sqlString, args...)
	return err
}
func (r *Repository) Add(ctx context.Context, containerType string) error {
	return r.wrapExecContext(ctx, `INSERT INTO "containers" (type) VALUES ($1)`, containerType)
}
func (r *Repository) Update(ctx context.Context, containerId int, newContainerType string) error {
	return r.wrapExecContext(ctx, `UPDATE "containers" AS c type = $1 WHERE c.id = $2`, newContainerType, containerId)

}
func (r *Repository) Delete(ctx context.Context, containerId int) error {
	return r.wrapExecContext(ctx, `DELETE "containers" AS c WHERE c.id = $2`, containerId)

}
