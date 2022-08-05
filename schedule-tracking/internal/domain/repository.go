package domain

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"schedule-tracking/pkg/util"
)

type NoTasksError struct{}

func (n *NoTasksError) Error() string {
	return "no tasks in table"
}

type IRepository interface {
	Add(ctx context.Context, req *BaseTrackReq, isContainer bool) error
	GetAll(ctx context.Context) ([]*TrackingTask, error)
	GetByNumber(ctx context.Context, number string) (*TrackingTask, error)
	AddEmails(ctx context.Context, numbers []string, emails []string) error
	DeleteEmail(ctx context.Context, number string, email string) error
	UpdateTime(ctx context.Context, numbers []string, newTime string) error
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
			_, err := r.db.ExecContext(ctx, `INSERT INTO "tasks" (number,user_id,country,time,emails,is_container) VALUES ($1,$2,$3,$4,$5,$6)`, v, req.UserId, req.Country, req.Time, pq.Array(req.Emails), isContainer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (r *Repository) GetAll(ctx context.Context) ([]*TrackingTask, error) {
	var res []*TrackingTask
	rows, err := r.db.QueryContext(ctx, `SELECT number,user_id,country,time,emails,is_container FROM "tasks"`)
	if err != nil {
		return res, err
	}
	var i = 0
	for rows.Next() {
		i++
		var oneItem TrackingTask
		if scanErr := rows.Scan(&oneItem.Number, &oneItem.UserId, &oneItem.Country, &oneItem.Time, pq.Array(&oneItem.Emails), &oneItem.IsContainer); scanErr != nil {
			return res, scanErr
		}
		res = append(res, &oneItem)
	}
	if i == 0 {
		return res, &NoTasksError{}
	}
	return res, nil
}
func (r *Repository) GetByNumber(ctx context.Context, number string) (*TrackingTask, error) {
	row := r.db.QueryRowContext(ctx, `SELECT number,user_id,country,time,emails,is_container FROM "tasks" AS t WHERE t.number = $1`, number)
	var task TrackingTask
	if err := row.Scan(&task.Number, &task.UserId, &task.Country, &task.Time, pq.Array(&task.Emails), &task.IsContainer); err != nil {
		return &TrackingTask{}, err
	}
	return &task, nil
}
func (r *Repository) getEmails(ctx context.Context, number string) ([]string, error) {
	var s []string
	if err := r.db.QueryRowContext(ctx, `SELECT t.emails AS t FROM "tasks" AS t WHERE t.number = $1`, number).Scan(pq.Array(&s)); err != nil {
		return s, err
	}
	return s, nil
}
func (r *Repository) AddEmails(ctx context.Context, numbers []string, emails []string) error {
	for _, v := range numbers {
		_, err := r.db.ExecContext(ctx, `UPDATE "tasks" AS t SET emails = t.emails || $1 WHERE t.number = $2`, pq.Array(emails), v)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *Repository) DeleteEmail(ctx context.Context, number string, email string) error {
	oldEmailsFromDb, err := r.getEmails(ctx, number)
	if err != nil {
		return err
	}
	updatedEmails := util.Pop(oldEmailsFromDb, util.GetIndex(email, util.ConvertArgsToInterface(oldEmailsFromDb)...))
	_, addErr := r.db.ExecContext(ctx, `UPDATE "tasks" AS t SET emails = $1 WHERE t.number = $2`, pq.Array(updatedEmails), number)
	return addErr
}
func (r *Repository) UpdateTime(ctx context.Context, numbers []string, newTime string) error {
	for _, number := range numbers {
		_, err := r.db.ExecContext(ctx, `UPDATE "tasks" AS t SET time = $1 WHERE t.number = $2`, newTime, number)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *Repository) Delete(ctx context.Context, numbers []string) error {
	for _, number := range numbers {
		_, err := r.db.ExecContext(ctx, `DELETE FROM "tasks" AS t WHERE t.number = $1`, number)
		if err != nil {
			return err
		}
	}
	return nil
}
