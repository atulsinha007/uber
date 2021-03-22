package api

import (
	"github.com/atulsinha007/uber/internal/driverTask"
	"github.com/atulsinha007/uber/pkg/server"
)

func GetEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{Path: "/driverTask/{driverId}", Method: "POST", Handler: driverTask.ApiHandler.GetDriverHistory},
	}
}
