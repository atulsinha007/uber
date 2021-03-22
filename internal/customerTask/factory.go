package customerTask

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

	customerTaskDao, err := NewDaoImpl(pgConf)
	if err != nil {
		log.L.With(zap.Error(err)).Fatal("error initializing customerTask dao")
	}
	log.L.Info("customerTask dao initialized")

	customerTaskCtrl := NewCtrl(customerTaskDao)
	ApiHandler = NewHandler(customerTaskCtrl)
}
