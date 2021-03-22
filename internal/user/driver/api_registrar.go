package driver

import "github.com/atulsinha007/uber/pkg/server"

func GetEndpoints() []server.Endpoint {
	return []server.Endpoint{
		//{Path: "/driver", Method: "POST", Handler: CreateDriver},
	}
}
