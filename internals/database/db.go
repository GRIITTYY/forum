package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(connStr string) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = 25                       // Maximum active connections
	poolConfig.MinConns = 5                        // Minimum connections to keep warm
	poolConfig.MaxConnLifetime = 30 * time.Minute  // Recycle connections every 30 mins
	poolConfig.MaxConnIdleTime = 10 * time.Minute  // Drop connections idle for > 10 mins
	poolConfig.HealthCheckPeriod = 1 * time.Minute // Ping database every minute

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	log.Println("SUCCESSFULLY CONNECTED TO POSTGRESQL... 🎉🎉🎉🎉🎉")

	return pool, nil
}
