package transactions

import (
	"context"
	"database/sql"
	"time"
)

type IRepository interface {
	Add(ctx context.Context, userId int, value float64, trType int) error
	GetByDates(ctx context.Context, userId int, fromTimestamp, toTimestamp time.Time) ([]*Transaction, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Add(ctx context.Context, userId int, value float64, trType int) error {
	if _, err := r.db.ExecContext(ctx, `INSERT INTO "balance_transaction" AS tr (user_id,value,type) VALUES ($1,$2,$3)`, userId, value, trType); err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetByDates(ctx context.Context, userId int, fromTimestamp, toTimestamp time.Time) ([]*Transaction, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT * FROM "balance_transaction" AS tr WHERE tr.user_id = $1 AND (tr.created_at > $2 AND tr.created_at < $3)`, userId, fromTimestamp, toTimestamp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*Transaction

	for rows.Next() {
		var t *Transaction
		if err := rows.Scan(&t.ID, t.UserID, t.Value, t.Type, t.TimeStamp.Format(time.RFC3339)); err != nil {
			return nil, err
		}
	}

	return transactions, nil
}
