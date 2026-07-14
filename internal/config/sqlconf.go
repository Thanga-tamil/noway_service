package config

import (
	"database/sql"
	"github.com/sirupsen/logrus"

	// Import the SQLite driver
	// The underscore import registers the driver with database/sql
	_ "github.com/mattn/go-sqlite3"
)

// var MasterDB *sql.DB
// update it postgresql for multi tenant support 
// for now, keep sqlite running
// var DBs *[]sql.DB

func InitSqlite(c Cfg) (*sql.DB, error) {

	logrus.Info("Initializing sqlite db")

	// since we're relaying on sqlite, 
	// there is no tcp connection, matter of fact.. 
	// we don't even have a database server :) 
	// isn't it awesome 
	// but it's only for small scale
	Type := c.Sqlite.Host
	port := c.Sqlite.Port

	logrus.Printf("sqlite type '%s' and file: '%s'\n", Type, port)
	
	sqlite, err := sql.Open(Type, port)

	return sqlite, err

}
