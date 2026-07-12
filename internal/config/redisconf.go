package config

import (
	"log"
	"github.com/redis/go-redis/v9"
)

var GoRedis *redis.Client

func InitRedis(c Cfg) {

	log.Println("Initializing redis instance")

	addr := c.Rcache.Host + ":" + c.Rcache.Port

	log.Printf("redis address '%s'\n", addr)

	GoRedis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	log.Println("Redis init success")

	// ctx := context.Background()
	// res := rdb.Get(ctx, "one")
}
