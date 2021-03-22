package driver

import (
	_ "github.com/lib/pq"
)
//
//type Dao interface {
//	Set(loc Location) (string, error)
//	Get(id string) (Location, error)
//}
//
//type DaoImpl struct {
//	db *sql.DB
//}
//
//func NewDaoImpl(conf postgres.PgConf, dbName string) (*DaoImpl, error) {
//	conn, err := postgres.GetDbConn(conf, dbName)
//	if err != nil {
//		return nil, err
//	}
//
//	return &DaoImpl{db: conn}, nil
//}
//
//func (*DaoImpl) Set(loc Location) (string, error) {
//	return "", nil
//}
//
//func (*DaoImpl) Get(id string) (Location, error) {
//	return Location{}, nil
//}
