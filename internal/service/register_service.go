package service

import (
	"time"
	"errors"
	"strings"
	"net/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"gateway/internal/dto"
	"gateway/internal/utils"
	"gateway/internal/repository"
)

func RegisterService(user dto.UserRegisterReqPayload) (string, error) {

	userId := uuid.New()

	logrus.Printf("generated uuid: %v for email_id: %s", userId, user.EmailID)

	result, err := repository.SaveRegisterUser(userId, user.Username, 
						user.MobileNumber, user.EmailID, false, time.Now())

	logrus.Info("sqlite user register insertion result: ", result)

	jwt, err := ServeJwt(userId.String())

	if err != nil {
		logrus.Errorf("Error on return of ServeJwt: %s", err.Error())
		panic(err)
	}

	return jwt, err
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
