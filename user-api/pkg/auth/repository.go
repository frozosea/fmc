package auth

import (
	"context"
	"database/sql"
	"errors"
	"user-api/pkg/user"
)

type IRepository interface {
	Register(ctx context.Context, user user.User) error
	Login(ctx context.Context, user user.User) (int, error)
	CheckAccess(ctx context.Context, userId int) (bool, error)
}

type Repository struct {
	db   *sql.DB
	hash IHash
}

func (r *Repository) Register(ctx context.Context, user user.User) error {
	hashPassword, hashErr := r.hash.Hash(user.Password)
	if hashErr != nil {
		return hashErr
	}
	_, err := r.db.ExecContext(ctx, `INSERT INTO "user"(username, password) VALUES ($1,$2)`, user.Username, hashPassword)
	return err
}
func (r *Repository) Login(ctx context.Context, user user.User) (int, error) {
	var id int
	var userPassword string
	row, queryErr := r.db.QueryContext(ctx, `SELECT id, password FROM "user" AS u WHERE u.username = $1`, user.Username)
	if queryErr != nil {
		return id, queryErr
	}
	switch err := row.Scan(&id, &userPassword); err {
	case sql.ErrNoRows:
		return id, sql.ErrNoRows
	case nil:
		if r.hash.CheckPasswordHash(userPassword, user.Password) {
			return id, nil
		}
		return id, errors.New(`password is invalid`)
	default:
		return 1, errors.New(`something went wrong`)
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
