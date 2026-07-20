package repository

import (
	"time"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func SaveRegisterUser(db *sqlx.DB, useId string, username, mobileNumber, 
			emailID string, isDeleted bool, now time.Time) error  {

	stmt := `insert into 
				users (user_id, username, email_id, mobile_number, created_at, is_deleted)
		     values ($1, $2, $3, $4, $5, $6)`

	result, err := db.Exec(stmt, useId, username, emailID, mobileNumber, now, isDeleted)

	if err != nil {
		logrus.Error("error while inserting into table user:: ", err.Error())
		return err
	}

	logrus.Info("user register db insert result:: ", result)

	return nil
}
