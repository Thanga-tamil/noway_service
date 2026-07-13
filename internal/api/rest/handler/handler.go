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
		err := service.ValidateInput(w, user)
		if err != nil {
			msg := map[string]any{"message": err.Error(), "status": 400}

			val, _ := json.Marshal(msg)
			w.Write([]byte(val))

			log.Println("input payload validation error: ", err)
		}
	}

	return user
}
