package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitops-beyond/beyond-sync/internal/redis"
)

// SyncResponse represents a single sync operation result
type SyncResponse struct {
    Timestamp string          `json:"timestamp"`
    Data      redis.RedisValue `json:"data"`
}

// SyncListResponse is a collection of sync responses
type SyncListResponse []SyncResponse

// GetAllSyncs godoc
// @Summary      Get all sync records
// @Description  Retrieves all sync records from Redis
// @Tags         sync
// @Accept       json
// @Produce      json
// @Success      200  {array}   SyncResponse
// @Router       /sync [get]
func GetAllSyncs(c *gin.Context) {
    redisRecords, err := redis.GetSyncRecords("*")
    if err != nil && err.Error() == "key not found"{
        c.JSON(404, gin.H{"error": err.Error()})
        return
    } else if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // Convert Redis records to response format
    response := make(SyncListResponse, 0)
    for timestamp, value := range redisRecords {
        sync := SyncResponse{
            Timestamp: timestamp,
            Data:     value,
        }
        response = append(response, sync)
    }

    c.JSON(200, response)
}

// GetSyncByDate godoc
// @Summary      Get sync record by timestamp
// @Description  Retrieves a specific sync record by its timestamp
// @Tags         sync
// @Accept       json
// @Produce      json
// @Param        timestamp   path      string  true  "Timestamp of the sync record"
// @Success      200  {object}  SyncResponse
// @Router       /sync/{timestamp} [get]
func GetSyncByDate(c *gin.Context) {
    redisKey := c.Param("timestamp")
    redisValue, err := redis.GetSyncRecords(redisKey)
    if err != nil && err.Error() == "key not found" {
        c.JSON(404, gin.H{"error": err.Error()})
        return
    } else if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    response := SyncResponse{
        Timestamp: redisKey,
        Data: redisValue[redisKey],
    }

    c.JSON(200, response)
}

// TriggerSync godoc
// @Summary      Trigger new sync operation
// @Description  Triggers a new synchronization operation
// @Tags         sync
// @Accept       json
// @Produce      json
// @Success      201  {string}  string    "Sync trigger is requested"
// @Router       /sync/trigger [post]
func TriggerSync(c *gin.Context) {
    err := redis.PublishMessage()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, "Sync trigger is requested")
}