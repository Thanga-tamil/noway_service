package handler

import (
	"log"
	"time"
	"encoding/json"
	"net/http"
	"gateway/internal/config"
)

type User struct {
	Id int `json:"id"`
	UserId string `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
	IsDeleted bool `json:"is_deleted"`
}

func HandleRegister(w http.ResponseWriter, req *http.Request) {
    log.Println("req is in hanlder layer")
	sayHello(w, req)
	
}

// handler func(ResponseWriter, *Request)
func sayHello(w http.ResponseWriter, req *http.Request) { 
	defer log.Println("final defer")

	Map := map[string]string{"net/http: API caller Addr": req.RemoteAddr, "message": "hello"}
	val, _ := json.Marshal(Map) // conv map to bin
	_, _ = w.Write([]byte(val)) 

	stmt := `select * from users;`

	var user User

	start := time.Now()

	err := config.DB.QueryRow(stmt).Scan(&user.Id, &user.UserId, &user.IsAdmin, &user.IsDeleted)

    elapsed := time.Since(start)
    log.Printf("Binomial took %s", elapsed)

	if err != nil {
		log.Println("Error while scanning sqlite: ", err)
		return
	}

	log.Println("retrieved data from sqlite: ", user)

}
