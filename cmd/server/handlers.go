package main

import (
	"encoding/json/v2"
	"net/http"

	"github.com/griittyy/forum/internal/models"
)

func (app *Application) Healthcheck(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.Config.Environment,
			"version":     app.Config.Version,
		},
	}
	writeJson(w, response, http.StatusOK)
}

func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info("register request received", "method", r.Method, "path", r.URL.Path, "content_type", r.Header.Get("Content-Type"))
	var data models.UserRegistrationData

	if err := json.UnmarshalRead(r.Body, &data, json.RejectUnknownMembers(true)); err != nil {
		app.Logger.Warn("register request JSON could not be decoded", "error", err)
		app.ClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var response = jsonResponse{
		"message": "user successfully registered",
	}
	writeJson(w, response, http.StatusCreated)
}
