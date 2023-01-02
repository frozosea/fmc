package scac_accessory

import (
	"context"
	"database/sql"
	"errors"
)

type IRepository interface {
	Get(ctx context.Context, number string) (string, error)
	Add(ctx context.Context, scac, number string) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Get(ctx context.Context, number string) (string, error) {
	query := r.db.QueryRowContext(ctx, `SELECT s.scac FROM "scac_accessory" AS s WHERE s.number = $1`, number)
	var nullString sql.NullString
	if err := query.Scan(&nullString); err != nil {
		return "", err
	}
	if nullString.Valid {
		return nullString.String, nil
	}
	return "", errors.New("no scac by your param")
}

func (r *Repository) Add(ctx context.Context, scac, number string) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO "scac_accessory" (scac,number) VALUES ($1,$2) ON CONFLICT DO NOTHING`, scac, number)
	return err
}
