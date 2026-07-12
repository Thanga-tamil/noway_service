package repository

import (
	"time"
	"database/sql"
	"github.com/google/uuid"
	"gateway/internal/config"
)

type User struct {
	ID           int       `json:"id" db:"id"`
	UserID       string    `json:"user_id" db:"user_id"`
	Username     string    `json:"username" db:"username"`
	MobileNumber string    `json:"mobilenumber" db:"mobilenumber"`
	EmailID      string    `json:"email_id" db:"email_id"`
	IsDeleted    bool      `json:"is_deleted" db:"is_deleted"`
	Image        string    `json:"image" db:"image"`
	ThumbImage   string    `json:"thumbimage" db:"thumbimage"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

func SaveRegisterUser(useId uuid.UUID, username, mobileNumber, emailID string, IsDeleted bool, now time.Time) (sql.Result, error)  {

	Q := `
		INSERT INTO USERS 
		(user_id, username, mobilenumber, email_id, 
		is_deleted, created_at) 
		VALUES (?, ?, ?, ?, ?, ?)
		`
	result, err := config.MasterDB.Exec(Q, useId, username, mobileNumber, emailID, IsDeleted, now)

	return result, err
}
