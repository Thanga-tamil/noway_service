package config

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/redis/go-redis/v9"
)

var GoRedis *redis.Client

var ctx = context.Background()

func InitRedis(c Cfg) (string, error) {

	logrus.Info("Initializing redis instance")

	addr := c.Rcache.Host + ":" + c.Rcache.Port

	logrus.Printf("redis address '%s'\n", addr)

	// Connect to Redis
	GoRedis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // No password for local development
		DB:       0, // Default DB
	})

	// Ping the Redis server to check the connection
	pong, err := GoRedis.Ping(ctx).Result()

	return pong, err
}


