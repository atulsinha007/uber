package api

import (
	"github.com/atulsinha007/uber/internal/customerTask"
	"github.com/atulsinha007/uber/pkg/server"
)

func GetEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{Path: "/customerTask", Method: "POST", Handler: customerTask.ApiHandler.CreateRideRequest},
		{Path: "/customerTask/{customerTaskId}", Method: "PATCH", Handler: customerTask.ApiHandler.UpdateRideStops},
		{Path: "/customerTask/{customerTaskId}", Method: "DELETE", Handler: customerTask.ApiHandler.CancelRide},
		{Path: "/customerTask/{customerId}", Method: "GET", Handler: customerTask.ApiHandler.GetCustomerHistory},
	}
}
