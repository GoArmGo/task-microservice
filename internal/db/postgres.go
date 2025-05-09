package db

import (
	"fmt"
	"log"
	"task-microservice/pkg/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgres(cfg *config.Config) *Postgres {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sqlx.Connect(cfg.DBDriver, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &Postgres{DB: db}
}

// пинг для чека
func (p *Postgres) Ping() error {
	return p.DB.Ping()
}
