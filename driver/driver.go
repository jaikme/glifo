package driver

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// DB ...
type DB struct {
	SQL *sql.DB
}

// DBConn ...
var dbConn = &DB{}

// ConnectSQL ...
func ConnectSQL(dbname string) (*DB, error) {
	d, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		panic(err)
	}

	dbConn.SQL = d.DB()

	return dbConn, err
}
