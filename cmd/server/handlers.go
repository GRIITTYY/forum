package main

import (
	"encoding/json/v2"
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
