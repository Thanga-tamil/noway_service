package app

import (
	//"log"
	"gateway/internal/config"
)

func App() {

	config.InitCqlShema("chat", "localhost:")
}
