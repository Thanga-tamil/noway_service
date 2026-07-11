package config

import (
	"log"
	"context"
	"github.com/redis/go-redis/v9"
)

func InitRedis() {

	log.Println("Initializing redis instance")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	ctx := context.Background()
	
	res := rdb.Get(ctx, "one")

	log.Println("retrieved result from redis: ", res)
}
