package app

import (
	"github.com/sirupsen/logrus"

	"gateway/internal/config"
	"gateway/internal/service"
)

func App(c config.Cfg) {

	logrus.Info("Initialize required services from app.go")

	config.InitSql(c)

	if pong, err := config.InitRedis(c); err != nil {
		logrus.Fatalf("Error connecting to Redis: %s", err)
	} else {
		logrus.Info("Connected to Redis:", pong)
		logrus.Info("Redis init success")
	}

	if err := service.LoadJwtSignKeyInCache(); err != nil {
		// todo
	}
 
	logrus.Info("Required services initialization completed successfully")
}
