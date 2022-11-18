package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"user-api/internal/domain"
)

type IRepository interface {
	Register(ctx context.Context, user *domain.RegisterUser) error
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

func (r *Repository) Register(ctx context.Context, user *domain.RegisterUser) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO "user" (email, username, password) VALUES ($1,$2,$3)`, user.Email, user.Username, user.Password)
	if err != nil {
		return NewAlreadyRegisterError()
	}
	return nil
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
		return -1, errors.New(fmt.Sprintf(`something went wrong: %s`, err.Error()))
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
	if scanErr := row.Scan(&id); scanErr != nil {
		return -1, scanErr
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
