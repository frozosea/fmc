package company

import (
	"context"
	"database/sql"
)

type IRepository interface {
	Add(ctx context.Context, contact BaseCompany) error
	GetAll(ctx context.Context) ([]*Company, error)
	Update(ctx context.Context, id int, contact *BaseCompany) error
	Delete(ctx context.Context, id int) error
}
type Repository struct {
	db *sql.DB
}

func (repo *Repository) Add(ctx context.Context, contact BaseCompany) error {
	_, err := repo.db.ExecContext(ctx, `INSERT INTO "company" (url,email,agent_name,phone_number,company_name) VALUES ($1,$2,$3,$4,$5)`, contact.Url, contact.Email, contact.Name, contact.PhoneNumber, contact.Name)
	return err
}
func (repo *Repository) GetAll(ctx context.Context) ([]*Company, error) {
	var contacts []*Company
	rows, err := repo.db.QueryContext(ctx, `SELECT * FROM "company"`)
	if err != nil {
		return contacts, err
	}
	defer rows.Close()
	for rows.Next() {
		oneContact := new(Company)
		if scanErr := rows.Scan(&oneContact.Id, &oneContact.Url, &oneContact.Email, &oneContact.Name, &oneContact.PhoneNumber, &oneContact.Name); scanErr != nil {
			return contacts, scanErr
		}
		contacts = append(contacts, oneContact)

	}
	return contacts, nil
}
func (repo *Repository) Update(ctx context.Context, id int, company *BaseCompany) error {
	_, err := repo.db.ExecContext(ctx,
		`UPDATE "company" AS c SET 
                          url = $1, 
                          email = $2,
                          agent_name = $3,
                          phone_number = $4, 
                          company_name = $5 WHERE c.id = $6`,
		company.Url,
		company.Email,
		company.Name,
		company.PhoneNumber,
		company.Name,
		id)
	return err
}
func (repo *Repository) Delete(ctx context.Context, id int) error {
	if _, err := repo.db.ExecContext(ctx, `DELETE FROM "company" AS c WHERE c.id = $1`, id); err != nil {
		return err
	}
	return nil
}
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
