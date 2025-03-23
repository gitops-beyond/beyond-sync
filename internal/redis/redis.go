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

// Struct definition to set sync parameters to save in Redis
type RedisValue struct {
    Sha string `json:"sha"`
    Status string `json:"status"`
    Message string `json:"message"`
}

// Save sync record to Redis
func AddSyncRecord(sha string, status string, message string) {
    // Create client connection with Redis
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
        Password: "",
        DB:       0,
    })

    ctx := context.Background()
    // Close the connection after function execution
    defer rdb.Close()

    // Check connection with Redis
    if err := rdb.Ping(ctx).Err(); err != nil {
        fmt.Printf("Error connecting to Redis: %v\n", err)
        return
    }

    // Set sync time as key
    key := time.Now().Format("2006-01-02T15:04:05")
    // Put sync parameters to save into variable
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

// Get sync record/records
func GetSyncRecords(query string) (map[string]RedisValue, error) {
    // Create client connection with Redis
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
        Password: "",
        DB:       0,
    })

    ctx := context.Background()
    // Close the connection after function execution
    defer rdb.Close()

    // Check connection first
    if err := rdb.Ping(ctx).Err(); err != nil {
        return nil, fmt.Errorf("failed connecting to Redis: %v", err)
    }

    // Get all keys from Redis
    keys, err := rdb.Keys(ctx, query).Result()
    if err != nil {
        return nil, fmt.Errorf("failed getting keys from Redis: %v", err)
    } else if len(keys) == 0 {
        return nil, errors.New("key not found")
    }

    // Map to store data pulled from Redis
    result := make(map[string]RedisValue)
    
    // Iterate through keys
    for _, key := range keys {
        // Get value from Redis of each key
        value, err := rdb.Get(ctx, key).Result()
        if err != nil {
            continue
        }
        
        // Put value into RedisValue struct
        var redisValue RedisValue
        if err := json.Unmarshal([]byte(value), &redisValue); err != nil {
            continue
        }
        
        // Put the value into map of Redis key-values
        result[key] = redisValue
    }

    return result, nil
}

// Redis message publishing
func PublishMessage() error {
    // Create client connection with Redis
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
        Password: "",
        DB:       0,
    })

    ctx := context.Background()
    // Close the connection after function execution
    defer rdb.Close()

    // Check connection first
    if err := rdb.Ping(ctx).Err(); err != nil {
        return fmt.Errorf("failed connecting to Redis: %v", err)
    }

    // Publish message to "triggers channel"
    err := rdb.Publish(ctx, "triggers", `sync trigger`).Err()
    if err != nil {
        return fmt.Errorf("failed publishing message in Redis: %v", err)
    }

    return nil
}

// Redis channel subscription
func Subscribe() (*redis.PubSub, error) {
    // Create client connection with Redis
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
        Password: "",
        DB:       0,
    })

    ctx := context.Background()
    // We do not close connection here to be left subscribed to channel

    // Check connection first
    if err := rdb.Ping(ctx).Err(); err != nil {
        return nil, fmt.Errorf("failed connecting to Redis: %v", err)
    }

    // Subscribe to channel
    sub := rdb.Subscribe(ctx, "triggers")

    return sub, nil
}