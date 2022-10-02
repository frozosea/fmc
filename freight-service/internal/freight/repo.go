package freight

import (
	"context"
	"database/sql"
)

type IRepository interface {
	Get(ctx context.Context, freight GetFreight) ([]BaseFreight, error)
	GetAll(ctx context.Context) ([]BaseFreight, error)
	Add(ctx context.Context, freight AddFreight) error
	Update(ctx context.Context, id int, freight *AddFreight) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) Get(ctx context.Context, freight GetFreight) ([]BaseFreight, error) {
	var freights []BaseFreight
	rows, err := r.db.QueryContext(ctx, `SELECT 
       freight.id as price_id,
       from_city.ru_full_name as from_city_ru_full_name,
       from_city.en_full_name as from_city_en_full_name,
       from_city.id as from_city_id,
       from_city_country.id as from_city_country_id,
       from_city_country.ru_full_name as from_city_country_ru_full_name,
       from_city_country.en_full_name as from_city_country_en_full_name,
       to_city.ru_full_name as to_city_ru_full_name,
       to_city.en_full_name as to_city_en_full_name,
       to_city_country.id as to_city_country_id,
       to_city_country.ru_full_name as to_city_country_ru_full_name,
       to_city_country.en_full_name as to_city_country_en_full_name,
	   to_city.id as to_city_id,
       container.full_name as container_type,
       container.id as container_type_id,
       freight.usd_price,
       freight.from_date as from_date,
       freight.expires_date as expires_date,
       company.url as contact_url,
       company.email,
       company.name,
       company.phone_number	
	   FROM "price" AS freight 
	    JOIN "cities" AS from_city ON from_city.id = freight.from_city_id 
	    JOIN "cities" AS to_city ON to_city.id = freight.to_city_id
	    JOIN "countries" AS from_city_country ON from_city.country_id = from_city_country.id
	    JOIN "countries" AS to_city_country ON to_city.country_id = to_city_country.id
	    JOIN "container" AS container ON container.id = freight.container_id 
	    JOIN "company" AS company ON company.id = freight.contact_id
		WHERE from_city.id = $1 AND to_city.id = $2 AND freight.expires_date <= NOW()::DATE
		AND container.id = $3
		ORDER BY freight.usd_price
		LIMIT $4
`, freight.FromCityId, freight.ToCityId, freight.ContainerTypeId, freight.Limit)
	defer rows.Close()
	if err != nil {
		return freights, err
	}
	for rows.Next() {
		baseFreight := new(BaseFreight)
		if scanErr := rows.Scan(&baseFreight.Id,
			&baseFreight.FromCity.RuFullName,
			&baseFreight.FromCity.EnFullName,
			&baseFreight.FromCity.Id,
			&baseFreight.FromCity.Country.Id,
			&baseFreight.FromCity.Country.RuFullName,
			&baseFreight.FromCity.Country.EnFullName,
			&baseFreight.ToCity.RuFullName,
			&baseFreight.ToCity.EnFullName,
			&baseFreight.ToCity.Country.Id,
			&baseFreight.ToCity.Country.RuFullName,
			&baseFreight.ToCity.Country.EnFullName,
			&baseFreight.ToCity.Id,
			&baseFreight.Container.Type,
			&baseFreight.Container.Id,
			&baseFreight.UsdPrice,
			&baseFreight.FromDate,
			&baseFreight.ExpiresDate,
			&baseFreight.Company.Url,
			&baseFreight.Company.Email,
			&baseFreight.Company.Name,
			&baseFreight.Company.PhoneNumber); scanErr != nil {
			return freights, scanErr
		}
		freights = append(freights, *baseFreight)
	}
	return freights, nil

}
func (r *repository) GetAll(ctx context.Context) ([]BaseFreight, error) {
	var freights []BaseFreight
	rows, err := r.db.QueryContext(ctx, `SELECT 
       freight.id as price_id,
       from_city.ru_full_name as from_city_ru_full_name,
       from_city.en_full_name as from_city_en_full_name,
       from_city.id as from_city_id,
       from_city_country.id as from_city_country_id,
       from_city_country.ru_full_name as from_city_country_ru_full_name,
       from_city_country.en_full_name as from_city_country_en_full_name,
       to_city.ru_full_name as to_city_ru_full_name,
       to_city.en_full_name as to_city_en_full_name,
       to_city_country.id as to_city_country_id,
       to_city_country.ru_full_name as to_city_country_ru_full_name,
       to_city_country.en_full_name as to_city_country_en_full_name,
	   to_city.id as to_city_id,
       container.full_name as container_type,
       container.id as container_type_id,
       freight.usd_price,
       freight.from_date as from_date,
       freight.expires_date as expires_date,
       company.id as contact_id,
       company.url as contact_url,
       company.email,
       company.agent_name,
       company.phone_number
	FROM "price" AS freight 
	    JOIN "cities" AS from_city ON from_city.id = freight.from_city_id 
	    JOIN "cities" AS to_city ON to_city.id = freight.to_city_id
	    JOIN "countries" AS from_city_country ON from_city.country_id = from_city_country.id
	    JOIN "countries" AS to_city_country ON to_city.country_id = to_city_country.id
	    JOIN "container" AS container ON container.id = freight.container_id 
	    JOIN "company" AS company ON company.id = freight.contact_id
		ORDER BY freight.usd_price
`)
	defer rows.Close()
	if err != nil {
		return freights, err
	}
	for rows.Next() {
		baseFreight := new(BaseFreight)
		if scanErr := rows.Scan(&baseFreight.Id,
			&baseFreight.FromCity.RuFullName,
			&baseFreight.FromCity.EnFullName,
			&baseFreight.FromCity.Id,
			&baseFreight.FromCity.Country.Id,
			&baseFreight.FromCity.Country.RuFullName,
			&baseFreight.FromCity.Country.EnFullName,
			&baseFreight.ToCity.RuFullName,
			&baseFreight.ToCity.EnFullName,
			&baseFreight.ToCity.Country.Id,
			&baseFreight.ToCity.Country.RuFullName,
			&baseFreight.ToCity.Country.EnFullName,
			&baseFreight.ToCity.Id,
			&baseFreight.Container.Type,
			&baseFreight.Container.Id,
			&baseFreight.UsdPrice,
			&baseFreight.FromDate,
			&baseFreight.ExpiresDate,
			&baseFreight.Company.Id,
			&baseFreight.Company.Url,
			&baseFreight.Company.Email,
			&baseFreight.Company.Name,
			&baseFreight.Company.PhoneNumber); scanErr != nil {
			return freights, scanErr
		}
		freights = append(freights, *baseFreight)
	}
	return freights, nil
}
func (r *repository) Add(ctx context.Context, freight AddFreight) error {
	tx, getTxErr := r.db.BeginTx(ctx, nil)
	if getTxErr != nil {
		return getTxErr
	}
	defer tx.Rollback()
	_, execErr := tx.ExecContext(ctx, `
		INSERT INTO "price" (
			from_city_id,
			usd_price,
			container_id,
			contact_id,
			to_city_id,
			from_date,
			expires_date)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		`,
		freight.FromCityId,
		freight.UsdPrice,
		freight.ContainerTypeId,
		freight.ContactId,
		freight.ToCityId,
		freight.FromDate,
		freight.ExpiresDate)
	if execErr != nil {
		return execErr
	}
	return tx.Commit()
}
func (r *repository) Update(ctx context.Context, id int, freight *AddFreight) error {
	_, err := r.db.ExecContext(ctx, `UPDATE "price" AS p SET 
                        from_city_id = $1, 
                        to_city_id = $2,
                        container_id = $3,
                        contact_id = $4, 
                        from_date = $5, 
                        to_date = $6, 
                        usd_price = $7
						WHERE p.id = $8`,
		freight.FromCityId,
		freight.ToCityId,
		freight.ContainerTypeId,
		freight.ContactId,
		freight.FromDate,
		freight.ExpiresDate,
		freight.UsdPrice,
		id)
	return err
}
func (r *repository) Delete(ctx context.Context, id int) error {
	if _, err := r.db.ExecContext(ctx, `DELETE "freights" AS f WHERE f.id = $1`, id); err != nil {
		return err
	}
	return nil
}
func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}
