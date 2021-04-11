package customerTask


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

	stmts := []string{
		`delete from customer_task`,
		`delete from users`,
		`alter sequence users_user_id_seq RESTART WITH 1`,
		`insert into users(first_name, last_name, phone, user_type) values($1, $2, $3, $4)`,
	}
	args := [][]interface{}{
		{},
		{},
		{},
		{"atul", "sinha", "123", "customer"},
	}

	for i := range stmts {
		_, err = suite.dao.db.Exec(stmts[i], args[i]...)
		if err != nil {
			log.L.With(zap.Error(err), zap.String("stmt", stmts[i])).Fatal("unable to delete")
		}
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

func (suite *DaoTestSuite) Test_CreateRide() {
	req := getCreateRideRequest()
	_, err := suite.dao.CreateRide(req)
	suite.Nil(err)
}

func (suite *DaoTestSuite) Test_CancelRide() {
	req := getCreateRideRequest()
	id, err := suite.dao.CreateRide(req)
	suite.Nil(err)

	err = suite.dao.CancelRide(id)
	suite.Nil(err)
}

func (suite *DaoTestSuite) Test_GetHistory() {
	_, err := suite.dao.GetHistory(1)
	suite.Nil(err)
}

func (suite *DaoTestSuite) Test_GetRideDetails() {
	req := getCreateRideRequest()
	id, err := suite.dao.CreateRide(req)
	suite.Nil(err)

	expDetails := CustomerTask{
		CustomerTaskId: id,
		CustomerId:     req.CustomerId,
		Status:         "CREATED",
		PayableAmount:  req.PayableAmount,
		RideType:       req.PreferredRideType,
	}

	details, err := suite.dao.GetRideDetails(id)
	suite.Nil(err)
	suite.Equal(expDetails, details)
}
