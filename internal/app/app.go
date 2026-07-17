package app

import (
	//"database/sql"
	"github.com/sirupsen/logrus"

	"gateway/internal/config"
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

 
	logrus.Info("Required services initialization completed successfully")
}
