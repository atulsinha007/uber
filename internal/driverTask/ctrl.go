package driverTask

//go:generate mockgen -destination=mock_ctrl.go -package=driverTask -source=./ctrl.go
type Ctrl interface {
	AcceptRideRequest(req AcceptRideReq) error
	UpdateRide(req UpdateRideReq) error
}

type CtrlImpl struct {
	dao Dao
}

func NewCtrl(dao Dao) *CtrlImpl {
	return &CtrlImpl{dao: dao}
}

func (c *CtrlImpl) AcceptRideRequest(req AcceptRideReq) error {
	return c.dao.AcceptRideRequest(req)
}

func (c *CtrlImpl) UpdateRide(req UpdateRideReq) error {
	return c.dao.UpdateRide(req)
}
