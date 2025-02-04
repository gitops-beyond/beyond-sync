package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitops-beyond/beyond-sync/internal/redis"
)

type SyncResponse struct {
    Timestamp string     `json:"timestamp"`
    Data      redis.RedisValue `json:"data"`
}

type SyncListResponse []SyncResponse

func GetAllSyncs(c *gin.Context) {
    records, err := redis.GetSyncRecords("*")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    response := make(SyncListResponse, 0)
    for timestamp, value := range records {
        sync := SyncResponse{
            Timestamp: timestamp,
            Data:     value,
        }
        response = append(response, sync)
    }

    c.JSON(200, response)
}

func GetSyncByDate(c *gin.Context) {
    records, err := redis.GetSyncRecords(c.Param("timestamp"))
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    response := make(SyncListResponse, 0)
    for timestamp, value := range records {
        sync := SyncResponse{
            Timestamp: timestamp,
            Data:     value,
        }
        response = append(response, sync)
    }

    c.JSON(200, response)
}