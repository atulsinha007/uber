package customerTask

import (
	"github.com/atulsinha007/uber/internal/driverTask"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mockDao          *MockDao
	driverTaskDao    *driverTask.MockDao
	customerTaskCtrl *CtrlImpl
}

func (suite *TestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func (suite *TestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockDao = NewMockDao(suite.ctrl)
	suite.driverTaskDao = driverTask.NewMockDao(suite.ctrl)
	suite.customerTaskCtrl = NewCtrl(suite.mockDao, suite.driverTaskDao)
}

// suite initialization
func TestCtrlTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) Test_CreateRide() {
	req := getCreateRideRequest()
	customerTaskId := 1
	driverId := 1
	distance := 100.0
	pickUpLoc := getCreateRideRequest().PickupLocation
	preferredRideType := "MICRO"
	task := driverTask.DriverTask{
		DriverTaskId:   1,
		CustomerTaskId: 1,
		DriverId:       1,
		Status:         "ACCEPTED",
		PayableAmount:  100,
		RideType:       "MICRO",
		Distance:       100,
	}
	expectedResp := CreateRideResponseOnDriverAcceptance{
		PickupLocation: pickUpLoc,
		ETA:            100,
	}

	suite.mockDao.EXPECT().CreateRide(req).Return(customerTaskId, nil)
	suite.driverTaskDao.EXPECT().FindNearestDriver(pickUpLoc, preferredRideType).Return(driverId, distance, nil)
	suite.driverTaskDao.EXPECT().CreateDriverTask(driverTask.DriverTask{
		CustomerTaskId: customerTaskId,
		DriverId:       driverId,
		Status:         "CREATED",
		PayableAmount:  distance,
		RideType:       preferredRideType,
		Distance:       distance,
	}).Return(nil)
	suite.driverTaskDao.EXPECT().GetFromDriverIdAndCustomerTaskId(customerTaskId, driverId).Return(task, nil)

	resp, err := suite.customerTaskCtrl.CreateRide(req)
	suite.Nil(err)
	suite.Equal(expectedResp, resp)
}

func (suite *TestSuite) Test_CancelRide() {
	customerTaskId := 1
	suite.mockDao.EXPECT().CancelRide(customerTaskId).Return(nil)
	err := suite.customerTaskCtrl.CancelRide(customerTaskId)
	suite.Nil(err)
}

func (suite *TestSuite) Test_GetHistory() {
	customerId := 1
	expectedHistory := getCustomerHistoryResponse()
	suite.mockDao.EXPECT().GetHistory(customerId).Return(expectedHistory, nil)
	history, err := suite.customerTaskCtrl.GetHistory(customerId)
	suite.Nil(err)
	suite.Equal(expectedHistory, history)
}

func (suite *TestSuite) Test_GetHistoryNilError() {
	customerId := 1
	var expectedHistory []CustomerHistoryResponse
	suite.mockDao.EXPECT().GetHistory(customerId).Return(expectedHistory, nil)
	_, err := suite.customerTaskCtrl.GetHistory(customerId)
	suite.NotNil(err)
	suite.Equal("record not found", err.Error())
}

func (suite *TestSuite) Test_AssignNearestDriver() {
	driverId := 1
	distance := 10.0
	customerTaskId := 1
	pickUpLoc := getCreateRideRequest().PickupLocation
	preferredRideType := "MICRO"

	suite.driverTaskDao.EXPECT().FindNearestDriver(pickUpLoc, preferredRideType).Return(driverId, distance, nil)
	suite.driverTaskDao.EXPECT().CreateDriverTask(driverTask.DriverTask{
		CustomerTaskId: customerTaskId,
		DriverId:       driverId,
		Status:         "CREATED",
		PayableAmount:  distance,
		RideType:       preferredRideType,
		Distance:       distance,
	}).Return(nil)

	id, err := suite.customerTaskCtrl.AssignNearestDriver(customerTaskId, pickUpLoc, preferredRideType)
	suite.Nil(err)
	suite.Equal(driverId, id)
}
