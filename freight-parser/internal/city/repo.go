package city

import (
	"context"
	"database/sql"
)

type IRepository interface {
	Add(ctx context.Context, city *CountryWithId) error
	AddCountry(ctx context.Context, country BaseEntity) error
	UpdateCountry(ctx context.Context, id int, country *BaseEntity) error
	GetAllCountries(ctx context.Context) ([]*Country, error)
	GetAllCities(ctx context.Context) ([]*City, error)
	UpdateCity(ctx context.Context, id int, city *CountryWithId) error
	DeleteCity(ctx context.Context, id int) error
	DeleteCountry(ctx context.Context, id int) error
}
type Repository struct {
	db *sql.DB
}

func (repo *Repository) Add(ctx context.Context, city *CountryWithId) error {
	_, err := repo.db.ExecContext(ctx, `INSERT INTO "cities" (ru_full_name,en_full_name,country_id) VALUES ($1,$2,$3)`, city.RuFullName, city.EnFullName, city.CountryId)
	return err
}
func (repo *Repository) AddCountry(ctx context.Context, country BaseEntity) error {
	_, err := repo.db.ExecContext(ctx, `INSERT INTO "countries" (ru_full_name,en_full_name) VALUES ($1,$2,$3)`, country.RuFullName, country.EnFullName)
	return err
}
func (repo *Repository) GetAllCountries(ctx context.Context) ([]*Country, error) {
	var allCoutries []*Country
	rows, err := repo.db.QueryContext(ctx, `SELECT c.id,c.ru_full_name,c.en_full_name FROM "countries" AS c`)
	if err != nil {
		return allCoutries, err
	}
	for rows.Next() {
		country := new(Country)
		if err := rows.Scan(&country.Id, &country.RuFullName, &country.EnFullName); err != nil {
			return allCoutries, err
		}
		allCoutries = append(allCoutries, country)
	}
	return allCoutries, nil
}
func (repo *Repository) GetAllCities(ctx context.Context) ([]*City, error) {
	var cities []*City
	rows, err := repo.db.QueryContext(ctx, `SELECT c.id,c.ru_full_name,c.en_full_name, co.id,co.ru_full_name,co.en_full_name FROM "cities" AS c JOIN "countries" AS co ON c.country_id = co.id`)
	if err != nil {
		return cities, err
	}
	defer rows.Close()
	for rows.Next() {
		oneCity := new(City)
		if scanErr := rows.Scan(&oneCity.Id, &oneCity.RuFullName, &oneCity.EnFullName, &oneCity.Country.Id, &oneCity.Country.RuFullName, &oneCity.Country.EnFullName); scanErr != nil {
			return cities, scanErr
		}
		cities = append(cities, oneCity)
	}
	return cities, nil
}
func (repo *Repository) UpdateCountry(ctx context.Context, id int, country *BaseEntity) error {
	_, err := repo.db.ExecContext(ctx, `
			UPDATE "countries" AS c SET 
                            ru_full_name = $1,
                            en_full_name= $2
			WHERE c.id = $3`,
		country.RuFullName,
		country.EnFullName,
		id)
	return err
}
func (repo *Repository) UpdateCity(ctx context.Context, id int, city *CountryWithId) error {
	_, err := repo.db.ExecContext(ctx, `
			UPDATE "cities" AS c SET 
                         ru_full_name = $1 ,
                         en_full_name = $2 , 
                         country_id = $3 
			WHERE c.id = $4`,
		city.RuFullName,
		city.EnFullName,
		city.CountryId,
		id)
	return err
}
func (repo *Repository) DeleteCity(ctx context.Context, id int) error {
	if _, err := repo.db.ExecContext(ctx, `DELETE "cities" AS c WHERE c.id = $1`); err != nil {
		return err
	}
	return nil
}
func (repo *Repository) DeleteCountry(ctx context.Context, id int) error {
	if _, err := repo.db.ExecContext(ctx, `DELETE "countries" AS c WHERE c.id = $1`); err != nil {
		return err
	}
	return nil
}
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
