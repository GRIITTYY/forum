package main

import (
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	DBPool *pgxpool.Pool
	Logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &Application{
		Logger: logger,
	}

	app.
}
