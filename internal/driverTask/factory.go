package driverTask

import (
	"github.com/atulsinha007/uber/config"
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/atulsinha007/uber/pkg/postgres"
	"go.uber.org/zap"
)

var (
	ApiHandler *Handler
)

func init() {
	pgConf := postgres.PgConf{
		Host:     config.V.GetString("PG_HOST"),
		Port:     config.V.GetString("PG_PORT"),
		Username: config.V.GetString("PG_USERNAME"),
		Password: config.V.GetString("PG_PASSWORD"),
		DbName:   config.V.GetString("PG_DB_NAME"),
	}

	driverTaskDao, err := NewDaoImpl(pgConf)
	if err != nil {
		log.L.With(zap.Error(err)).Fatal("error initializing driverTask dao")
	}
	log.L.Info("driverTask dao initialized")

	driverTaskCtrl := NewCtrl(driverTaskDao)
	ApiHandler = NewHandler(driverTaskCtrl)
}
