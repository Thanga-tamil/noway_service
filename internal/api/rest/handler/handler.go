package handler

import (
	"fmt"
	"log"
	"time"
	"strings"
	"errors"

	"encoding/json"
	//"gateway/internal/config"
	"gateway/internal/repository"
	"net/http"

	"github.com/google/uuid"
)

const (
	USERNAME_EMPTY = "username must not be null or empty"
	EMAILID_EMPTY = "email id must not be null or empty"
	MOBILENUMBER_EMPTY = "mobile number must not be null or empty"
)
type UserRegisterReqPayload struct {
	Username     string    `json:"username" db:"username"`
	MobileNumber string    `json:"mobilenumber" db:"mobilenumber"`
	EmailID      string    `json:"email_id" db:"email_id"`
	Image        string    `json:"image" db:"image"`
}

func HandleUserRegister(w http.ResponseWriter, req *http.Request) {

	user := parseInputFromReq(w, req)
	fmt.Printf("parsed register user input: %#v\n", user)

	userId := uuid.New()

	log.Printf("generated uuid: %v for email_id: %s", userId, user.EmailID)

	result, err := repository.SaveRegisterUser(userId, user.Username, 
						user.MobileNumber, user.EmailID, false, time.Now())

	if err != nil {
		resp := map[string]any{"message": err.Error(), "status": 400}
		val, _ := json.Marshal(resp)
		w.Write([]byte(val))
	}

	fmt.Println(result)

	resp := map[string]any{
				"message": "Registration completed successfully", 
		 		"status": 200}
	val, _ := json.Marshal(resp)
	w.Write([]byte(val))
}

func parseInputFromReq(w http.ResponseWriter, req *http.Request) UserRegisterReqPayload {
	var user UserRegisterReqPayload

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		resp := map[string]any{
					"message": "Request body must not be null",
					"status": 400}

		val, _ := json.Marshal(resp)

		w.Write([]byte(val))

		log.Println("Error while Decode input payload: ", err)
	} else {
		err := validateInput(w, user)
		if err != nil {
			msg := map[string]any{"message": err.Error(), "status": 400}

			val, _ := json.Marshal(msg)
			w.Write([]byte(val))

			log.Println("input payload validation error: ", err)
		}
	}

	// EOF : end of file error might occur
	return user
}

func validateInput(w http.ResponseWriter, user UserRegisterReqPayload) error {
	if len(strings.TrimSpace(user.Username)) == 0 { 
		return errors.New(USERNAME_EMPTY)
	} else if len(strings.TrimSpace(user.EmailID)) == 0 { 
		return errors.New(EMAILID_EMPTY)
	} else if len(strings.TrimSpace(user.MobileNumber)) == 0 { 
		return errors.New(MOBILENUMBER_EMPTY)
	} 

	return nil
}
