package vehicle

import (
	"database/sql"
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/atulsinha007/uber/pkg/postgres"
	"go.uber.org/zap"
)

type Dao interface {
	CreateVehicle(request CreateVehicleRequest) (string, error)
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

func (d *DaoImpl) CreateVehicle(request CreateVehicleRequest) (string, error) {
	tx, err := d.db.Begin()
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("createRideReq", request)).
			Error("error creating customer task")
		return "", err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var id string

	query := `insert into vehicle(model, registration_no) values($1, $2) returning id;`
	err = tx.QueryRow(query, request.Model, request.RegistrationNo).Scan(&id)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("req", request)).Error("error in inserting vehicle")
		tx.Rollback()
	}

	query = `insert into vehicle_ride_type_mapping(vehicle_id, ride_type) values($1, $2);`

	for _, rideType := range request.PermittedRideTypes {
		_, err = tx.Exec(query, id, rideType)
		if err != nil {
			log.L.With(zap.Error(err), zap.Any("req", request)).
				Error("error in inserting vehicleRideTypeMapping")
			tx.Rollback()
		}
	}

	if err = tx.Commit(); err != nil {
		log.L.With(zap.Error(err)).Error("error creating vehicle")
		return "", err
	}

	return id, nil
}
