package factory

import (
	"github.com/atulsinha007/uber/internal/customerTask"
	"github.com/atulsinha007/uber/internal/driverTask"
	"github.com/atulsinha007/uber/internal/user"
)

//var (
//	UserHandler       *api2.Handler
//	DriverTaskHandler *api.Handler
//)

func Init() {
	//pgConf := postgres.PgConf{
	//	Host:     viper.GetString("PG_HOST"),
	//	Port:     viper.GetString("PG_PORT"),
	//	Username: viper.GetString("PG_USERNAME"),
	//	Password: viper.GetString("PG_PASSWORD"),
	//}

	//addrDao, err := address.NewDaoImpl(pgConf, AddressDbName)
	//if err != nil {
	//	log.L.With(zap.Error(err)).Fatal("error initializing address dao")
	//}
	//log.L.Info("address dao initialized")
	//
	//userDao, err := user.NewDaoImpl(pgConf, DriverDbName)
	//if err != nil {
	//	log.L.With(zap.Error(err)).Fatal("error initializing user dao")
	//}
	//log.L.Info("user dao initialized")
	//
	//userCtrl := user.NewCtrl(userDao)
	//UserHandler = api2.NewHandler(userCtrl)

	user.Init()

	//driverTaskDao, err := driverTask.NewDaoImpl(pgConf, DriverTaskDbName)
	//if err != nil {
	//	log.L.With(zap.Error(err)).Fatal("error initializing driverTask dao")
	//}
	//log.L.Info("driverTask dao initialized")
	//
	//driverTaskCtrl := driverTask.NewCtrl(driverTaskDao)
	//DriverTaskHandler = api.NewHandler(driverTaskCtrl)
	customerTask.Init()
	driverTask.Init()

	//rideStopsDao, err := ride.NewDaoImpl(pgConf, RideStopsDbName)
	//if err != nil {
	//	log.L.With(zap.Error(err)).Fatal("error initializing rideStops dao")
	//}
	//log.L.Info("rideStops dao initialized")

	//documentsDao, err := documents.NewDaoImpl(pgConf, DocumentsDbName)
	//if err != nil {
	//	log.L.With(zap.Error(err)).Fatal("error initializing documents dao")
	//}
	//log.L.Info("documents dao initialized")

}
