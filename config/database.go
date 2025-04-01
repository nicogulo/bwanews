package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgress() (*Postgres, error) {
	dbConnString  := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	cfg.PsqlDB.User,
	cfg.PsqlDB.Password,
	cfg.PsqlDB.Host,
	cfg.PsqlDB.Port,
	cfg.PsqlDB.DbName,
)

	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgress-1] Failed connecting to database" + cfg.PsqlDB.Host)
		return nil, err
	}

	sqlDb, err := db.DB()

	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgress-2] Failed connecting to database" + cfg.PsqlDB.Host)
		return nil, err
	}

	sqlDb.SetMaxOpenConns(cfg.PsqlDB.DbMaxOpen)
	sqlDb.SetMaxIdleConns(cfg.PsqlDB.DbMaxIdle)

	return &Postgres{DB: db}, nil
}