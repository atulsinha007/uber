package user

type Ctrl interface {
	AddUser(user User) error
	GetDriverProfile(driverId string) (DriverProfileResponse, error)
	UpdateLocation(request UpdateCurrentLocationRequest) error
}

type CtrlImpl struct {
	dao Dao
}

func NewCtrl(dao Dao) *CtrlImpl {
	return &CtrlImpl{dao: dao}
}

func (c *CtrlImpl) AddUser(user User) error {
	return c.dao.Set(user)
}

func (c *CtrlImpl) GetDriverProfile(driverId string) (DriverProfileResponse, error) {
	return c.dao.GetDriverProfile(driverId)
}

func (c *CtrlImpl) UpdateLocation(request UpdateCurrentLocationRequest) error {
	return c.dao.UpdateLocation(request)
}
