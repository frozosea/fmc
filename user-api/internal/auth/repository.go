package auth

import (
	"context"
	"database/sql"
	"fmt"
	"user-api/internal/domain"
)

type IRepository interface {
	Register(ctx context.Context, user *domain.RegisterUser) (int64, error)
	Login(ctx context.Context, user *domain.User) (int, error)
	CheckAccess(ctx context.Context, userId int) (bool, error)
	GetUserId(ctx context.Context, email string) (int, error)
	CheckEmailExist(ctx context.Context, email string) (bool, error)
	SetNewPassword(ctx context.Context, userId int, newHashPassword string) error
}
type AlreadyRegisterError struct{}

func NewAlreadyRegisterError() *AlreadyRegisterError {
	return &AlreadyRegisterError{}
}
func (a *AlreadyRegisterError) Error() string {
	return "Cannot register with these parameters"
}

type InvalidUserError struct{}

func (i *InvalidUserError) Error() string {
	return "Cannot find user with these parameters"
}

type Repository struct {
	db   *sql.DB
	hash IHash
}

func NewRepository(db *sql.DB, hash IHash) *Repository {
	return &Repository{db: db, hash: hash}
}

func (r *Repository) Register(ctx context.Context, user *domain.RegisterUser) (int64, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		return -1, err
	}

	var userId sql.NullInt64

	if err := tx.QueryRowContext(ctx, `INSERT INTO "user" (email, username, password) VALUES ($1,$2,$3) RETURNING id`,
		user.Email,
		user.Username,
		user.Password,
	).Scan(&userId); err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		return -1, NewAlreadyRegisterError()
	}

	if !userId.Valid {
		fmt.Println(err.Error())
		tx.Rollback()
		return -1, NewAlreadyRegisterError()
	}

	if _, err := tx.ExecContext(ctx,
		`INSERT INTO "company" AS c (
                            user_id,
                            company_full_name,
                            company_abbreviated_name,
                            inn,
                            ogrn,
                            legal_address,
							post_address,
							work_email
							) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		userId,
		user.CompanyFullName,
		user.CompanyAbbreviatedName,
		user.INN,
		user.OGRN,
		user.LegalAddress,
		user.PostAddress,
		user.WorkEmail,
	); err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		return -1, err
	}

	return userId.Int64, tx.Commit()
}
func (r *Repository) Login(ctx context.Context, user *domain.User) (int, error) {
	var id int
	var userPassword string
	row := r.db.QueryRowContext(ctx, `SELECT id, password FROM "user" AS u WHERE u.email = $1`, user.Email)
	if row.Err() != nil {
		return id, row.Err()
	}
	switch err := row.Scan(&id, &userPassword); err {
	case sql.ErrNoRows:
		return id, &InvalidUserError{}
	case nil:
		if r.hash.CheckPasswordHash(userPassword, user.Password) {
			return id, nil
		}
		return id, &InvalidUserError{}
	default:
		return -1, err
	}

}
func (r *Repository) CheckAccess(ctx context.Context, userId int) (bool, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id FROM "user" AS u WHERE u.id = $1`, userId)
	var id sql.NullInt64
	if scanErr := row.Scan(&id); scanErr != nil {
		return false, scanErr
	}
	if id.Valid {
		return true, nil
	}
	return false, &InvalidUserError{}
}
func (r *Repository) GetUserId(ctx context.Context, email string) (int, error) {
	row := r.db.QueryRowContext(ctx, `SELECT u.id FROM "user" AS u WHERE u.email = $1`, email)
	var id sql.NullInt64
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	if id.Valid {
		return int(id.Int64), nil
	}
	return -1, nil
}

func (r *Repository) SetNewPassword(ctx context.Context, userId int, newHashPassword string) error {
	_, err := r.db.ExecContext(ctx, `UPDATE "user" AS u SET password = $1 WHERE u.id = $2`, newHashPassword, userId)
	return err
}
func (r *Repository) CheckEmailExist(ctx context.Context, email string) (bool, error) {
	row := r.db.QueryRowContext(ctx, `SELECT u.email FROM "user" AS u WHERE u.email = $1`, email)
	var nullEmail sql.NullString
	if scanErr := row.Scan(&nullEmail); scanErr != nil {
		return false, scanErr
	}
	if nullEmail.Valid {
		return true, nil
	}
	return false, nil
}
