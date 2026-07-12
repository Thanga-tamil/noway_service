package service

import (
	"log"
	"time"
	"errors"
	"strings"
	"net/http"
	"database/sql"
	"github.com/google/uuid"

	"gateway/internal/dto"
	"gateway/internal/utils"
	"gateway/internal/repository"
)

func RegisterService(user dto.UserRegisterReqPayload) (sql.Result, error) {

	userId := uuid.New()

	log.Printf("generated uuid: %v for email_id: %s", userId, user.EmailID)

	result, err := repository.SaveRegisterUser(userId, user.Username, 
						user.MobileNumber, user.EmailID, false, time.Now())


	return result, err
}

func ValidateInput(w http.ResponseWriter, user dto.UserRegisterReqPayload) error {
	if len(strings.TrimSpace(user.Username)) == 0 { 
		return errors.New(utils.USERNAME_EMPTY)
	} else if len(strings.TrimSpace(user.EmailID)) == 0 { 
		return errors.New(utils.EMAILID_EMPTY)
	} else if len(strings.TrimSpace(user.MobileNumber)) == 0 { 
		return errors.New(utils.MOBILENUMBER_EMPTY)
	} 

	return nil
}
