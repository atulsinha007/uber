package vehicle

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mockDao     *MockDao
	vehicleCtrl *CtrlImpl
}

func (suite *TestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func (suite *TestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockDao = NewMockDao(suite.ctrl)
	suite.vehicleCtrl = NewCtrl(suite.mockDao)
}

// suite initialization
func TestCtrlTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) Test_CreateVehicleSuccess() {
	v := getCreateVehicleRequest()
	suite.mockDao.EXPECT().CreateVehicle(v).Return("", nil)

	err := suite.vehicleCtrl.CreateVehicle(v)
	suite.Nil(err)
}

func (suite *TestSuite) Test_CreateVehicleFailure() {
	v := getCreateVehicleRequest()
	someError := errors.New("some error")
	suite.mockDao.EXPECT().CreateVehicle(v).Return("", someError)

	err := suite.vehicleCtrl.CreateVehicle(v)
	suite.Equal(someError, err)
}

func getCreateVehicleRequest() CreateVehicleRequest {
	return CreateVehicleRequest{
		Model:              "m1",
		RegistrationNo:     "r1",
		PermittedRideTypes: []string{"MICRO"},
	}
}
