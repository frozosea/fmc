package sklu

import (
	"context"
	"database/sql"
	"errors"
)

type IRepository interface {
	GetFullName(ctx context.Context, unlocode string) (string, error)
	Add(ctx context.Context, unlocode, fullName string) error
}
type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetFullName(ctx context.Context, unlocode string) (string, error) {
	query := r.db.QueryRowContext(ctx, `SELECT s.fullname FROM "unlocodes" AS s WHERE s.unlocode = $1`, unlocode)
	var nullString *sql.NullString
	if err := query.Scan(&nullString); err != nil {
		return "", err
	}
	if nullString.Valid {
		return nullString.String, nil
	}
	return "", errors.New("no scac by your param")
}

func (r *Repository) Add(ctx context.Context, unlocode, fullName string) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO "unlocodes" (unlocode,fullname) VALUES ($1,$2)`, unlocode, fullName)
	return err
}
