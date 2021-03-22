package api

import (
	customerTaskApi "github.com/atulsinha007/uber/internal/customerTask/api"
	driverTaskApi "github.com/atulsinha007/uber/internal/driverTask/api"
	userApi "github.com/atulsinha007/uber/internal/user/api"
	"github.com/atulsinha007/uber/internal/user/customer"
	"github.com/atulsinha007/uber/internal/user/driver"
	"github.com/atulsinha007/uber/pkg/server"
)

func GetEndpoints() []server.Endpoint {
	var endPoints []server.Endpoint

	endPoints = append(endPoints, userApi.GetEndpoints()...)
	endPoints = append(endPoints, driver.GetEndpoints()...)
	endPoints = append(endPoints, customer.GetEndpoints()...)
	endPoints = append(endPoints, driverTaskApi.GetEndpoints()...)
	endPoints = append(endPoints, customerTaskApi.GetEndpoints()...)

	return endPoints
}
