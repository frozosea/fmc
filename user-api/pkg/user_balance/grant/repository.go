package grant

import (
	"context"
	"database/sql"
)

type IRepository interface {
	Add(ctx context.Context, userId int64, value float64) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Add(ctx context.Context, userId int64, value float64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO "balance" AS user_balance(user_id, value) VALUES ($1,$2)`, userId, value); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO "balance_transaction" AS tr (user_id,value,type) VALUES ($1,$2,$3)`, userId, value, 1); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
