package user

import (
	"github.com/atulsinha007/uber/internal/vehicle"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mockUserDao    *MockDao
	mockVehicleDao *vehicle.MockDao
	userCtrl       *CtrlImpl
}

func (suite *TestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func (suite *TestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockUserDao = NewMockDao(suite.ctrl)
	suite.mockVehicleDao = vehicle.NewMockDao(suite.ctrl)
	suite.userCtrl = NewCtrl(suite.mockUserDao, suite.mockVehicleDao)
}

// suite initialization
func TestCtrlTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) Test_AddUser() {
	user := CreateUserRequestToUser(getCreateUserRequest())
	suite.mockUserDao.EXPECT().Set(user).Return("", nil)
	err := suite.userCtrl.AddUser(user)
	suite.Nil(err)
}

func (suite *TestSuite) Test_AddDriverWithVehicle() {
	req := getDriverWithVehicleReq()
	id := "1"
	suite.mockVehicleDao.EXPECT().CreateVehicle(vehicle.CreateVehicleRequest{
		Model:              req.Model,
		RegistrationNo:     req.RegistrationNo,
		PermittedRideTypes: req.PermittedRideTypes,
	}).Return(id, nil)
	suite.mockUserDao.EXPECT().AddDriverWithVehicle(id, CreateUserRequestToUser(req.CreateUserRequest)).Return(nil)

	err := suite.userCtrl.AddDriverWithVehicle(req)
	suite.Nil(err)
}

func (suite *TestSuite) Test_GetDriverProfile() {
	resp := getDriverProfileResponse()
	suite.mockUserDao.EXPECT().GetDriverProfile(1).Return(resp, nil)
	res, err := suite.userCtrl.GetDriverProfile(1)
	suite.Nil(err)
	suite.Equal(resp, res)
}

func (suite *TestSuite) Test_UpdateLocation() {
	req := getUpdateCurrentLocationRequest()
	suite.mockUserDao.EXPECT().UpdateLocation(req).Return(nil)
	err := suite.userCtrl.UpdateLocation(req)
	suite.Nil(err)
}

func (suite *TestSuite) Test_GetDriverHistory() {
	driverId := 1
	resp := getDriverHistoryResponse()
	suite.mockUserDao.EXPECT().GetDriverHistory(driverId).Return(resp, nil)
	history, err := suite.userCtrl.GetDriverHistory(driverId)
	suite.Nil(err)
	suite.Equal(resp, history)
}
