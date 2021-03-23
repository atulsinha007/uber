package main

import (
	"github.com/atulsinha007/uber/config"
	"github.com/atulsinha007/uber/internal/api"
	"github.com/atulsinha007/uber/internal/customerTask"
	"github.com/atulsinha007/uber/internal/driverTask"
	"github.com/atulsinha007/uber/internal/user"
	"github.com/atulsinha007/uber/internal/vehicle"
	"github.com/atulsinha007/uber/pkg/server"
)

func main() {
	config.SetUpEnv()

	user.Init()
	vehicle.Init()
	customerTask.Init()
	driverTask.Init()

	server.Init()
	server.RegisterAndStart(api.GetEndpoints())
}
