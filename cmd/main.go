package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"gateway/internal/api/rest/middleware"
	"gateway/internal/api/rest/router"
	"gateway/internal/app"
	"gateway/internal/config"
)


var (
    WarningLog *log.Logger
    InfoLog   *log.Logger
    ErrorLog   *log.Logger
)

func main() {
	
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		DisableColors:   false,
		// TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.Info("Serve http request response using engine 'net/http'")

	conf := config.LoadConfig()

	app.App(conf)

	r := chi.NewRouter()

	r.Use(middleware.MyMiddleware)

	r.Route("/api/v1", router.Route)

	http.ListenAndServe(conf.Host + ":" + strconv.Itoa(conf.Port), r)

}
