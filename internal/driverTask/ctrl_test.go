package driverTask

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mockDao        *MockDao
	driverTaskCtrl *CtrlImpl
}

func (suite *TestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func (suite *TestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockDao = NewMockDao(suite.ctrl)
	suite.driverTaskCtrl = NewCtrl(suite.mockDao)
}

// suite initialization
func TestCtrlTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) Test_AcceptRideRequest() {
	req := getAcceptRideReq()
	suite.mockDao.EXPECT().AcceptRideRequest(req).Return(nil)
	err := suite.driverTaskCtrl.AcceptRideRequest(req)
	suite.Nil(err)
}

func (suite *TestSuite) Test_AcceptRideRequestError() {
	req := getAcceptRideReq()
	someError := errors.New("some error")
	suite.mockDao.EXPECT().AcceptRideRequest(req).Return(someError)
	err := suite.driverTaskCtrl.AcceptRideRequest(req)
	suite.Equal(someError, err)
}


func (suite *TestSuite) Test_UpdateRide() {
	req := getUpdateRideReq()
	suite.mockDao.EXPECT().UpdateRide(req).Return(nil)
	err := suite.driverTaskCtrl.UpdateRide(req)
	suite.Nil(err)
}

func (suite *TestSuite) Test_UpdateRideError() {
	req := getUpdateRideReq()
	someError := errors.New("some error")
	suite.mockDao.EXPECT().UpdateRide(req).Return(someError)
	err := suite.driverTaskCtrl.UpdateRide(req)
	suite.Equal(someError, err)
}