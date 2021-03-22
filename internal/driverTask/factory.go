package driverTask

import (
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/atulsinha007/uber/pkg/postgres"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	ApiHandler *Handler
)

func Init() {
	pgConf := postgres.PgConf{
		Host:     viper.GetString("PG_HOST"),
		Port:     viper.GetString("PG_PORT"),
		Username: viper.GetString("PG_USERNAME"),
		Password: viper.GetString("PG_PASSWORD"),
		DbName:   viper.GetString("PG_DB_NAME"),
	}

	driverTaskDao, err := NewDaoImpl(pgConf)
	if err != nil {
		log.L.With(zap.Error(err)).Fatal("error initializing driverTask dao")
	}
	log.L.Info("driverTask dao initialized")

	driverTaskCtrl := NewCtrl(driverTaskDao)
	ApiHandler = NewHandler(driverTaskCtrl)
}
