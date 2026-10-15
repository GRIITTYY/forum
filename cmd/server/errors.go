package main

import "net/http"

func (app *Application) ClientError(w http.ResponseWriter, message string, status int) {
	data := jsonResponse{
		"message": message,
	}
	writeJson(w, data, status)
}
