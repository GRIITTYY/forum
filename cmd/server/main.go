package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/griittyy/forum/internals/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	DBPool *pgxpool.Pool
	Logger *slog.Logger
	Config Config
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
		logger.Error("Erro creating Database Pool", "msg", err.Error())
		os.Exit(1)
	}

	app := &Application{
		Logger: logger,
		DBPool: dbPool,
		Config: config,
	}

	srv := http.Server{
		Addr:    config.Addr,
		Handler: app.Routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		app.Logger.Error("Failed to start server", "msg", err.Error())
		os.Exit(1)
	}
}
