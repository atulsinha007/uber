package driverTask

import (
	"database/sql"
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/atulsinha007/uber/pkg/postgres"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Dao interface {
	GetDriverHistory(driverId string) ([]DriverHistoryResponse, error)
	AcceptRideRequest(req AcceptRideReq) error
	UpdateRide(req UpdateRideReq) error
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

func (d *DaoImpl) GetDriverHistory(driverId string) (resp []DriverHistoryResponse, err error) {
	query := `select(payable_amount, distance, rating, status) from driver_task where driver_id=$1 order by created_at desc`

	rows, err := d.db.Query(query, driverId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("driverId", driverId)).Error("error getting driver history")
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.L.With(zap.Error(e)).Error("error getting driver history")
		}
	}()

	for rows.Next() {
		var row DriverHistoryResponse
		if err = rows.Scan(&row.PayoutAmount, &row.DistanceCovered, &row.Rating, &row.Status); err != nil {
			return
		}

		resp = append(resp, row)
	}

	return
}

func (d *DaoImpl) AcceptRideRequest(req AcceptRideReq) error {
	query := `update driver_task set status='ACCEPTED' where driver_task_id=$1`

	_, err := d.db.Exec(query, req.DriverTaskId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("req", req)).Error("error in accepting ride request")
	}

	return err
}

func (d *DaoImpl) UpdateRide(req UpdateRideReq) error {
	query := `update driver_task set status=$1 where driver_task_id=$2`

	_, err := d.db.Exec(query, req.Status, req.DriverTaskId)
	if err != nil {
		log.L.With(zap.Error(err), zap.Any("req", req)).Error("error in updating ride request")
	}

	return err
}
