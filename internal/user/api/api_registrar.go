package api

import (
	"github.com/atulsinha007/uber/internal/user"
	"github.com/atulsinha007/uber/pkg/server"
)

func GetEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{Path: "/user", Method: "POST", Handler: user.ApiHandler.CreateUser},
		{Path: "/driver/{driverId}", Method: "GET", Handler: user.ApiHandler.GetDriverProfile},
		{Path: "/user/{userId}/location", Method: "PATCH", Handler: user.ApiHandler.UpdateLocation},
		{Path: "/driver", Method: "POST", Handler: user.ApiHandler.AddDriverWithVehicle},
	}
}
