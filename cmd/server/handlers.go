package main

import (
	"encoding/json/v2"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/griittyy/forum/internal/models"
	"github.com/griittyy/forum/internal/validator"
)

func (app *Application) Healthcheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":      "available",
		"environment": app.Config.Environment,
		"version":     app.Config.Version,
	}
	writeJson(w, response, http.StatusOK)
}

func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	var regData models.UserRegistrationData

	if err := json.UnmarshalRead(r.Body, &regData, json.RejectUnknownMembers(true)); err != nil {
		app.LogError(r, "register request JSON could not be decoded", err)
		log.Printf("JSON Unmarshalling Error: %T\n", err)
		log.Printf("JSON Unmarshalling Error: %v\n", err)

		if errors.Is(err, io.ErrUnexpectedEOF) {
			app.ClientError(w, "request body must not be empty", http.StatusBadRequest)
			return
		}

		var semErr *json.SemanticError
		if errors.As(err, &semErr) {
			if semErr.Err == json.ErrUnknownName {
				msg := fmt.Sprintf("request body contains unsupported field: %v", semErr.JSONPointer)
				app.ClientError(w, msg, http.StatusBadRequest)
				return
			}

			app.ClientError(w, "request body contains badly formed JSON", http.StatusBadRequest)
			return
		}

		app.ClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := validator.ValidateRegistrationData(regData)
	if !result.Valid() {
		writeJson(w, result, http.StatusBadRequest)
		return
	}

	var response = map[string]string{
		"message": "user successfully registered",
	}
	writeJson(w, response, http.StatusCreated)
}
