package database

import (
	"database/sql"
	"fmt"

	"github.com/homecloud/go-api/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("pgx", dsn)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
