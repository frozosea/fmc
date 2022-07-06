package city

import (
	"context"
	"database/sql"
)

type ICityRepository interface {
	Add(ctx context.Context, city BaseCity) error
	GetAll(ctx context.Context) ([]*City, error)
}
type repository struct {
	db *sql.DB
}

func (repo *repository) Add(ctx context.Context, city BaseCity) error {
	_, err := repo.db.ExecContext(ctx, `INSERT INTO "cities" (unlocode, full_name) VALUES ($1,$2)`, city.Unlocode, city.FullName)
	return err
}

func (repo *repository) GetAll(ctx context.Context) ([]*City, error) {
	var cities []*City
	rows, err := repo.db.QueryContext(ctx, `SELECT * FROM "cities"`)
	if err != nil {
		return cities, err
	}
	defer rows.Close()
	for rows.Next() {
		oneCity := new(City)
		if scanErr := rows.Scan(&oneCity.Id, &oneCity.FullName, &oneCity.Unlocode); scanErr != nil {
			return cities, scanErr
		}
		cities = append(cities, oneCity)
	}
	return cities, nil
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}
