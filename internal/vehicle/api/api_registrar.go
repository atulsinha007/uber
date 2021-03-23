package api

import (
	"github.com/atulsinha007/uber/internal/vehicle"
	"github.com/atulsinha007/uber/pkg/server"
)

func GetEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{Path: "/vehicle", Method: "POST", Handler: vehicle.ApiHandler.CreateVehicle},
	}
}
