package user

import "github.com/atulsinha007/uber/internal/vehicle"

type Ctrl interface {
	AddUser(user User) error
	AddDriverWithVehicle(driverWithVehicleReq DriverWithVehicleReq) error
	GetDriverProfile(driverId int) (DriverProfileResponse, error)
	UpdateLocation(request UpdateCurrentLocationRequest) error
	GetDriverHistory(driverId int) ([]DriverHistoryResponse, error)
}

type CtrlImpl struct {
	dao        Dao
	vehicleDao vehicle.Dao
}

func NewCtrl(dao Dao, vehicleDao vehicle.Dao) *CtrlImpl {
	return &CtrlImpl{dao: dao, vehicleDao: vehicleDao}
}

func (c *CtrlImpl) AddUser(user User) error {
	_, err := c.dao.Set(user)
	return err
}

func (c *CtrlImpl) AddDriverWithVehicle(driverWithVehicleReq DriverWithVehicleReq) error {
	id, err := c.vehicleDao.CreateVehicle(vehicle.CreateVehicleRequest{
		Model:              driverWithVehicleReq.Model,
		RegistrationNo:     driverWithVehicleReq.RegistrationNo,
		PermittedRideTypes: driverWithVehicleReq.PermittedRideTypes,
	})
	if err != nil {
		return err
	}
	return c.dao.AddDriverWithVehicle(id, CreateUserRequestToUser(driverWithVehicleReq.CreateUserRequest))
}

func (c *CtrlImpl) GetDriverProfile(driverId int) (DriverProfileResponse, error) {
	return c.dao.GetDriverProfile(driverId)
}

func (c *CtrlImpl) UpdateLocation(request UpdateCurrentLocationRequest) error {
	return c.dao.UpdateLocation(request)
}

func (c *CtrlImpl) GetDriverHistory(driverId int) ([]DriverHistoryResponse, error) {
	return c.dao.GetDriverHistory(driverId)
}