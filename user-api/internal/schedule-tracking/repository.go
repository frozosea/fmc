package schedule_tracking

import (
	"context"
	"database/sql"
	"errors"
)

type IRepository interface {
	AddMarkContainerOnTrack(ctx context.Context, number string, userId int64) error
	AddMarkContainerWasArrived(ctx context.Context, number string, userId int64) error
	AddMarkContainerWasRemovedFromTrack(ctx context.Context, number string, userId int64) error
	AddMarkBillNoOnTrack(ctx context.Context, number string, userId int64) error
	AddMarkBillNoWasArrived(ctx context.Context, number string, userId int64) error
	AddMarkBillNoWasRemovedFromTrack(ctx context.Context, number string, userId int64) error
	CheckNumberExists(ctx context.Context, number string, userId int64, isContainer bool) (bool, error)
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
func (r *Repository) AddMarkContainerOnTrack(ctx context.Context, number string, userId int64) error {
	return r.wrapperExec(ctx, `UPDATE "containers" AS c SET is_on_track = true WHERE c.number = $1 AND c.user_id = $2`, number, userId)
}
func (r *Repository) AddMarkContainerWasArrived(ctx context.Context, number string, userId int64) error {
	return r.wrapperExec(ctx, `UPDATE "containers" AS c SET is_arrived = true WHERE c.number = $1 AND c.user_id = $2`, number, userId)
}
func (r *Repository) AddMarkContainerWasRemovedFromTrack(ctx context.Context, number string, userId int64) error {
	return r.wrapperExec(ctx, `UPDATE "containers" AS c SET is_on_track = false WHERE c.number = $1 AND c.user_id = $2`, number, userId)
}
func (r *Repository) AddMarkBillNoOnTrack(ctx context.Context, number string, userId int64) error {
	return r.wrapperExec(ctx, `UPDATE "bill_numbers" AS b SET is_on_track = true WHERE b.number = $1 AND b.user_id = $2`, number, userId)
}
func (r *Repository) AddMarkBillNoWasArrived(ctx context.Context, number string, userId int64) error {
	return r.wrapperExec(ctx, `UPDATE "bill_numbers" AS b SET is_arrived = true WHERE b.number = $1 AND b.user_id = $2`, number, userId)
}
func (r *Repository) AddMarkBillNoWasRemovedFromTrack(ctx context.Context, number string, userId int64) error {
	return r.wrapperExec(ctx, `UPDATE "bill_numbers" AS b SET is_on_track = false WHERE b.number = $1 AND b.user_id = $2`, number, userId)
}
func (r *Repository) CheckNumberExists(ctx context.Context, number string, userId int64, isContainer bool) (bool, error) {
	var returnedNumber sql.NullString
	if isContainer {
		if err := r.db.QueryRowContext(ctx, `SELECT c.number FROM "containers" AS c WHERE c.number = $1 AND c.user_id = $2`, number, userId).Scan(&returnedNumber); err != nil {
			return false, err
		}
		if returnedNumber.Valid {
			return true, nil
		}
	} else {
		if err := r.db.QueryRowContext(ctx, `SELECT c.number FROM "bill_numbers" AS c WHERE c.number = $1 AND c.user_id = $2`, number, userId).Scan(&returnedNumber); err != nil {
			return false, err
		}
		if returnedNumber.Valid {
			return true, nil
		}
	}
	return false, errors.New("no number exists")
}
