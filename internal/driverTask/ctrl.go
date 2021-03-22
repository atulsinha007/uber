package driverTask

type Ctrl interface {
	GetDriverHistory(driverId string) ([]DriverHistoryResponse, error)
	AcceptRideRequest(req AcceptRideReq) error
	UpdateRide(req UpdateRideReq) error
}

type CtrlImpl struct {
	dao Dao
}

func NewCtrl(dao Dao) *CtrlImpl {
	return &CtrlImpl{dao: dao}
}

func (c *CtrlImpl) GetDriverHistory(driverId string) ([]DriverHistoryResponse, error) {
	return c.dao.GetDriverHistory(driverId)
}

func (c *CtrlImpl) AcceptRideRequest(req AcceptRideReq) error {
	return c.dao.AcceptRideRequest(req)
}

func (c *CtrlImpl) UpdateRide(req UpdateRideReq) error {
	return c.dao.UpdateRide(req)
}
