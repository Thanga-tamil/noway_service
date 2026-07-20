package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
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
	
	r := gin.Default() // r := gin.New()

	v1 := r.Group("/api/v1")

	v1.Use(middleware.MyMiddleware())
	
	router.Route(v1)

	r.Run(conf.Host + ":" + strconv.Itoa(conf.Port))

}

