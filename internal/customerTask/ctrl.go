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
	CancelRide(customerTaskId int) error
	GetHistory(customerId int) ([]CustomerHistoryResponse, error)
	AssignNearestDriver(customerTaskId int, pickupLocation address.Location, preferredRideType string) (int, error)
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

	assigned := false
	var driverId int
	for i := 0; i < 5; i++ { // there should be psuedo-infinite tries for assignment
		driverId, err = c.AssignNearestDriver(customerTaskId, req.PickupLocation, req.PreferredRideType)
		if err != nil {
			log.L.With(zap.Error(err), zap.Any("customerTaskId", customerTaskId)).
				Error("error finding nearest driver")
			continue
		}
		assigned = true
		break
	}

	if !assigned {
		return CreateRideResponseOnDriverAcceptance{}, nil
	}

	accepted := false
	for i := 0; i < 120; i++ { // this should be handled by some non-blocking approach
		time.Sleep(time.Second)
		accepted, resp = c.checkIfDriverAccepted(customerTaskId, driverId, req.PickupLocation)
		if accepted {
			break
		}
	}

	if !accepted {
		return CreateRideResponseOnDriverAcceptance{}, nil
	}

	return resp, nil
}

func (c *CtrlImpl) checkIfDriverAccepted(customerTaskId, driverId int, pickupLoc address.Location) (
	assigned bool, resp CreateRideResponseOnDriverAcceptance) {

	task, err := c.driverTaskDao.GetFromDriverIdAndCustomerTaskId(customerTaskId, driverId)
	if err != nil {
		return false, CreateRideResponseOnDriverAcceptance{}
	}

	if task.Status != "ACCEPTED" {
		return false, CreateRideResponseOnDriverAcceptance{}
	}

	return true, CreateRideResponseOnDriverAcceptance{
		PickupLocation: pickupLoc,
		ETA:            100, // some algorithm should decide ETA
	}
}

func (c *CtrlImpl) UpdateRide(req UpdateRideReq) error {
	return c.customerTaskDao.UpdateRide(req)
}

func (c *CtrlImpl) CancelRide(customerTaskId int) error {
	return c.customerTaskDao.CancelRide(customerTaskId)
}

func (c *CtrlImpl) GetHistory(customerId int) ([]CustomerHistoryResponse, error) {
	history, err := c.customerTaskDao.GetHistory(customerId)
	if err != nil {
		return nil, err
	}

	if len(history) == 0 {
		return nil, errors.New("record not found")
	}

	return history, nil
}

func (c *CtrlImpl) AssignNearestDriver(customerTaskId int, pickupLocation address.Location, preferredRideType string) (int, error) {
	driverId, distance, err := c.driverTaskDao.FindNearestDriver(pickupLocation, preferredRideType)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("customerTaskId", customerTaskId)).
			Error("error finding nearest driver")
		return 0, err
	}

	err = c.driverTaskDao.CreateDriverTask(driverTask.DriverTask{
		CustomerTaskId: customerTaskId,
		DriverId:       driverId,
		Status:         "CREATED",
		PayableAmount:  distance, // let's take amount=distance for now
		RideType:       preferredRideType,
		Distance:       distance,
	})
	if err != nil {
		return 0, err
	}

	return driverId, nil
}
