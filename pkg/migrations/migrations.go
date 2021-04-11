package migrations

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var ErrNoChange = errors.New("no change")

type DbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func Up(cfg DbConfig, migrationsPath string) error {
	dbConn, err := connectToDb(cfg)
	if err != nil {
		log.L.Info("error1:" + err.Error())
		return err
	}

	m, err := newMigrate(dbConn, migrationsPath)
	if err != nil {
		log.L.Info("error2:" + err.Error())
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		// We want to identify this error at the caller. Typically when we call migrate.Up,
		// we don't care if the migrations are already applied, so we should return an identifiable error.
		return ErrNoChange
	}
	log.L.With(zap.Error(err)).Info("error3:")

	return err
}

func connectToDb(cfg DbConfig) (*sql.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DbName,
	)

	log.L.Info("connString:" + connString)

	return sql.Open("postgres", connString)
}

func newMigrate(dbConn *sql.DB, path string) (*migrate.Migrate, error) {
	log.L.Info("path11:" + path)
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		log.L.Info("err:" + err.Error())
		return nil, err
	}
	log.L.Info("path:" + path)

	return migrate.NewWithDatabaseInstance("file:///"+path, "postgres", driver)
}
