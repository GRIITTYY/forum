package main

import (
	"net/http"
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
