package redis

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
    Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
    Password: "",
    DB:       0,
})

type RedisValue struct {
    Sha string `json:"sha"`
    Status string `json:"status"`
    Message string `json:"message"`
}