package vehicle

import (
	"bytes"
	"encoding/json"
	"errors"
	handler "github.com/atulsinha007/uber/pkg/http-wrapper"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type HandlerTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mockCtrl       *MockCtrl
	vehicleHandler *Handler
}

func (suite *HandlerTestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func (suite *HandlerTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockCtrl = NewMockCtrl(suite.ctrl)
	suite.vehicleHandler = NewHandler(suite.mockCtrl)
}

// suite initialization
func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) Test_CreateVehicle() {
	payload := getCreateVehicleRequest()
	b, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(b))

	suite.mockCtrl.EXPECT().CreateVehicle(payload).Return(nil)
	response := suite.vehicleHandler.CreateVehicle(request)
	suite.Equal(http.StatusCreated, response.Code)
}

func (suite *HandlerTestSuite) Test_CreateVehicleError() {
	payload := getCreateVehicleRequest()
	b, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(b))

	someError := errors.New("some error")

	suite.mockCtrl.EXPECT().CreateVehicle(payload).Return(someError)
	response := suite.vehicleHandler.CreateVehicle(request)
	suite.Equal(http.StatusInternalServerError, response.Code)
	suite.Equal(response.Payload.(handler.Fields)["error"], someError.Error())
}

func (suite *HandlerTestSuite) Test_CreateVehicleInvalidPayloadError() {
	payload := getCreateVehicleRequest()
	payload.PermittedRideTypes[0] = "random"
	b, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(b))

	response := suite.vehicleHandler.CreateVehicle(request)
	suite.Equal(http.StatusBadRequest, response.Code)
}
