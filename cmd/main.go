package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"gateway/internal/api/rest/router"
	"gateway/internal/app"
	"gateway/internal/config"
)



func main() {

	log.Println("serve http request response")

	conf := config.LoadConfig()

	app.App(conf)

	chiRouter := chi.NewRouter()

	router.Route(chiRouter)

	serverAddr := conf.Host + ":" + conf.Port

	http.ListenAndServe(serverAddr, chiRouter)

	log.Println("HTTP server started with addr: ", serverAddr)

}
