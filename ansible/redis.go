package ansible

import (
	"fmt"
	"os"
	"time"
	"context"

	"github.com/redis/go-redis/v9"
)

func addNonAnsibleErrorRecord() {
    ctx := context.Background()
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
        Password: "",
        DB:       0,
    })
    defer rdb.Close()

    // Check connection first
    if err := rdb.Ping(ctx).Err(); err != nil {
        fmt.Printf("Error connecting to Redis: %v\n", err)
        return
    }

    // Set the value using the context
    key := time.Now().Format("2006-01-02 15:04:05")
    if err := rdb.Set(ctx, string(key), "test", 0).Err(); err != nil {
        fmt.Printf("Error setting value in Redis: %v\n", err)
        return
    }
}
