package customerTask

import (
	"database/sql"
	"github.com/atulsinha007/uber/internal/address"
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/atulsinha007/uber/pkg/postgres"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type Dao interface {
	CreateRide(createRideReq CreateRideRequest) error
	UpdateRide(req UpdateRideReq) error
	CancelRide(customerTaskId string) error
	GetHistory(customerId string) ([]CustomerHistoryResponse, error)
}

type DaoImpl struct {
	db *sql.DB
}

func NewDaoImpl(conf postgres.PgConf) (*DaoImpl, error) {
	conn, err := postgres.GetDbConn(conf)
	if err != nil {
		return nil, err
	}

	return &DaoImpl{db: conn}, nil
}

func (d *DaoImpl) CreateRide(req CreateRideRequest) error {
	tx, err := d.db.Begin()
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("createRideReq", req)).
			Error("error creating customer task")
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var customerTaskId string
	query := `insert into customer_task(customer_id, payable_amount, ride_type, status) values($1, $2, $3, $4) returning id;`
	err = tx.QueryRow(query, req.CustomerId, req.PayableAmount, req.PreferredRideType, "CREATED").Scan(&customerTaskId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("createRideReq", req)).
			Error("error creating customer task")
		tx.Rollback()
		return err
	}

	var pickupLocId, dropLocId string
	query = `insert into address(lat, lng, house_name, street_name, landmark, city, country) values($1,$2,$3,$4,$5,$6,$7) returning id;`
	err = tx.QueryRow(query, req.PickupLocation.Lat, req.PickupLocation.Lng, req.PickupLocation.Name,
		req.PickupLocation.StreetName, req.PickupLocation.Landmark, req.PickupLocation.City, req.PickupLocation.Country).Scan(&pickupLocId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("createRideReq", req)).
			Error("error creating customer task")
		tx.Rollback()
		return err
	}

	err = tx.QueryRow(query, req.DropLocation.Lat, req.DropLocation.Lng, req.DropLocation.Name,
		req.DropLocation.StreetName, req.DropLocation.Landmark, req.DropLocation.City, req.DropLocation.Country).Scan(&dropLocId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("createRideReq", req)).
			Error("error creating customer task")
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		log.L.With(zap.Error(err)).Error("error creating customer task")
		return err
	}

	return d.createRideStops(customerTaskId, pickupLocId, dropLocId)
}

func (d *DaoImpl) createRideStops(customerTaskId, pickupLocId, dropLocId string) error {
	query := `insert into ride_stops(customer_task_id, location_id, prev_location_id, next_location_id) 
			  values ($1, $2, $3, $4), ($5, $6, $7, $8);`

	_, err := d.db.Exec(query, customerTaskId, pickupLocId, pickupLocId, dropLocId, customerTaskId, dropLocId,
		pickupLocId, dropLocId)
	if err != nil {
		log.L.With(zap.Error(err)).Error("error in creating ride stops")
	}

	return err
}

func (d *DaoImpl) UpdateRide(req UpdateRideReq) error {
	return nil
}

func (d *DaoImpl) CancelRide(customerTaskId string) error {
	tx, err := d.db.Begin()
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("customerTaskId", customerTaskId)).
			Error("error cancelling customer task")
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	query := `update customer_task set status='CANCELLED', updated_at=$1 where id=$2`

	_, err = tx.Exec(query, time.Now().UTC(), customerTaskId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("customerTaskId", customerTaskId)).
			Error("error cancelling customer task")
	}

	query = `update driver_task set status='CANCELLED', updated_at=$1 where customer_task_id=$2`

	_, err = tx.Exec(query, time.Now().UTC(), customerTaskId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("customerTaskId", customerTaskId)).
			Error("error cancelling customer task")
	}

	if err = tx.Commit(); err != nil {
		log.L.With(zap.Error(err)).Error("error cancelling customer task")
	}

	return err
}

func (d *DaoImpl) GetHistory(customerId string) ([]CustomerHistoryResponse, error) {
	query := `select a.id, a.status, a.payable_amount, a.created_on, b.rating, b.driver_id from customer_task a 
    		  INNER JOIN driver_task b on a.id=b.customer_task_id where a.customer_id=$1`

	rows, err := d.db.Query(query, customerId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("customerId", customerId)).
			Error("error getting customer ride history")
	}

	var resp []CustomerHistoryResponse

	for rows.Next() {
		var customerTaskId, customerTaskStatus, driverId, rating string
		var createdOn time.Time
		var payableAmount float64
		var ratingPtr *int64

		err = rows.Scan(&customerTaskId, &customerTaskStatus, &payableAmount, &createdOn, &ratingPtr, &driverId)
		if err != nil {
			log.L.With(zap.Error(err), zap.Any("customerId", customerId)).
				Error("error scanning row while fetching customer ride history")
		}

		if ratingPtr != nil {
			rating = "null"
		} else {
			rating = strconv.FormatInt(*ratingPtr, 64)
		}

		resp = append(resp, CustomerHistoryResponse{
			RideId:        customerTaskId,
			RideStops:     d.getRideStops(customerTaskId),
			Status:        customerTaskStatus,
			PayableAmount: payableAmount,
			PaymentStatus: "done",
			DriverInfo:    d.getDriverInfo(driverId),
			RatingGiven:   rating,
			DateOfJourney: createdOn.Format(time.RFC850),
		})

	}

	return resp, nil
}

func (d *DaoImpl) getDriverInfo(driverId string) DriverInfo {
	query := `select first_name, last_name, phone from users where driver_id=$1`

	rows, err := d.db.Query(query, driverId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("driverId", driverId)).
			Error("error getting driver details")
	}

	var info DriverInfo
	var firstName, lastName string

	for rows.Next() {
		err = rows.Scan(&firstName, &lastName, &info.PhoneNo)
	}

	info.Name = firstName + " " + lastName

	return info
}

func (d *DaoImpl) getRideStops(customerTaskId string) []address.Location {
	query := `select a.lat, a.lng, a.name, a.landmark, a.street_name, a.city, a.country from address a 
    		  inner join ride_stops b on a.id=b.location_id where b.customer_task_id=$1 order by b.id `

	rows, err := d.db.Query(query, customerTaskId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("customerTaskId", customerTaskId)).
			Error("error getting ride stops details")
	}

	var stops []address.Location

	for rows.Next() {
		var loc address.Location
		err = rows.Scan(&loc.Lat, &loc.Lng, &loc.Name, &loc.Landmark, &loc.StreetName, &loc.City, &loc.Country)
		if err != nil {
			log.L.With(zap.Error(err), zap.Any("customerTaskId", customerTaskId)).
				Error("error scanning stop")
		}
		stops = append(stops, loc)
	}

	return stops
}
