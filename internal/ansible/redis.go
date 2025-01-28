package ansible

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisValue struct {
    Sha string `json:"sha"`
    Status string `json:"status"`
    Message string `json:"message"`
}

var rdb = redis.NewClient(&redis.Options{
    Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
    Password: "",
    DB:       0,
})

func addSyncRecord(sha string, status string, message string) {
    ctx := context.Background()
    defer rdb.Close()

    // Check connection first
    if err := rdb.Ping(ctx).Err(); err != nil {
        fmt.Printf("Error connecting to Redis: %v\n", err)
        return
    }

    // Set data
    key := time.Now().Format("2006-01-02 15:04:05")
    value, err := json.Marshal(RedisValue{sha, status, message})
    if err != nil {
        fmt.Printf("Error encoding json value: %v\n", err)
        return
    }

    // Set the value using the context
    if err := rdb.Set(ctx, string(key), string(value), 0).Err(); err != nil {
        fmt.Printf("Error setting value in Redis: %v\n", err)
        return
    }
}