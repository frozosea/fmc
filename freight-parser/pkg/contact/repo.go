package contact

import (
	"context"
	"database/sql"
)

type IRepository interface {
	Add(ctx context.Context, contact BaseContact) error
	GetAll(ctx context.Context) ([]*Contact, error)
}
type repository struct {
	db *sql.DB
}

func (repo *repository) Add(ctx context.Context, contact BaseContact) error {
	_, err := repo.db.ExecContext(ctx, `INSERT INTO "contact" (url,email,agent_name,phone_number) VALUES ($1,$2,$3,$4)`, contact.Url, contact.Email, contact.AgentName, contact.PhoneNumber)
	return err
}
func (repo *repository) GetAll(ctx context.Context) ([]*Contact, error) {
	var contacts []*Contact
	rows, err := repo.db.QueryContext(ctx, `SELECT * FROM "contact"`)
	if err != nil {
		return contacts, err
	}
	defer rows.Close()
	for rows.Next() {
		oneContact := new(Contact)
		if scanErr := rows.Scan(&oneContact.Id, &oneContact.Url, &oneContact.Email, &oneContact.AgentName, &oneContact.PhoneNumber); scanErr != nil {
			return contacts, scanErr
		}
		contacts = append(contacts, oneContact)

	}
	return contacts, nil
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}
