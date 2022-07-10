package freight

import (
	"context"
	"database/sql"
)

type IRepository interface {
	Get(ctx context.Context, freight GetFreight) ([]BaseFreight, error)
	Add(ctx context.Context, freight AddFreight) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) Get(ctx context.Context, freight GetFreight) ([]BaseFreight, error) {
	var freights []BaseFreight
	rows, err := r.db.QueryContext(ctx, `SELECT 
       freight.id as price_id,
       from_city.full_name as from_city_full_name,
       from_city.id as from_city_id,
       from_city.unlocode as from_city_unlocode,
       to_city.full_name as to_city_full_name,
	   to_city.id as to_city_id,
       to_city.unlocode as to_city_unlocode,
       container.full_name as container_type,
       container.id as containert_type_id,
       freight.usd_price,
       lines.full_name as line_full_name,
       lines.id as line_id,
       lines.image_url as line_image_url,
       lines.scac as scac,
       freight.from_date as from_date,
       freight.expires_date as expires_date,
       contact.url as contact_url,
       contact.email,
       contact.agent_name,
       contact.phone_number
	FROM "price" AS freight JOIN "lines" AS lines ON lines.id = freight.line_id 
	    JOIN "cities" AS from_city ON from_city.id = freight.from_city_id 
	    JOIN "cities" AS to_city ON to_city.id = freight.to_city_id  
	    JOIN "container" AS container ON container.id = freight.container_id 
	    JOIN "contact" AS contact ON contact.id = freight.contact_id
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
		if scanErr := rows.Scan(&baseFreight.Id, &baseFreight.FromCity.FullName, &baseFreight.FromCity.Id, &baseFreight.FromCity.Unlocode, &baseFreight.ToCity.FullName, &baseFreight.ToCity.Id, &baseFreight.ToCity.Unlocode, &baseFreight.Type, &baseFreight.Id, &baseFreight.UsdPrice, &baseFreight.Line, &baseFreight.Id, &baseFreight.ImageUrl, &baseFreight.Scac, &baseFreight.FromDate, &baseFreight.ExpiresDate, &baseFreight.Contact.Url, &baseFreight.Contact.Email, &baseFreight.Contact.AgentName, &baseFreight.Contact.PhoneNumber); scanErr != nil {
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
			line_id,
			from_city_id,
			usd_price,
			container_id,
			contact_id,
			to_city_id,
			from_date,
			expires_date)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		`, freight.LineId, freight.FromCityId, freight.UsdPrice, freight.ContainerTypeId, freight.ContactId, freight.ToCityId, freight.FromDate, freight.ExpiresDate)
	if execErr != nil {
		return execErr
	}
	return tx.Commit()
}
func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}
