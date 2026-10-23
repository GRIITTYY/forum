package main

import (
	"encoding/json/v2"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

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
	var response = make(map[string]string)

	if err := json.UnmarshalRead(r.Body, &regData, json.RejectUnknownMembers(true)); err != nil {
		app.LogError(r, "register request JSON could not be decoded", err)

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

			app.ClientError(w, "request body contains badly formed JSON or may be INCOMPLETE", http.StatusBadRequest)
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

	hashPassword, err := generatePWHash(regData.Password)
	if err != nil {
		app.LogError(r, "error with argon2id", err)
		app.ServerError(w, "An error occured while processing your request.", http.StatusInternalServerError)
		return
	}
	regData.Password = hashPassword

	userId, err := app.UserModel.RegisterNewUser(r.Context(), regData)
	if err != nil {
		if errors.Is(err, models.ErrEmailExists) || errors.Is(err, models.ErrUsernameExists) {
			response["message"] = err.Error()
			writeJson(w, response, http.StatusCreated)
			return
		}

		app.LogError(r, "error with inserting new user record", err)
		app.ServerError(w, "An error occured while processing your request.", http.StatusInternalServerError)
		return
	}

	response["message"] = "welcome to forum, you have registered successfully"
	response["user_id"] = strconv.Itoa(userId)
	writeJson(w, response, http.StatusCreated)
}
