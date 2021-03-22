package customerTask

type Ctrl interface {
	CreateRide(createRideReq CreateRideRequest) (CreateRideResponseOnDriverAcceptance, error)
	UpdateRide(req UpdateRideReq) error
	CancelRide(customerTaskId string) error
	GetHistory(customerId string) ([]CustomerHistoryResponse, error)
}

type CtrlImpl struct {
	dao Dao
}

func NewCtrl(dao Dao) *CtrlImpl {
	return &CtrlImpl{dao: dao}
}

func (c *CtrlImpl) CreateRide(req CreateRideRequest) (CreateRideResponseOnDriverAcceptance, error) {
	return CreateRideResponseOnDriverAcceptance{}, c.dao.CreateRide(req)
}

func (c *CtrlImpl) UpdateRide(req UpdateRideReq) error {
	return c.dao.UpdateRide(req)
}

func (c *CtrlImpl) CancelRide(customerTaskId string) error {
	return c.dao.CancelRide(customerTaskId)
}

func (c *CtrlImpl) GetHistory(driverId string) ([]CustomerHistoryResponse, error) {
	return c.dao.GetHistory(driverId)
}
