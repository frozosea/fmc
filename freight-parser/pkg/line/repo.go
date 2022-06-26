package line

import (
	"context"
	"database/sql"
)

type IRepository interface {
	Add(ctx context.Context, object AddRepoLine) error
	GetAll(ctx context.Context) ([]*Line, error)
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll(ctx context.Context) ([]*Line, error) {
	var allLines []*Line
	rows, err := r.db.QueryContext(ctx, `SELECT * FROM "lines"`)
	if err != nil {
		return allLines, err
	}
	defer rows.Close()
	for rows.Next() {
		oneLine := new(Line)
		if scanErr := rows.Scan(&oneLine.Id, &oneLine.Scac, &oneLine.FullName, &oneLine.ImageUrl); scanErr != nil {
			return allLines, scanErr
		}
		allLines = append(allLines, oneLine)
	}
	return allLines, nil
}
func (r *repository) Add(ctx context.Context, object AddRepoLine) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO "lines" (scac,full_name,image_url) VALUES ($1,$2,$3)`, object.Scac, object.FullName, object.ImageUrl)
	return err
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}
