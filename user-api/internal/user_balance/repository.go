package user_balance

import (
	"context"
	"database/sql"
	"fmt"
)

type IRepository interface {
	GetAllNumbersOnTrack(ctx context.Context, userId int64) (int64, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllNumbersOnTrack(ctx context.Context, userId int64) (int64, error) {
	var containerOnTrackQuantity, billsOnTrackQuantity sql.NullInt64

	err := r.db.QueryRowContext(ctx, `
	SELECT (SELECT COUNT(*) FROM "containers" AS c WHERE c.is_on_track AND c.user_id = $1) AS containers_on_track,
  (SELECT COUNT(*) FROM "bill_numbers" AS c WHERE c.is_on_track AND c.user_id = $2) AS bills_on_track
`, userId, userId).Scan(&containerOnTrackQuantity, &billsOnTrackQuantity)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	if !containerOnTrackQuantity.Valid || !billsOnTrackQuantity.Valid {
		return 0, nil
	}

	return containerOnTrackQuantity.Int64 + billsOnTrackQuantity.Int64, nil
}
