package vehicle

type Ctrl interface {
	CreateVehicle(req CreateVehicleRequest) error
}

type CtrlImpl struct {
	dao Dao
}

func NewCtrl(dao Dao) *CtrlImpl {
	return &CtrlImpl{dao: dao}
}

func (c *CtrlImpl) CreateVehicle(req CreateVehicleRequest) error {
	_, err := c.dao.CreateVehicle(req)
	return err
}
