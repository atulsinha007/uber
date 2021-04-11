package main

import (
	_ "github.com/atulsinha007/uber/config"
	"github.com/atulsinha007/uber/internal/api"
	"github.com/atulsinha007/uber/pkg/server"
)

func main() {
	server.Init()
	server.RegisterAndStart(api.GetEndpoints())
}
