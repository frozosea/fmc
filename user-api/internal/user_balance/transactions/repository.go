package transactions

import (
	"context"
	"database/sql"
	"time"
)

type IRepository interface {
	Add(ctx context.Context, transactionId int, number string) error
	GetTransactionsByDates(ctx context.Context, userId int, fromTimestamp, toTimestamp time.Time) ([]*Transaction, error)
	GetTransactionsByNumber(ctx context.Context, userId int, number string) (*Transaction, error)
	GetInfoByMonth(ctx context.Context, userId int) ([]*Transaction, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Add(ctx context.Context, transactionId int, number string) error {
	if _, err := r.db.ExecContext(ctx, `INSERT INTO "number_transaction" AS tr (transaction_id,number) VALUES ($1,$2)`, transactionId, number); err != nil {
		return err
	}
	return nil
}

// SELECT b_tr.*,num_tr.* FROM "balance_transaction" AS b_tr LEFT JOIN "number_transaction" AS num_tr ON num_tr.transaction_id = b_tr.id WHERE b_tr.user_id = 19 AND (b_tr.created_at > '2023-05-23 17:01:23.515615+10' AND b_tr.created_at < '2022-05-23 17:01:23.515615+10') AND b_tr.type = 2
func (r *Repository) GetTransactionsByDates(ctx context.Context, userId int, fromTimestamp, toTimestamp time.Time) ([]*Transaction, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT b_tr.*,num_tr.* FROM "balance_transaction" AS b_tr LEFT JOIN "number_transaction" AS num_tr ON num_tr.transaction_id = b_tr.id WHERE b_tr.user_id = $1 AND (b_tr.created_at > $2 AND b_tr.created_at < $3) AND b_tr.type = 2`, userId, fromTimestamp, toTimestamp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*Transaction

	var nullNumberStr sql.NullString

	for rows.Next() {
		var t *Transaction
		if err := rows.Scan(&t.ID, t.UserID, t.Value, t.Type, t.TimeStamp.Format(time.RFC3339), &nullNumberStr); err != nil {
			return nil, err
		}
		if nullNumberStr.Valid {
			t.Number = nullNumberStr.String
		}
	}

	return transactions, nil
}

func (r *Repository) GetTransactionsByNumber(ctx context.Context, userId int, number string) (*Transaction, error) {
	return nil, nil
}
func (r *Repository) GetInfoByMonth(ctx context.Context, userId int) ([]*Transaction, error) {
	return nil, nil
}
