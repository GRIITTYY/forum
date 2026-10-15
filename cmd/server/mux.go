package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.Healthcheck)
	router.HandlerFunc(http.MethodPost, "/v1/register", app.Register)

	return router
}
