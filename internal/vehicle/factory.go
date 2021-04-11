package vehicle

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

	vehicleDao, err := NewDaoImpl(pgConf)
	if err != nil {
		log.L.With(zap.Error(err)).Fatal("error initializing vehicle dao")
	}
	log.L.Info("vehicle dao initialized")

	userCtrl := NewCtrl(vehicleDao)
	ApiHandler = NewHandler(userCtrl)
}
