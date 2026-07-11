package main

import (
	"log"
	"net/http"

	"gateway/internal/api/rest/router"
)

func main() {
	log.Println("serve http request response")

	r := router.Route()

	http.ListenAndServe(":6969", r)

}
