package config

import (
	"log"
	"os"
	"encoding/json"
)

// {
//     "host": "localhost",
//     "port": "6969",
//     "sqlite": {
//         "host": "sqlite3",
//         "port": "chats.db"
//     },
//     "redis": {
//         "host": "localhost",
//         "port": "6379"
//     }
// }

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

func LoadConfig() Cfg {

	file, err := os.Open("config.json")

	if err != nil {
		log.Fatalf("Error while opening %s file \n", err.Error())
	}

	var c Cfg 
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	log.Println("decoded config.json file: ", c)

	return c
}
