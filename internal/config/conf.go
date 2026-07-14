package config

import (
	"os"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type Cfg struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Sqlite sqlite `json:"sqlite"`
	Rcache redisCache `json:"redis"`
}

type sqlite struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
	
type redisCache struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

const(
	CONF_PATH = "config.json"
)

func LoadConfig() Cfg {

	file, err := os.Open(CONF_PATH)

	if err != nil {
		logrus.Fatalf("Error while opening %s file \n", err.Error())
	}

	var c Cfg 
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		logrus.Fatalf("Error decoding JSON: %v", err)
	}
	logrus.Printf("decoded config.json file: %#v\n", c)

	return c
}
