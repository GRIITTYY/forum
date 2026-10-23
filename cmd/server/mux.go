package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *Application) Routes() http.Handler {
	router := httprouter.New()

	needSession := alice.New(commonHeaders).Then(router)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.Healthcheck)
	router.HandlerFunc(http.MethodPost, "/v1/register", app.Register)

	return needSession
}
