package freight

import (
	"context"
	"database/sql"
	"fmc-newest/pkg/domain"
)

func NewGetFreight(fromCity string, toCity string, containerType string, limit int64) domain.GetFreight {
	return domain.GetFreight{FromCity: fromCity, ToCity: toCity, ContainerType: containerType, Limit: limit}
}

type IFreightRepository interface {
	GetFreight(ctx context.Context, freight domain.GetFreight) ([]domain.BaseFreight, error)
	AddFreight(ctx context.Context, freight domain.AddFreight) error
}

type ICityRepository interface {
	AddCity(ctx context.Context, city domain.AddCity) error
	GetAllCities(ctx context.Context) ([]domain.GetCity, error)
}
type IContactRepository interface {
	AddContact(ctx context.Context, contact domain.Contact) error
	GetAllContacts(ctx context.Context) ([]domain.Contact, error)
}
type IContainerRepository interface {
	AddContainerType(ctx context.Context, containerType string) error
	GetAllContainers(ctx context.Context) ([]*domain.Container, error)
}

type FreightRepository struct {
	db *sql.DB
}

func (repo *FreightRepository) GetFreight(ctx context.Context, freight domain.GetFreight) ([]domain.BaseFreight, error) {
	var freights []domain.BaseFreight
	rows, err := repo.db.QueryContext(ctx, `select 
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
	from "price" as freight join "lines" as lines on lines.id = freight.line_id 
	    join "cities" as from_city on from_city.id = freight.from_city_id 
	    join "cities" as to_city on to_city.id = freight.to_city_id  
	    join "container" as container on container.id = freight.container_id 
	    join "contact" as contact on contact.id = freight.contact_id
		where from_city.full_name = $1 and to_city.full_name = $2 and freight.expires_date <= NOW()::date
		and container.full_name = $3
		order by freight.usd_price
		limit $4
`, freight.FromCity, freight.ToCity, freight.ContainerType, freight.Limit)
	defer rows.Close()
	if err != nil {
		return freights, err
	}
	for rows.Next() {
		baseFreight := new(domain.BaseFreight)
		if scanErr := rows.Scan(&baseFreight.Id, &baseFreight.FromCity.CityName, &baseFreight.FromCity.CityId, &baseFreight.FromCity.CityUnlocode, &baseFreight.ToCity.CityName, &baseFreight.ToCity.CityId, &baseFreight.ToCity.CityUnlocode, &baseFreight.ContainerType, &baseFreight.ContainerTypeId, &baseFreight.UsdPrice, &baseFreight.Line, &baseFreight.LineId, &baseFreight.LineImage, &baseFreight.Scac, &baseFreight.FromDate, &baseFreight.ExpiresDate, &baseFreight.Contact.Url, &baseFreight.Contact.Email, &baseFreight.Contact.AgentName, &baseFreight.Contact.PhoneNumber); scanErr != nil {
			return freights, scanErr
		}
		freights = append(freights, *baseFreight)
	}
	return freights, nil

}

func (repo *FreightRepository) AddFreight(ctx context.Context, freight domain.AddFreight) error {
	tx, getTxErr := repo.db.BeginTx(ctx, nil)
	if getTxErr != nil {
		return getTxErr
	}
	defer tx.Rollback()
	_, execErr := tx.ExecContext(ctx, `
		insert into "price" (
			line_id,
			from_city_id,
			usd_price,
			container_id,
			contact_id,
			to_city_id,
			from_date,
			expires_date)
		values ($1,$2,$3,$4,$5,$6,$7,$8)
		`, freight.LineId, freight.FromCityId, freight.UsdPrice, freight.ContainerTypeId, freight.ContactId, freight.ToCityId, freight.FromDate, freight.ExpiresDate)
	if execErr != nil {
		return execErr
	}
	return tx.Commit()
}

type CityRepository struct {
	db *sql.DB
}

func (repo *CityRepository) AddCity(ctx context.Context, city domain.AddCity) error {
	_, err := repo.db.ExecContext(ctx, `insert into "cities" (unlocode, full_name) values ($1,$2)`, city.Unlocode, city.CityFullName)
	return err
}

func (repo *CityRepository) GetAllCities(ctx context.Context) ([]domain.GetCity, error) {
	var cities []domain.GetCity
	rows, err := repo.db.QueryContext(ctx, `select * from "cities"`)
	if err != nil {
		return cities, err
	}
	defer rows.Close()
	for rows.Next() {
		oneCity := new(domain.GetCity)
		if scanErr := rows.Scan(&oneCity.Id, &oneCity.CityFullName, &oneCity.Unlocode); scanErr != nil {
			return cities, scanErr
		}
		cities = append(cities, *oneCity)
	}
	return cities, nil
}

type ContactRepository struct {
	db *sql.DB
}

func (repo *ContactRepository) AddContact(ctx context.Context, contact domain.Contact) error {
	_, err := repo.db.ExecContext(ctx, `insert into "contact" (url,email,agent_name,phone_number) values ($1,$2,$3,$4)`, contact.Url, contact.Email, contact.AgentName, contact.PhoneNumber)
	return err
}
func (repo *ContactRepository) GetAllContacts(ctx context.Context) ([]domain.Contact, error) {
	var contacts []domain.Contact
	rows, err := repo.db.QueryContext(ctx, `select * from "contact"`)
	if err != nil {
		return contacts, err
	}
	defer rows.Close()
	for rows.Next() {
		oneContact := new(domain.Contact)
		if scanErr := rows.Scan(&oneContact.Id, &oneContact.Url, &oneContact.Email, &oneContact.AgentName, &oneContact.PhoneNumber); scanErr != nil {
			return contacts, scanErr
		}
		contacts = append(contacts, *oneContact)

	}
	return contacts, nil
}

type ContainerRepository struct {
	db *sql.DB
}

func (repo *ContainerRepository) AddContainerType(ctx context.Context, containerType string) error {
	_, err := repo.db.ExecContext(ctx, `insert into "container" (full_name) values ($1)`, containerType)
	return err
}
func (repo *ContainerRepository) GetAllContainers(ctx context.Context) ([]*domain.Container, error) {
	var containers []*domain.Container
	rows, err := repo.db.QueryContext(ctx, `select * from "container"`)
	if err != nil {
		return containers, err
	}
	defer rows.Close()
	for rows.Next() {
		oneContainer := new(domain.Container)
		if scanErr := rows.Scan(&oneContainer.ContainerTypeId, &oneContainer.ContainerType); scanErr != nil {
			return containers, scanErr
		}
		containers = append(containers, oneContainer)
	}
	return containers, nil
}
