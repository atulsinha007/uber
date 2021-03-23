package api

import (
	"github.com/atulsinha007/uber/internal/driverTask"
	"github.com/atulsinha007/uber/pkg/server"
)

func GetEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{Path: "/driverTask/{driverId}/history", Method: "POST", Handler: driverTask.ApiHandler.GetDriverHistory},
		{Path: "/driverTask/{driverTaskId}/accept", Method: "POST", Handler: driverTask.ApiHandler.AcceptRideRequest},
		{Path: "/driverTask/{driverTaskId}/update", Method: "PATCH", Handler: driverTask.ApiHandler.UpdateRide},
	}
}
