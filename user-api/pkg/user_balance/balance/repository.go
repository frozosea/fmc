package balance

import (
	"context"
	"database/sql"
	"errors"
)

type IRepository interface {
	Add(ctx context.Context, userId int, value float64) error
	Sub(ctx context.Context, userId int, value float64) error
	Get(ctx context.Context, userId int) (float64, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Add(ctx context.Context, userId int, value float64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE "balance" SET value = value + $1 WHERE user_id = $2`, value, userId); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.ExecContext(ctx, `INSERT INTO "balance_transaction" AS tr (user_id,value,type) VALUES ($1,$2,$3)`, userId, value, 1); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *Repository) Sub(ctx context.Context, userId int, value float64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE "balance" SET value = value - $1 WHERE user_id = $2`, value, userId); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.ExecContext(ctx, `INSERT INTO "balance_transaction" AS tr (user_id,value,type) VALUES ($1,$2,$3)`, userId, value, 2); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *Repository) getBalance(ctx context.Context, userId int) (float64, error) {
	row := r.db.QueryRowContext(ctx, `SELECT b.value FROM "balance" AS b WHERE b.user_id = $1`, userId)
	if err := row.Err(); err != nil {
		return -1, err
	}
	var balance sql.NullFloat64
	if err := row.Scan(&balance); err != nil {
		return -1, err
	}

	if !balance.Valid {
		return -1, errors.New("balance is unreadable")
	}

	return balance.Float64, nil
}

func (r *Repository) Get(ctx context.Context, userId int) (float64, error) {
	row := r.db.QueryRowContext(ctx, `SELECT b.value FROM "balance" AS b WHERE b.user_id = $1`, userId)
	if err := row.Err(); err != nil {
		return -1, err
	}
	var balance sql.NullFloat64
	if err := row.Scan(&balance); err != nil {
		return -1, err
	}

	if !balance.Valid {
		return -1, errors.New("balance is unreadable")
	}

	return balance.Float64, nil
}
