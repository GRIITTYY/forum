package models

import (
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

// func (u *UserModel) NewUser(data UserRegistrationData) (jsonResponse, error) {

// }

