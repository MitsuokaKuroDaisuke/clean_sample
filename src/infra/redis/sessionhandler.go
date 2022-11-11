package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

type SessionHandler struct {
	client *redis.Client
}

func NewSessionHandler() SessionHandler {

	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	sh := SessionHandler{}
	sh.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	return sh
}
