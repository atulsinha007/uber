package address

//import (
//	"database/sql"
//	"github.com/atulsinha007/uber/pkg/postgres"
//	_ "github.com/lib/pq"
//)
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
//func NewDaoImpl(conf postgres.PgConf) (*DaoImpl, error) {
//	conn, err := postgres.GetDbConn(conf)
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
