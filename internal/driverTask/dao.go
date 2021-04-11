package driverTask

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"github.com/atulsinha007/uber/internal/address"
	"github.com/atulsinha007/uber/pkg/distance_util"
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/atulsinha007/uber/pkg/postgres"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"math"
	"strconv"
	"time"
)

//go:generate mockgen -destination=mock_dao.go -package=driverTask -source=./dao.go
type Dao interface {
	AcceptRideRequest(req AcceptRideReq) error
	UpdateRide(req UpdateRideReq) error
	GetFromDriverIdAndCustomerTaskId(customerTaskId, driverId int) (DriverTask, error)
	FindNearestDriver(pickupLocation address.Location, preferredRideType string) (int, float64, error)
	CreateDriverTask(task DriverTask) error
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

func (d *DaoImpl) AcceptRideRequest(req AcceptRideReq) error {
	tx, err := d.db.Begin()
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("req", req)).
			Error("error creating driver task")
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	query := `update driver_task set status='ACCEPTED', updated_at=$1 where id=$2`

	_, err = tx.Exec(query, time.Now().UTC(), req.DriverTaskId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("req", req)).Error("error in accepting ride request")
		tx.Rollback()
		return err
	}

	query = `update driver_profile set is_available=$1, updated_at=$2 where driver_id=$3;`
	_, err = tx.Exec(query, false, time.Now().UTC(), req.DriverId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("req", req)).Error("error in updating is_available for driver")
		tx.Rollback()
		return err
	}

	query = `update customer_task set status=$1, updated_at=$2 where id=$3;`
	_, err = tx.Exec(query, "ONGOING", time.Now().UTC(), req.CustomerTaskId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("req", req)).Error("error in updating customer_task status")
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		log.L.With(zap.Error(err)).Error("error accepting driver task")
		return err
	}

	return nil
}

func (d *DaoImpl) UpdateRide(req UpdateRideReq) error {
	tx, err := d.db.Begin()
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("req", req)).
			Error("error updating driver task")
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var driverId int
	query := `update driver_task set status=$1 where id=$2 returning driver_id`
	err = tx.QueryRow(query, req.Status, req.DriverTaskId).Scan(&driverId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("req", req)).Error("error in updating ride request")
	}

	if req.Status == "COMPLETED" {
		query = `update driver_profile set is_available=$1, updated_at=$2 where driver_id=$3;`
		_, err = tx.Exec(query, true, time.Now().UTC(), driverId)
		if err != nil {
			log.L.With(zap.Error(err), zap.Any("req", req)).Error("error in updating is_available for driver")
			tx.Rollback()
			return err
		}

		query = `update customer_task set status=$1, updated_at=$2 where id=$3;`
		_, err = tx.Exec(query, "COMPLETED", time.Now().UTC(), req.CustomerTaskId)
		if err != nil {
			log.L.With(zap.Error(err), zap.Any("req", req)).Error("error in updating customer_task status")
			tx.Rollback()
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		log.L.With(zap.Error(err)).Error("error updating driver task")
		return err
	}

	return err
}

func (d *DaoImpl) GetFromDriverIdAndCustomerTaskId(customerTaskId, driverId int) (DriverTask, error) {
	query := `select id, status, payable_amount, ride_type, distance from driver_task where customer_task_id=$1 and driver_id=$2;`

	var dt DriverTask
	err := d.db.QueryRow(query, customerTaskId, driverId).Scan(&dt.DriverTaskId, &dt.Status, &dt.PayableAmount,
		&dt.RideType, &dt.Distance)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("customerTaskId", customerTaskId)).
			Error("error getting driverTask from customerTaskId")
		return DriverTask{}, err
	}

	dt.CustomerTaskId = customerTaskId
	dt.DriverId = driverId

	return dt, err
}

func (d *DaoImpl) FindNearestDriver(pickupLocation address.Location, preferredRideType string) (int, float64, error) {
	driverIds, err := d.findSameRideTypeDriverIds(preferredRideType)
	if err != nil {
		return 0, 0, err
	}

	if len(driverIds) == 0 {
		return 0, 0, errors.New("no nearby driver found")
	}

	buf := bytes.NewBufferString("select user_id, current_lat, current_lng from users where user_id IN(")
	for i, v := range driverIds {
		if i > 0 {
			buf.WriteString(",")
		}

		buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteString(")")

	fmt.Println(buf.String())

	rows, err := d.db.Query(buf.String())
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("preferredRideType", preferredRideType)).
			Error("error getting driver locations")
		return 0, 0, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.L.With(zap.Error(e)).Error("error getting nearest driver")
		}
	}()

	var nearestDriverId int
	dist := math.Inf(1)

	for rows.Next() {
		var lat, lng float64
		var driverId int

		err = rows.Scan(&driverId, &lat, &lng)
		if err != nil {
			return 0, 0, err
		}
		di := distance_util.Haversine(pickupLocation.Lat, pickupLocation.Lng, lat, lng)
		fmt.Println("found driver: ", driverId, lat, lng, di)
		if di < dist {
			nearestDriverId = driverId
			dist = di
		}
	}

	fmt.Println(nearestDriverId, dist)

	return nearestDriverId, dist, nil

}

func (d *DaoImpl) findSameRideTypeDriverIds(preferredRideType string) ([]int, error) {

	query := `select a.driver_id from driver_profile a INNER JOIN 
              vehicle_ride_type_mapping b ON a.vehicle_id=b.vehicle_id where b.ride_type=$1 and a.is_available=$2;`

	rows, err := d.db.Query(query, preferredRideType, true)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("preferredRideType", preferredRideType)).
			Error("error getting driver ids")
		return nil, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.L.With(zap.Error(e)).Error("error getting driver ids")
		}
	}()
	var driverIds []int

	for rows.Next() {
		var driverId int
		if err = rows.Scan(&driverId); err != nil {
			log.L.With(zap.Error(err), zap.Any("preferredRideType", preferredRideType)).
				Error("error getting driver id")
			return nil, err
		}

		driverIds = append(driverIds, driverId)
	}

	return driverIds, nil
}

func (d *DaoImpl) CreateDriverTask(task DriverTask) error {
	query := `insert into driver_task(driver_id, customer_task_id, distance, payable_amount, ride_type, status) 
			  values ($1, $2, $3, $4, $5, $6);`

	_, err := d.db.Exec(query, task.DriverId, task.CustomerTaskId, task.Distance, task.PayableAmount, task.RideType, task.Status)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("task", task)).Error("error in creating driver task")
		return err
	}

	return nil
}