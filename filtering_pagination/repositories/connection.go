package repositories

import (
	"context"
	"fmt"
	"time"
    "log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func Connect(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname,
	)

	db, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Println("(connect.go)failed to create connection pool:", err)
		return nil, fmt.Errorf("failed to create connection pool: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.Ping(ctx); err != nil {
		log.Println("(connect.go)failed to connect to PostgreSQL:", err)	
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
	return db, nil
}
