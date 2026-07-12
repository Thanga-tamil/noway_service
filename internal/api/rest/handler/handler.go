package handler

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"encoding/json"
	"gateway/internal/dto"
	"gateway/internal/repository"
	"gateway/internal/utils"
	"net/http"

	"github.com/google/uuid"
)


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

func parseInputFromReq(w http.ResponseWriter, req *http.Request) dto.UserRegisterReqPayload {

	var user dto.UserRegisterReqPayload
	err := json.NewDecoder(req.Body).Decode(&user)

	// EOF : end of file error might occur
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

	return user
}

func validateInput(w http.ResponseWriter, user dto.UserRegisterReqPayload) error {
	if len(strings.TrimSpace(user.Username)) == 0 { 
		return errors.New(utils.USERNAME_EMPTY)
	} else if len(strings.TrimSpace(user.EmailID)) == 0 { 
		return errors.New(utils.EMAILID_EMPTY)
	} else if len(strings.TrimSpace(user.MobileNumber)) == 0 { 
		return errors.New(utils.MOBILENUMBER_EMPTY)
	} 

	return nil
}
