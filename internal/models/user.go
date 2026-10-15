package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type jsonResponse map[string]any

type UserRegistrationData struct {
	Username        string `json:"username"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm-password"`
}

type UserModel struct {
	DBPool *pgxpool.Pool
}

// func (u *UserModel) NewUser(data UserRegistrationData) (jsonResponse, error) {

// }
