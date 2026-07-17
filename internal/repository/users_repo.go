package repository

import (
	"time"
	"database/sql"
	"github.com/google/uuid"

)

func SaveRegisterUser(useId uuid.UUID, username, mobileNumber, emailID string, IsDeleted bool, now time.Time) (sql.Result, error)  {

	return nil, nil
}
