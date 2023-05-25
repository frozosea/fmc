package balance

import (
	"context"
	"database/sql"
	"errors"
	"user-api/internal/domain"
)

type IRepository interface {
	Add(ctx context.Context, userId int64, value float64) (*domain.Transaction, error)
	Sub(ctx context.Context, userId int64, value float64) (*domain.Transaction, error)
	Get(ctx context.Context, userId int64) (float64, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Add(ctx context.Context, userId int64, value float64) (*domain.Transaction, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE "balance" SET value = value + $1 WHERE user_id = $2`, value, userId); err != nil {
		tx.Rollback()
		return nil, err
	}

	var tr *domain.Transaction
	if err := tx.QueryRowContext(ctx, `INSERT INTO "balance_transaction" AS tr (user_id,value,type) VALUES ($1,$2,$3) RETURNING *;`, userId, value, 1).Scan(
		&tr.ID, &tr.UserID, &tr.Value, &tr.Type, &tr.TimeStamp,
	); err != nil {
		tx.Rollback()
		return nil, err
	}
	return tr, tx.Commit()
}

func (r *Repository) Sub(ctx context.Context, userId int64, value float64) (*domain.Transaction, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE "balance" SET value = value - $1 WHERE user_id = $2`, value, userId); err != nil {
		tx.Rollback()
		return nil, err
	}
	var tr *domain.Transaction
	if err := tx.QueryRowContext(ctx, `INSERT INTO "balance_transaction" AS tr (user_id,value,type) VALUES ($1,$2,$3) RETURNING *;`, userId, value, 2).Scan(
		&tr.ID, &tr.UserID, &tr.Value, &tr.Type, &tr.TimeStamp,
	); err != nil {
		tx.Rollback()
		return nil, err
	}
	return tr, tx.Commit()
}

func (r *Repository) getBalance(ctx context.Context, userId int64) (float64, error) {
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

func (r *Repository) Get(ctx context.Context, userId int64) (float64, error) {
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
