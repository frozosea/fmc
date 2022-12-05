package domain

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
)

type NoTasksError struct{}

func (n *NoTasksError) Error() string {
	return "no tasks in table"
}

type IRepository interface {
	Add(ctx context.Context, req *BaseTrackReq, isContainer bool) error
	GetAll(ctx context.Context) ([]*TrackingTask, error)
	Update(ctx context.Context, req *BaseTrackReq, isContainer bool) error
	GetByNumber(ctx context.Context, number string) (*TrackingTask, error)
	Delete(ctx context.Context, numbers []string) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
func (r *Repository) checkNumber(ctx context.Context, number string) bool {
	row := r.db.QueryRowContext(ctx, `SELECT t.number FROM "tasks" AS t WHERE t.number = $1`, number)
	var s sql.NullString
	if scanErr := row.Scan(&s); scanErr != nil {
		return false
	}
	if s.String != "" || s.Valid {
		return true
	}
	return false
}
func (r *Repository) Add(ctx context.Context, req *BaseTrackReq, isContainer bool) error {
	for _, v := range req.Numbers {
		if !r.checkNumber(ctx, v) {
			_, err := r.db.ExecContext(ctx, `INSERT INTO "tasks" (number,user_id,country,time,emails,is_container,email_subject) VALUES ($1,$2,$3,$4,$5,$6,$7)`, v, req.UserId, req.Country, req.Time, pq.Array(req.Emails), isContainer, req.EmailMessageSubject)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (r *Repository) GetAll(ctx context.Context) ([]*TrackingTask, error) {
	var res []*TrackingTask
	rows, err := r.db.QueryContext(ctx, `SELECT number,user_id,country,time,emails,is_container,email_subject FROM "tasks"`)
	if err != nil {
		return res, err
	}
	var i = 0
	for rows.Next() {
		i++
		var oneItem TrackingTask
		if scanErr := rows.Scan(&oneItem.Number, &oneItem.UserId, &oneItem.Country, &oneItem.Time, pq.Array(&oneItem.Emails), &oneItem.IsContainer, &oneItem.EmailMessageSubject); scanErr != nil {
			return res, scanErr
		}
		res = append(res, &oneItem)
	}
	if i == 0 {
		return res, &NoTasksError{}
	}
	return res, nil
}
func (r *Repository) Update(ctx context.Context, req *BaseTrackReq, isContainer bool) error {
	for _, v := range req.Numbers {
		if _, err := r.db.ExecContext(ctx, `UPDATE "tasks" SET number = $1 ,user_id = $2 ,country = $3,time = $4,emails = $5, is_container = $6,email_subject = $7`, v, req.UserId, req.Country, req.Time, pq.Array(req.Emails), isContainer, req.EmailMessageSubject); err != nil {
			return err
		}
	}
	return nil
}
func (r *Repository) GetByNumber(ctx context.Context, number string) (*TrackingTask, error) {
	row := r.db.QueryRowContext(ctx, `SELECT number,user_id,country,time,emails,is_container,email_subject FROM "tasks" AS t WHERE t.number = $1`, number)
	var task TrackingTask
	if err := row.Scan(&task.Number, &task.UserId, &task.Country, &task.Time, pq.Array(&task.Emails), &task.IsContainer, &task.EmailMessageSubject); err != nil {
		return &TrackingTask{}, err
	}
	return &task, nil
}

func (r *Repository) Delete(ctx context.Context, numbers []string) error {
	for _, number := range numbers {
		if _, err := r.db.ExecContext(ctx, `DELETE FROM "tasks" AS t WHERE t.number = $1`, number); err != nil {
			return err
		}
	}
	return nil
}
func (r *Repository) ChangeEmailMessageSubject(ctx context.Context, number, newSubject string) error {
	_, err := r.db.ExecContext(ctx, `UPDATE "tasks" AS t SET email_subject = $1 WHERE t.number = $2`, newSubject, number)
	return err
}
