package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func AddSyncRecord(sha string, status string, message string) {
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

func GetAllSyncRecords() ([]string, error){
    ctx := context.Background()
    defer rdb.Close()

    if err := rdb.Ping(ctx).Err(); err != nil {
        return nil, fmt.Errorf("Error connecting to Redis: %v\n", err)
    }

    keys := rdb.Keys(ctx, "*").Val()
    return keys, nil
}