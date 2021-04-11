package customerTask

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HandlerTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
}

func (suite *HandlerTestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func (suite *HandlerTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
}

// suite initialization
func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) Test_CreateRideRequest() {

}

func (suite *HandlerTestSuite) Test_UpdateRideStops() {

}

func (suite *HandlerTestSuite) Test_CancelRide() {

}

func (suite *HandlerTestSuite) Test_GetCustomerHistory() {

}
