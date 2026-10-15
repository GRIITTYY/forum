package main

import "net/http"

func (app *Application) LogError(r *http.Request, info string, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	app.Logger.Error(info, "error", err.Error(), "method", method, "uri", uri)
}

func (app *Application) ClientError(w http.ResponseWriter, message string, status int) {
	data := map[string]string{
		"message": message,
	}
	writeJson(w, data, status)
}

func (app *Application) ServerError(w http.ResponseWriter, message string, status int) {
	data := map[string]string{
		"message": message,
	}
	writeJson(w, data, status)
}
