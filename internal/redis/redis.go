package redis

import (
	"context"
	"encoding/json"
	"errors"
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

func AddSyncRecord(sha string, status string, message string) {
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
        Password: "",
        DB:       0,
    })

    ctx := context.Background()
    defer rdb.Close()

    // Check connection first
    if err := rdb.Ping(ctx).Err(); err != nil {
        fmt.Printf("Error connecting to Redis: %v\n", err)
        return
    }

    // Set data
    key := time.Now().Format("2006-01-02T15:04:05")
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

func GetSyncRecords(query string) (map[string]RedisValue, error) {
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
        Password: "",
        DB:       0,
    })

    ctx := context.Background()
    defer rdb.Close()

    // Check connection first
    if err := rdb.Ping(ctx).Err(); err != nil {
        return nil, fmt.Errorf("failed connecting to Redis: %v", err)
    }

    // Get all keys
    keys, err := rdb.Keys(ctx, query).Result()
    if err != nil {
        return nil, fmt.Errorf("failed getting keys from Redis: %v", err)
    } else if len(keys) == 0 {
        return nil, errors.New("key not found")
    }

    result := make(map[string]RedisValue)
    
    // Get all values
    for _, key := range keys {
        value, err := rdb.Get(ctx, key).Result()
        if err != nil {
            continue
        }
        
        var redisValue RedisValue
        if err := json.Unmarshal([]byte(value), &redisValue); err != nil {
            continue
        }
        
        result[key] = redisValue
    }

    return result, nil
}