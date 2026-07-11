package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi"

	"gateway/internal/api/rest/router"
	"gateway/internal/app"
)

const (
	SERVER_PORT = ":6969"
)

func main() {
	log.Println("serve http request response")

	app.App()

	chiRouter := chi.NewRouter()

	router.Route(chiRouter)

	http.ListenAndServe(SERVER_PORT, chiRouter)


}
