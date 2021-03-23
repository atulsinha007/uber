package user

import (
	"github.com/atulsinha007/uber/internal/vehicle"
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

	userDao, err := NewDaoImpl(pgConf)
	if err != nil {
		log.L.With(zap.Error(err)).Fatal("error initializing user dao")
	}
	log.L.Info("user dao initialized")

	vehicleDao, err := vehicle.NewDaoImpl(pgConf)
	if err != nil {
		log.L.With(zap.Error(err)).Fatal("error initializing vehicle dao")
	}
	log.L.Info("vehicle dao initialized")

	userCtrl := NewCtrl(userDao, vehicleDao)
	ApiHandler = NewHandler(userCtrl)
}
