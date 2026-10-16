package models

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

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

func (u *UserModel) RegisterNewUser(r *http.Request, data UserRegistrationData) error {
	stmt := `INSERT INTO auth(email, username, name, password) VALUES($1, $2, $3, $4)`
}
