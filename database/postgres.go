package database

import (
	"bwastartup/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func ConnectionPostgres(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.Psql.Host,
		cfg.Psql.User,
		cfg.Psql.Password,
		cfg.Psql.DBName,
		cfg.Psql.Port,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres] Failed to create DB Object")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres] Failed to ping database " + cfg.Psql.DBName)
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Psql.DBMaxOpen)
	db.SetMaxIdleConns(cfg.Psql.DBMaxIdle)
	db.SetConnMaxLifetime(time.Hour)

	log.Info().Msg("[ConnectionPostgres] Connected to PostgreSQL successfully")

	return db, nil
}
