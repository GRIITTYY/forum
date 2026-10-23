package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"net/http"

	"github.com/alexedwards/argon2id"
)

func writeJson(w http.ResponseWriter, data any, status int) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	return json.MarshalWrite(w, data, jsontext.WithIndent("\t"))
}

func generatePWHash(password string) (string, error) {
	hashPassword, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}

	return hashPassword, nil
}
