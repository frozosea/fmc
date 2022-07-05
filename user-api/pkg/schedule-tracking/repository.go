package schedule_tracking

import (
	"context"
	"database/sql"
)

type IRepository interface {
	AddMarkContainerOnTrack(ctx context.Context, number string, userId int) error
	AddMarkContainerWasArrived(ctx context.Context, number string, userId int) error
	AddMarkContainerWasRemovedFromTrack(ctx context.Context, number string, userId int) error
	AddMarkBillNoOnTrack(ctx context.Context, number string, userId int) error
	AddMarkBillNoWasArrived(ctx context.Context, number string, userId int) error
	AddMarkBillNoWasRemovedFromTrack(ctx context.Context, number string, userId int) error
}
type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) wrapperExec(ctx context.Context, sql string, args ...interface{}) error {
	_, err := r.db.ExecContext(ctx, sql, args...)
	return err
}
func (r *Repository) AddMarkContainerOnTrack(ctx context.Context, number string, userId int) error {
	return r.wrapperExec(ctx, `UPDATE "containers" AS c SET is_on_track = true WHERE c.number = $1 AND c.user_id = $2`, number, userId)
}
func (r *Repository) AddMarkContainerWasArrived(ctx context.Context, number string, userId int) error {
	return r.wrapperExec(ctx, `UPDATE "containers" AS c SET is_arrived = true WHERE c.number = $1 AND c.user_id = $2`, number, userId)
}
func (r *Repository) AddMarkContainerWasRemovedFromTrack(ctx context.Context, number string, userId int) error {
	return r.wrapperExec(ctx, `UPDATE "containers" AS c SET is_on_track = false WHERE c.number = $1 AND c.user_od = $2`, number, userId)
}
func (r *Repository) AddMarkBillNoOnTrack(ctx context.Context, number string, userId int) error {
	return r.wrapperExec(ctx, `UPDATE "bill_numbers" AS b SET is_on_track = true WHERE b.number = $1 AND b.user_id = $2`, number, userId)
}
func (r *Repository) AddMarkBillNoWasArrived(ctx context.Context, number string, userId int) error {
	return r.wrapperExec(ctx, `UPDATE "bill_numbers" AS b SET is_arrived = true WHERE b.number = $1 AND b.user_id = $2`, number, userId)
}
func (r *Repository) AddMarkBillNoWasRemovedFromTrack(ctx context.Context, number string, userId int) error {
	return r.wrapperExec(ctx, `UPDATE "bill_numbers" AS b SET is_on_track = false WHERE b.number = $1 AND b.user_od = $2`, number, userId)
}
