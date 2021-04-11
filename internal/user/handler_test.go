package user

import (
	"bytes"
	"encoding/json"
	"errors"
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

	mockCtrl    *MockCtrl
	userHandler *Handler
}

func (suite *HandlerTestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func (suite *HandlerTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())

	suite.mockCtrl = NewMockCtrl(suite.ctrl)
	suite.userHandler = NewHandler(suite.mockCtrl)
}

// suite initialization
func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) Test_CreateUser() {
	payload := getCreateUserRequest()
	suite.mockCtrl.EXPECT().AddUser(CreateUserRequestToUser(payload)).Return(nil)

	b, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(b))

	response := suite.userHandler.CreateUser(request)
	suite.Equal(http.StatusCreated, response.Code)
	suite.Equal(response.Payload.(handler.Fields)["data"], "user creation successful")
}

func (suite *HandlerTestSuite) Test_GetDriverProfile() {
	payload := "1"
	driverProfile := getDriverProfileResponse()
	suite.mockCtrl.EXPECT().GetDriverProfile(1).Return(driverProfile, nil)

	b, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(b))
	request = mux.SetURLVars(request, map[string]string{"driverId": payload})

	response := suite.userHandler.GetDriverProfile(request)
	suite.Equal(http.StatusOK, response.Code)
	suite.Equal(response.Payload.(handler.Fields)["data"], driverProfile)
}

func (suite *HandlerTestSuite) Test_GetDriverProfileError() {
	payload := "1"
	driverProfile := getDriverProfileResponse()
	someError := errors.New("some error")
	suite.mockCtrl.EXPECT().GetDriverProfile(1).Return(driverProfile, someError)

	request, _ := http.NewRequest(http.MethodGet, "", bytes.NewReader([]byte{}))
	request = mux.SetURLVars(request, map[string]string{"driverId": payload})

	response := suite.userHandler.GetDriverProfile(request)
	suite.Equal(http.StatusInternalServerError, response.Code)
	suite.Equal(response.Payload.(handler.Fields)["error"], someError.Error())
}

func (suite *HandlerTestSuite) Test_UpdateLocation() {
	payload := getUpdateCurrentLocationRequest()
	suite.mockCtrl.EXPECT().UpdateLocation(payload).Return(nil)

	b, _ := json.Marshal(payload.CurLocation)
	request, _ := http.NewRequest(http.MethodPatch, "", bytes.NewReader(b))
	request = mux.SetURLVars(request, map[string]string{"userId": strconv.Itoa(payload.UserId)})

	response := suite.userHandler.UpdateLocation(request)
	suite.Equal(http.StatusCreated, response.Code)
	suite.Equal(response.Payload.(handler.Fields)["data"], "location updated successfully")
}

func (suite *HandlerTestSuite) Test_UpdateLocationError() {
	payload := getUpdateCurrentLocationRequest()
	someError := errors.New("some error")
	suite.mockCtrl.EXPECT().UpdateLocation(payload).Return(someError)

	b, _ := json.Marshal(payload.CurLocation)
	request, _ := http.NewRequest(http.MethodPatch, "", bytes.NewReader(b))
	request = mux.SetURLVars(request, map[string]string{"userId": strconv.Itoa(payload.UserId)})

	response := suite.userHandler.UpdateLocation(request)
	suite.Equal(http.StatusInternalServerError, response.Code)
	suite.Equal(response.Payload.(handler.Fields)["error"], someError.Error())
}

func (suite *HandlerTestSuite) Test_AddDriverWithVehicle() {
	payload := getDriverWithVehicleReq()
	suite.mockCtrl.EXPECT().AddDriverWithVehicle(payload).Return(nil)

	b, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(b))

	response := suite.userHandler.AddDriverWithVehicle(request)
	suite.Equal(http.StatusCreated, response.Code)
	suite.Equal(response.Payload.(handler.Fields)["data"], "driver with vehicle created successfully")
}

func (suite *HandlerTestSuite) Test_AddDriverWithVehicleError() {
	payload := getDriverWithVehicleReq()
	someError := errors.New("some error")
	suite.mockCtrl.EXPECT().AddDriverWithVehicle(payload).Return(someError)

	b, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(b))

	response := suite.userHandler.AddDriverWithVehicle(request)
	suite.Equal(http.StatusInternalServerError, response.Code)
	suite.Equal(response.Payload.(handler.Fields)["error"], someError.Error())
}

func (suite *HandlerTestSuite) Test_GetDriverHistory() {
	payload := "1"
	driverHistory := getDriverHistoryResponse()
	suite.mockCtrl.EXPECT().GetDriverHistory(1).Return(driverHistory, nil)

	request, _ := http.NewRequest(http.MethodGet, "", bytes.NewReader([]byte{}))
	request = mux.SetURLVars(request, map[string]string{"driverId": payload})

	response := suite.userHandler.GetDriverHistory(request)
	suite.Equal(http.StatusOK, response.Code)
	suite.Equal(response.Payload.(handler.Fields)["data"], driverHistory)
}

func (suite *HandlerTestSuite) Test_GetDriverHistoryError() {
	payload := "1"
	driverHistory := getDriverHistoryResponse()
	someError := errors.New("some error")
	suite.mockCtrl.EXPECT().GetDriverHistory(1).Return(driverHistory, someError)

	request, _ := http.NewRequest(http.MethodGet, "", bytes.NewReader([]byte{}))
	request = mux.SetURLVars(request, map[string]string{"driverId": payload})

	response := suite.userHandler.GetDriverHistory(request)
	suite.Equal(http.StatusInternalServerError, response.Code)
	suite.Equal(response.Payload.(handler.Fields)["error"], someError.Error())
}

func getCreateUserRequest() CreateUserRequest {
	return CreateUserRequest{
		FirstName: "a",
		LastName:  "b",
		Phone:     "123",
		Type:      "customer",
	}
}

func getDriverProfileResponse() DriverProfileResponse {
	return DriverProfileResponse{
		Name:          "driver1",
		PhoneNo:       "123",
		TotalRides:    1,
		AverageRating: "2.3",
	}
}

func getUpdateCurrentLocationRequest() UpdateCurrentLocationRequest {
	return UpdateCurrentLocationRequest{
		UserId: 1,
		CurLocation: LatLng{
			Lat: 12,
			Lng: 77,
		},
	}
}

func getDriverWithVehicleReq() DriverWithVehicleReq {
	return DriverWithVehicleReq{
		CreateUserRequest:  getCreateUserRequest(),
		Model:              "1",
		RegistrationNo:     "1",
		PermittedRideTypes: []string{"MICRO"},
	}
}

func getDriverHistoryResponse() []DriverHistoryResponse {
	return []DriverHistoryResponse{{
		Status:          "COMPLETED",
		DistanceCovered: 10,
		Rating:          2,
		PayoutAmount:    100,
	},
	}
}
