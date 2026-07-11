package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi"

	"gateway/internal/api/rest/router"
)

const (
	SERVER_PORT = ":6969"
)

func main() {
	log.Println("serve http request response")

	chiRouter := chi.NewRouter()

	router.Route(chiRouter)

	http.ListenAndServe(SERVER_PORT, chiRouter)

}
