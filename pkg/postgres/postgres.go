package postgres

import (
	"database/sql"
	"fmt"
	"github.com/atulsinha007/uber/pkg/log"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type PgConf struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

func GetDbConn(cfg PgConf) (*sql.DB, error) {
	logger := log.L.With(zap.String("host", cfg.Host), zap.String("port", cfg.Port),
		zap.String("dbName", cfg.DbName))

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DbName)
	db, err := sql.Open("postgres", connString)

	if err != nil {
		logger.With(zap.Error(err)).Error("DB connection failed")
		return nil, err
	}

	logger.Info("Connected to database")

	return db, nil
}
