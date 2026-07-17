package config

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var MasterDB *sqlx.DB

func InitSql(c Cfg) {

	 db, err := sqlx.Open(c.Postgre.Type, get_db_uri(c))

	 if err != nil {
		 logrus.Error("error while opening tcp connection with postgreSQL: ", err.Error())
		 panic(err)
	 }

	 lookupDB := "SET SEARCH_PATH TO connection_conf;"

	 if _, err := db.ExecContext(ctx, lookupDB); err != nil {
		 logrus.Error("error when executing search path for tenant config: ", err.Error())
		 panic(err)
	 } else {
		 MasterDB = db
	 }

	 // DEV in progress logs ** unneccessary logs should be removed once crafting completed ** 
}


func get_db_uri(c Cfg) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	c.Postgre.Host, c.Postgre.Port, c.Postgre.User, c.Postgre.Password, c.Postgre.Dbname)
}
