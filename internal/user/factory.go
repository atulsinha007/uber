package user

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

	//addrDao, err := address.NewDaoImpl(pgConf, AddressDbName)
	//if err != nil {
	//	log.L.With(zap.Error(err)).Fatal("error initializing address dao")
	//}
	//log.L.Info("address dao initialized")

	userDao, err := NewDaoImpl(pgConf)
	if err != nil {
		log.L.With(zap.Error(err)).Fatal("error initializing user dao")
	}
	log.L.Info("user dao initialized")

	userCtrl := NewCtrl(userDao)
	ApiHandler = NewHandler(userCtrl)
}
