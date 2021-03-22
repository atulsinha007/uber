package main

import (
	"github.com/atulsinha007/uber/config"
	"github.com/atulsinha007/uber/internal/api"
	"github.com/atulsinha007/uber/internal/factory"
	"github.com/atulsinha007/uber/pkg/server"
)

func main() {
	config.SetUpEnv()
	factory.Init()
	server.Init()
	server.RegisterAndStart(api.GetEndpoints())
}
