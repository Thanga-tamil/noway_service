package main

import (
	"log"
	"strconv"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"gateway/internal/app"
	"gateway/internal/config"
	"gateway/internal/api/rest/router"
	"gateway/internal/api/rest/middleware"
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

	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.MyMiddleware)

	router.Route(chiRouter)

	serverAddr := conf.Host + ":" + strconv.Itoa(conf.Port)

	logrus.Info("Http server started in addr: ", serverAddr)

	http.ListenAndServe(serverAddr, chiRouter)

}
