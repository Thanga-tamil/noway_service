package service

import (
	"time"
	"errors"
	"strings"
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"gateway/internal/config"
	"gateway/internal/dto"
	"gateway/internal/repository"
	"gateway/internal/utils"
)

func RegisterService(db *sqlx.DB, userId string, user dto.UserRegisterReqPayload) {

	logrus.Infof("generated uuid: %v for email_id: %s", userId, user.EmailID)

	if err := repository.SaveRegisterUser(db, userId, user.Username, 
				  user.MobileNumber, user.EmailID, false, time.Now()); err != nil {
		logrus.Errorf("Error while saving register user: %s", err.Error())
		return
	}

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

func StoreJwtInRedis(userId, token string) error {
	logrus.Infof("store jwt token in redis cache for userId:: %s, with token=> %s", userId, token)
	ctx := context.Background()
	key := "user:" + userId + ":access"
	return config.GoRedis.Set(ctx, key, token, 0).Err()
}
