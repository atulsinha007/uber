package customerTask

import (
	"bytes"
	"encoding/json"
	"github.com/atulsinha007/uber/internal/address"
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

	mockCtrl *MockCtrl
	handler  *Handler
}

func (suite *HandlerTestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func (suite *HandlerTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockCtrl = NewMockCtrl(suite.ctrl)
	suite.handler = NewHandler(suite.mockCtrl)
}

// suite initialization
func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) Test_CreateRideRequest() {
	payload := getCreateRideRequest()
	b, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(b))
	resp := getCreateRideResponseOnDriverAcceptance()

	suite.mockCtrl.EXPECT().CreateRide(payload).Return(resp, nil)
	response := suite.handler.CreateRideRequest(request)
	suite.Equal(http.StatusCreated, response.Code)
	suite.Equal(resp, response.Payload.(handler.Fields)["data"])
}

func (suite *HandlerTestSuite) Test_CancelRide() {
	customerTaskId := 1
	request, _ := http.NewRequest(http.MethodDelete, "", bytes.NewReader([]byte{}))
	request = mux.SetURLVars(request, map[string]string{"customerTaskId": strconv.Itoa(customerTaskId)})

	suite.mockCtrl.EXPECT().CancelRide(customerTaskId).Return(nil)
	response := suite.handler.CancelRide(request)
	suite.Equal(http.StatusOK, response.Code)
	suite.Equal("ride cancelled successfully", response.Payload.(handler.Fields)["data"])
}

func (suite *HandlerTestSuite) Test_GetCustomerHistory() {
	customerId := 1
	request, _ := http.NewRequest(http.MethodGet, "", bytes.NewReader([]byte{}))
	request = mux.SetURLVars(request, map[string]string{"customerId": strconv.Itoa(customerId)})
	history := getCustomerHistoryResponse()

	suite.mockCtrl.EXPECT().GetHistory(customerId).Return(history, nil)
	response := suite.handler.GetCustomerHistory(request)
	suite.Equal(http.StatusOK, response.Code)
	suite.Equal(history, response.Payload.(handler.Fields)["data"])
}

func getCreateRideRequest() CreateRideRequest {
	return CreateRideRequest{
		CustomerId:    1,
		PayableAmount: 100,
		PickupLocation: address.Location{
			Lat:        12,
			Lng:        77,
			Name:       "a",
			StreetName: "",
			Landmark:   "",
			City:       "",
			Country:    "",
		},
		DropLocation: address.Location{
			Lat:        12.1,
			Lng:        77.1,
			Name:       "b",
			StreetName: "",
			Landmark:   "",
			City:       "",
			Country:    "",
		},
		PreferredRideType: "MICRO",
	}
}

func getCreateRideResponseOnDriverAcceptance() CreateRideResponseOnDriverAcceptance {
	return CreateRideResponseOnDriverAcceptance{
		PickupLocation: address.Location{
			Lat:        12,
			Lng:        77,
			Name:       "a",
			StreetName: "",
			Landmark:   "",
			City:       "",
			Country:    "",
		},
		ETA: 120,
	}
}

func getCustomerHistoryResponse() []CustomerHistoryResponse {
	return []CustomerHistoryResponse{{
		RideId: 1,
		RideStops: []address.Location{{
			Lat:        12,
			Lng:        77,
			Name:       "a",
			StreetName: "",
			Landmark:   "",
			City:       "",
			Country:    "",
		}, {
			Lat:        12.1,
			Lng:        77.1,
			Name:       "b",
			StreetName: "",
			Landmark:   "",
			City:       "",
			Country:    "",
		}},
		Status:        "COMPLETED",
		PayableAmount: 100,
		PaymentStatus: "PENDING",
		DriverInfo: DriverInfo{
			Name:    "driver1",
			PhoneNo: "123",
		},
		RatingGiven:   "",
		DateOfJourney: "Monday, 02-Jan-06 15:04:05 MST",
	}}
}