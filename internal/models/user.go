package models

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrEmailExists error = errors.New("An account with this email address already exists. Please try a different email address.")
var ErrUsernameExists error = errors.New("This username is already taken. Please try a different username")

type UserRegistrationData struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm-password"`
}

type UserModel struct {
	DBPool *pgxpool.Pool
}

func (u *UserModel) RegisterNewUser(ctx context.Context, data UserRegistrationData) (int, error) {
	var userId int

	exists, err := u.EmailExists(ctx, data.Email)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, ErrEmailExists
	}

	exists, err = u.UsernameExists(ctx, data.Username)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, ErrUsernameExists
	}

	stmt := `INSERT INTO auth(email, username, name, password) VALUES($1, $2, $3, $4) RETURNING user_id`

	err = u.DBPool.QueryRow(ctx, stmt, data.Email, data.Username, data.Name, data.Password).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (u *UserModel) EmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool

	stmt := `SELECT EXISTS(SELECT 1 FROM auth where email=$1)`
	if err := u.DBPool.QueryRow(ctx, stmt, email).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (u *UserModel) UsernameExists(ctx context.Context, username string) (bool, error) {
	var exists bool

	stmt := `SELECT EXISTS(SELECT 1 FROM auth username email=$1)`
	if err := u.DBPool.QueryRow(ctx, stmt, username).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}
