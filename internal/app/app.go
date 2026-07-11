package app

import (
	"gateway/internal/config"
)

func App(c config.Cfg) {

	config.InitRedis(c)
	
	config.InitSqlite(c)
}
