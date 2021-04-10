package vehicle

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
}

func (suite *TestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func (suite *TestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
}

//suit initialization
func TestCtrlTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) Test_CreateVehicle() {

}
