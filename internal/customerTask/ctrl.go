package customerTask

import (
	"errors"
	"github.com/atulsinha007/uber/internal/address"
	"github.com/atulsinha007/uber/internal/driverTask"
	"github.com/atulsinha007/uber/pkg/log"
	"go.uber.org/zap"
	"time"
)

// CREATED -> ONGOING -> CANCELLED/COMPLETED
type Ctrl interface {
	CreateRide(createRideReq CreateRideRequest) (CreateRideResponseOnDriverAcceptance, error)
	UpdateRide(req UpdateRideReq) error
	CancelRide(customerTaskId string) error
	GetHistory(customerId string) ([]CustomerHistoryResponse, error)
	AssignNearestDriver(customerTaskId string, pickupLocation address.Location, preferredRideType string) (string, error)
}

type CtrlImpl struct {
	customerTaskDao Dao
	driverTaskDao   driverTask.Dao
}

func NewCtrl(dao Dao, driverTaskDao driverTask.Dao) *CtrlImpl {
	return &CtrlImpl{customerTaskDao: dao, driverTaskDao: driverTaskDao}
}

func (c *CtrlImpl) CreateRide(req CreateRideRequest) (resp CreateRideResponseOnDriverAcceptance, err error) {
	customerTaskId, err := c.customerTaskDao.CreateRide(req)
	if err != nil {
		return CreateRideResponseOnDriverAcceptance{}, err
	}

	// assign nearby driver

	var driverId string
	for i := 0; i < 5; i++ { // there should be psuedo-infinite tries for assignment
		driverId, err = c.AssignNearestDriver(customerTaskId, req.PickupLocation, req.PreferredRideType)
		if err != nil {
			log.L.With(zap.Error(err), zap.Any("customerTaskId", customerTaskId)).
				Error("error finding nearest driver")
		} else {
			break
		}
	}

	assigned := false
	for i := 0; i < 120; i++ { // this should be handled by some non-blocking approach
		time.Sleep(time.Second)
		assigned, resp = c.checkIfDriverAccepted(customerTaskId, driverId)
	}

	if !assigned {
		return CreateRideResponseOnDriverAcceptance{}, nil
	}

	return resp, nil
}

func (c *CtrlImpl) checkIfDriverAccepted(customerTaskId, driverId string) (assigned bool, resp CreateRideResponseOnDriverAcceptance) {
	task, err := c.driverTaskDao.GetFromDriverIdAndCustomerTaskId(customerTaskId, driverId)
	if err != nil {
		return false, CreateRideResponseOnDriverAcceptance{}
	}

	if task.Status != "ACCEPTED" {
		return false, CreateRideResponseOnDriverAcceptance{}
	}

	return true, CreateRideResponseOnDriverAcceptance{
		PickupLocation: address.Location{},
		ETA:            100, // some algorithm should decide ETA
	}
}

func (c *CtrlImpl) UpdateRide(req UpdateRideReq) error {
	return c.customerTaskDao.UpdateRide(req)
}

func (c *CtrlImpl) CancelRide(customerTaskId string) error {
	return c.customerTaskDao.CancelRide(customerTaskId)
}

func (c *CtrlImpl) GetHistory(driverId string) ([]CustomerHistoryResponse, error) {
	history, err := c.customerTaskDao.GetHistory(driverId)
	if err != nil {
		return nil, err
	}

	if len(history) == 0 {
		return nil, errors.New("record not found")
	}

	return history, nil
}

func (c *CtrlImpl) AssignNearestDriver(customerTaskId string, pickupLocation address.Location, preferredRideType string) (string, error) {
	// find nearest driver
	//driver, err := c.driverTaskDao.FindNearestDriver(pickupLocation, preferredRideType)
	//if err != nil {
	//	log.L.With(zap.Error(err), zap.Any("customerTaskId", customerTaskId)).
	//		Error("error finding nearest driver")
	//	return "", err
	//}

	// assign
	return "", nil
}
