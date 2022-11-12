package feedback

import (
	"context"
	"database/sql"
)

type IRepository interface {
	Save(ctx context.Context, fb *Feedback) error
	GetByEmail(ctx context.Context, email string) ([]*Feedback, error)
	GetAll(ctx context.Context) ([]*Feedback, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(ctx context.Context, fb *Feedback) error {
	if _, err := r.db.ExecContext(ctx, `INSERT INTO "feedback" (email, message) VALUES ($1,$2)`, fb.Email, fb.Message); err != nil {
		return err
	}
	return nil
}
func (r *Repository) queryFeedbacks(ctx context.Context, queryRow string, args ...interface{}) ([]*Feedback, error) {
	rows, err := r.db.QueryContext(ctx, queryRow, args...)
	if err != nil {
		return nil, err
	}
	var feedbacks []*Feedback
	for rows.Next() {
		var oneFb *Feedback
		if err := rows.Scan(&oneFb.Email, oneFb.Message); err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, oneFb)
	}
	return feedbacks, nil
}
func (r *Repository) GetByEmail(ctx context.Context, email string) ([]*Feedback, error) {
	return r.queryFeedbacks(ctx, `SELECT email,message FROM "feedback" AS f WHERE f.email = $1`, email)
}

func (r *Repository) GetAll(ctx context.Context) ([]*Feedback, error) {
	return r.queryFeedbacks(ctx, `SELECT email,message FROM "feedback"`)
}
