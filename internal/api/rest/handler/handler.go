package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"

	"gateway/internal/dto"
	"gateway/internal/service"
)


func HandleUserRegister(w http.ResponseWriter, req *http.Request) {
	user, err := parseInputFromReq(w, req)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	logrus.Printf("parsed register user input: %#v\n", user)

	jwt, err := service.RegisterService(user)

	if err != nil {
		resp := map[string]any{"status": 400, "message": err.Error()}
		val, _ := json.Marshal(resp)
		w.Write([]byte(val))
		return
	}

	resp := map[string]any{
				"status": 200,
				"jwtToken": jwt,
				"message": "Registration completed successfully"}

	val, _ := json.Marshal(resp)
	w.Write([]byte(val))
}

func parseInputFromReq(w http.ResponseWriter,
	 req *http.Request) (dto.UserRegisterReqPayload, error) {

	var user dto.UserRegisterReqPayload
	err := json.NewDecoder(req.Body).Decode(&user)

	// EOF : end of file error might occur
	if err != nil {
		logrus.Info("Error while Decode input payload: ", err)
		resp := map[string]any{
					"status": 400,
					"message": "Request body must not be null"}
		val, _ := json.Marshal(resp)
		return dto.UserRegisterReqPayload{}, errors.New(string(val))
	} 
	
	if err := service.ValidateInput(w, user); err != nil {
		logrus.Error("input payload validation error: ", err)

		msg := map[string]any{"status": 400, "message": err.Error()}
		val, _ := json.Marshal(msg)
		return dto.UserRegisterReqPayload{}, errors.New(string(val))
	}

	return user, nil
}


func GenerateJwtToken(w http.ResponseWriter, req *http.Request) {

	userId := req.Header.Get("userId")

	if len(userId) == 0 {
		resp := map[string]any{"message": "userId can not be null or empty",
						"status": 400}

		val, _ := json.Marshal(resp)
		w.Write([]byte(val))
		return
	}

	jwt, err := service.ServeJwt(userId)

	if err != nil {
		logrus.Printf("Error on return of ServeJwt: %s", err.Error())
		panic(err)
	}

	resp := map[string]any{ "status": 200, "jwtToken": jwt}
	val, _ := json.Marshal(resp)
	w.Write([]byte(val))
}

