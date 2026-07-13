package dto

type UserRegisterReqPayload struct {
	Username     string    `json:"username" db:"username"`
	MobileNumber string    `json:"mobilenumber" db:"mobilenumber"`
	EmailID      string    `json:"email_id" db:"email_id"`
	Image        string    `json:"image" db:"image"`
}
