package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/griittyy/forum/internal/database"
	"github.com/griittyy/forum/internal/models"
)

type Application struct {
	Logger         *slog.Logger
	Config         Config
	UserModel      *models.UserModel
	SessionManager *scs.SessionManager
}

type Config struct {
	Version     string
	Addr        string
	Environment string
}

func main() {
	var config Config

	flag.StringVar(&config.Addr, "addr", ":5500", "Specify Host:Port for serving this program")
	flag.StringVar(&config.Environment, "env", "development", "Specify whether Development or Production")
	flag.StringVar(&config.Version, "version", "1", "Specify current version of this program")
	connString := flag.String("dsn", "postgres://forum_server:4op3wScWuEGff2ab@localhost:5432/forum_db?timezone=Africa/Lagos", "Database URL/Connection string for PostgreSQL Connection.")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	dbPool, err := database.NewDB(*connString)
	if err != nil {
		logger.Error("Error creating Database Pool", "error", err.Error())
		os.Exit(1)
	}

	sessionManager := scs.New()
	sessionManager.Store = pgxstore.New(dbPool)
	sessionManager.Lifetime = 48 * time.Hour
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Path = "/"
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode

	app := &Application{
		Logger:         logger,
		Config:         config,
		UserModel:      &models.UserModel{DBPool: dbPool},
		SessionManager: sessionManager,
	}

	srv := http.Server{
		Addr:    config.Addr,
		Handler: app.Routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		app.Logger.Error("Failed to start server", "error", err.Error())
		os.Exit(1)
	}
}
