package user

import (
	"github.com/atulsinha007/uber/config"
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/atulsinha007/uber/pkg/migrations"
	"github.com/atulsinha007/uber/pkg/postgres"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
)

type DaoTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	dao *DaoImpl
}

func (suite *DaoTestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func (suite *DaoTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())

	var err error

	conf := getTestPgConfig()
	port, _ := strconv.Atoi(conf.Port)

	setupDb(migrations.DbConfig{
		Host:     conf.Host,
		Port:     port,
		Username: conf.Username,
		Password: conf.Password,
		DbName:   conf.DbName,
	})

	suite.dao, err = NewDaoImpl(conf)
	suite.Nil(err)

	stmt := `delete from users; delete from driver_task;`
	_, err = suite.dao.db.Exec(stmt)
	if err != nil {
		log.L.With(zap.Error(err), zap.String("stmt", stmt)).Fatal("unable to delete")
	}
}

func setupDb(dbConfig migrations.DbConfig) {
	_, caller, _, _ := runtime.Caller(0)
	migrationsPath := filepath.Join(filepath.Dir(caller), "../../migrations")
	log.L.Info("migrationsPath:" + migrationsPath)
	err := migrations.Up(dbConfig, migrationsPath)
	if err != nil {
		logrus.WithField("dbName", dbConfig.DbName).WithError(err).Error("failed to setup db")
	}
}

func getTestPgConfig() postgres.PgConf {
	return postgres.PgConf{
		Host:     config.V.GetString("PG_HOST"),
		Port:     config.V.GetString("PG_PORT"),
		Username: config.V.GetString("PG_USERNAME"),
		Password: config.V.GetString("PG_PASSWORD"),
		DbName:   config.V.GetString("PG_DB_NAME"),
	}
}

// suite initialization
func TestDaoTestSuite(t *testing.T) {
	suite.Run(t, new(DaoTestSuite))
}

func (suite *DaoTestSuite) Test_Set() {
	user := CreateUserRequestToUser(getCreateUserRequest())
	_, err := suite.dao.Set(user)
	suite.Nil(err)
}

func (suite *DaoTestSuite) Test_GetDriverProfile() { // TODO: complete this by adding dummy data
	_, err := suite.dao.GetDriverProfile(1)
	suite.NotNil(err)
}

func (suite *DaoTestSuite) Test_UpdateLocation() {
	req := getUpdateCurrentLocationRequest()
	err := suite.dao.UpdateLocation(req)
	suite.Nil(err)
}

func (suite *DaoTestSuite) Test_AddDriverWithVehicle() {  // TODO: complete this by adding dummy data

}

func (suite *DaoTestSuite) Test_GetDriverHistory() { // TODO: complete this by adding dummy data

}
