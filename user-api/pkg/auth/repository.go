package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"user-api/pkg/domain"
)

type IRepository interface {
	Register(ctx context.Context, user domain.User) error
	Login(ctx context.Context, user domain.User) (int, error)
	CheckAccess(ctx context.Context, userId int) (bool, error)
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

func (r *Repository) Register(ctx context.Context, user domain.User) error {
	hashPassword, hashErr := r.hash.Hash(user.Password)
	if hashErr != nil {
		return hashErr
	}
	_, err := r.db.ExecContext(ctx, `INSERT INTO "user"(username, password) VALUES ($1,$2)`, user.Username, hashPassword)
	if err != nil {
		return NewAlreadyRegisterError()
	}
	return nil
}
func (r *Repository) Login(ctx context.Context, user domain.User) (int, error) {
	var id int
	var userPassword string
	row := r.db.QueryRowContext(ctx, `SELECT id, password FROM "user" AS u WHERE u.username = $1`, user.Username)
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
		return 1, errors.New(fmt.Sprintf(`something went wrong: %s`, err.Error()))
	}
}
func (r *Repository) CheckAccess(ctx context.Context, userId int) (bool, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT * FROM "user" AS u WHERE u.id = $1`, userId)
	if err != nil {
		return false, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)
	var scanId int
	if scanErr := rows.Scan(&scanId); scanErr != nil {
		return false, scanErr
	}
	return true, nil
}
