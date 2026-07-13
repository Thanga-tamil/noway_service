package handler

import (
	"fmt"
	"log"

	"net/http"
	"encoding/json"
	"gateway/internal/dto"
	"gateway/internal/service"
)


func HandleUserRegister(w http.ResponseWriter, req *http.Request) {

	user := parseInputFromReq(w, req)
	fmt.Printf("parsed register user input: %#v\n", user)

	result, err := service.RegisterService(user)

	if err != nil {
		resp := map[string]any{"status": 400, "message": err.Error()}
		val, _ := json.Marshal(resp)
		w.Write([]byte(val))
	}

	fmt.Println(result)

	resp := map[string]any{
				"status": 200,
				"message": "Registration completed successfully"}

	val, _ := json.Marshal(resp)
	w.Write([]byte(val))
}

func parseInputFromReq(w http.ResponseWriter, 
			req *http.Request) dto.UserRegisterReqPayload {

	var user dto.UserRegisterReqPayload
	err := json.NewDecoder(req.Body).Decode(&user)

	// EOF : end of file error might occur
	if err != nil {
		resp := map[string]any{
					"status": 400,
					"message": "Request body must not be null"}

		val, _ := json.Marshal(resp)

		w.Write([]byte(val))

		log.Println("Error while Decode input payload: ", err)
	} else {
		err := service.ValidateInput(w, user)
		if err != nil {
			msg := map[string]any{"status": 400, "message": err.Error()}

			val, _ := json.Marshal(msg)
			w.Write([]byte(val))

			log.Println("input payload validation error: ", err)
		}
	}

	return user
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
		log.Printf("Error on return of ServeJwt: %s", err.Error())
		panic(err)
	}

	resp := map[string]any{ "status": 200, "jwtToken": jwt}
	val, _ := json.Marshal(resp)
	w.Write([]byte(val))
}

