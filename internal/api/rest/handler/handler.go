package handler

import (
	"log"
	"encoding/json"
	"net/http"
)

func HandleRegister(w http.ResponseWriter, req *http.Request) {
    log.Println("req is in hanlder layer")
	sayHello(w, req)
	
}

// handler func(ResponseWriter, *Request)
func sayHello(w http.ResponseWriter, req *http.Request) { 
	defer log.Println("final defer")

	Map := map[string]string{"net/http: API caller Addr": req.RemoteAddr, "message": "hello"}
	val, _ := json.Marshal(Map)
	_, _ = w.Write([]byte(val))

	log.Println("I'm anticipating this to be logged in console after wrting the response to client")
	
}
