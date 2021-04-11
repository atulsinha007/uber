package user

import (
	"database/sql"
	"fmt"
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/atulsinha007/uber/pkg/postgres"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"strconv"
	"time"
)

//go:generate mockgen -destination=mock_dao.go -package=user -source=./dao.go
type Dao interface {
	Set(user User) (string, error)
	GetDriverProfile(driverId int) (DriverProfileResponse, error)
	UpdateLocation(req UpdateCurrentLocationRequest) error
	AddDriverWithVehicle(vehicleId string, user User) error
	GetDriverHistory(driverId int) ([]DriverHistoryResponse, error)
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

func (d *DaoImpl) Set(user User) (string, error) {
	query := `insert into users(first_name, last_name, phone, user_type) values($1, $2, $3, $4) returning user_id;`

	var id string
	err := d.db.QueryRow(query, user.FirstName, user.LastName, user.Phone, user.Type).Scan(&id)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("user", user)).Error("error in inserting user")
	}

	return id, err
}

func (d *DaoImpl) GetDriverProfile(driverId int) (profile DriverProfileResponse, err error) {
	query := `select first_name, last_name, phone from users where user_id=$1;`

	var firstName, lastName, phoneNo string
	row := d.db.QueryRow(query, driverId)

	err = row.Scan(&firstName, &lastName, &phoneNo)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("driverId", driverId)).Error("error in fetching driver")
		return
	}

	query = "select count(*), avg(rating) from driver_task where driver_id=$1 and status='COMPLETED';"

	var rating *float64
	err = d.db.QueryRow(query, driverId).Scan(&profile.TotalRides, &rating)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("driverId", driverId)).Error("error in fetching driver details")
	}

	profile.Name = firstName + " " + lastName
	profile.PhoneNo = phoneNo

	if rating != nil {
		profile.AverageRating = strconv.FormatFloat(*rating, 'f', -1, 64)
	} else {
		profile.AverageRating = "NULL"
	}

	return
}

func (d *DaoImpl) UpdateLocation(req UpdateCurrentLocationRequest) error {
	query := `update users set current_lat=$1, current_lng=$2, updated_at=$3 where user_id=$4`

	_, err := d.db.Exec(query, req.CurLocation.Lat, req.CurLocation.Lng, time.Now().UTC(), req.UserId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("req", req)).Error("error in updating location")
	}

	return err
}

func (d *DaoImpl) AddDriverWithVehicle(vehicleId string, user User) error {
	id, err := d.Set(user)
	if err != nil {
		log.L.With(zap.Error(err)).Error("error creating driver with vehicle")
		return err
	}

	query := `insert into driver_profile(driver_id, vehicle_id) values($1, $2);`

	_, err = d.db.Exec(query, id, vehicleId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("vehicleId", vehicleId), zap.Any("user", user)).
			Error("error creating driver with vehicle")
	}
	return err
}

func (d *DaoImpl) GetDriverHistory(driverId int) (resp []DriverHistoryResponse, err error) {
	query := `select payable_amount, distance, rating, status from driver_task where driver_id=$1 order by created_at desc`

	rows, err := d.db.Query(query, driverId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("driverId", driverId)).Error("error getting driver history")
		return nil, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.L.With(zap.Error(e)).Error("error getting driver history")
		}
	}()

	var rating *int
	for rows.Next() {
		var row DriverHistoryResponse
		if err = rows.Scan(&row.PayoutAmount, &row.DistanceCovered, &rating, &row.Status); err != nil {
			return
		}

		if rating == nil {
			row.Rating = 3
		} else {
			row.Rating = *rating
		}

		fmt.Println(row)

		resp = append(resp, row)
	}

	return
}
