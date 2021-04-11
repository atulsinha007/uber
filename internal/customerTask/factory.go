package customerTask

import (
	"github.com/atulsinha007/uber/config"
	"github.com/atulsinha007/uber/internal/driverTask"
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

	customerTaskDao, err := NewDaoImpl(pgConf)
	if err != nil {
		log.L.With(zap.Error(err)).Fatal("error initializing customerTask customerTaskDao")
	}
	log.L.Info("customerTask customerTaskDao initialized")

	driverTaskDao, err := driverTask.NewDaoImpl(pgConf)
	if err != nil {
		log.L.With(zap.Error(err)).Fatal("error initializing driverTask driverTaskDao")
	}
	log.L.Info("customerTask driverTaskDao initialized")

	customerTaskCtrl := NewCtrl(customerTaskDao, driverTaskDao)
	ApiHandler = NewHandler(customerTaskCtrl)
}
