package driverTask

import (
	"bytes"
	"encoding/json"
	handler "github.com/atulsinha007/uber/pkg/http-wrapper"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"net/http"
	"strconv"
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

func (suite *HandlerTestSuite) Test_AcceptRideRequest() {
	payload := getAcceptRideReq()
	b, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(b))
	request = mux.SetURLVars(request, map[string]string{"driverTaskId": strconv.Itoa(payload.DriverTaskId)})

	suite.mockCtrl.EXPECT().AcceptRideRequest(payload).Return(nil)
	response := suite.vehicleHandler.AcceptRideRequest(request)
	suite.Equal(http.StatusCreated, response.Code)
	suite.Equal("ride accepted successfully", response.Payload.(handler.Fields)["data"])
}

func (suite *HandlerTestSuite) Test_UpdateRide() {
	payload := getUpdateRideReq()
	b, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPatch, "", bytes.NewReader(b))
	request = mux.SetURLVars(request, map[string]string{"driverTaskId": strconv.Itoa(payload.DriverTaskId)})

	suite.mockCtrl.EXPECT().UpdateRide(payload).Return(nil)
	response := suite.vehicleHandler.UpdateRide(request)
	suite.Equal(http.StatusOK, response.Code)
	suite.Equal(response.Payload.(handler.Fields)["data"], "ride updated successfully")
}

func getAcceptRideReq() AcceptRideReq {
	return AcceptRideReq{
		DriverTaskId:   1,
		DriverId:       1,
		CustomerTaskId: 1,
	}
}

func getUpdateRideReq() UpdateRideReq {
	return UpdateRideReq{
		DriverTaskId:   1,
		Status:         "COMPLETED",
		CustomerTaskId: 1,
	}
}
