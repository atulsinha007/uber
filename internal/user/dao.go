package user

import (
	"database/sql"
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/atulsinha007/uber/pkg/postgres"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type Dao interface {
	Set(user User) error
	GetDriverProfile(driverId string) (DriverProfileResponse, error)
	UpdateLocation(req UpdateCurrentLocationRequest) error
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

func (d *DaoImpl) Set(user User) error {
	query := `insert into users(first_name, last_name, phone, user_type) values($1, $2, $3, $4);`

	_, err := d.db.Exec(query, user.FirstName, user.LastName, user.Phone, user.Type)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("user", user)).Error("error in inserting user")
	}

	return err
}

func (d *DaoImpl) GetDriverProfile(driverId string) (profile DriverProfileResponse, err error) {
	query := `select first_name, last_name, phone from users where user_id=$1;`

	var firstName, lastName, phoneNo string
	row := d.db.QueryRow(query, driverId)

	err = row.Scan(&firstName, &lastName, &phoneNo)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("driverId", driverId)).Error("error in fetching driver")
		return
	}

	query = "select count(*), avg(rating) from driver_task where driver_id=$1 and rating is not null and status='COMPLETED';"

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
