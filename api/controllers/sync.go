package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitops-beyond/beyond-sync/api/models"
	"github.com/gitops-beyond/beyond-sync/internal/redis"
)

func GetSyncs(c *gin.Context) {
    records, err := redis.GetAllSyncRecords()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    response := make(models.SyncListResponse, 0)
    for timestamp, value := range records {
        sync := models.SyncResponse{
            Timestamp: timestamp,
            Data:     value,
        }
        response = append(response, sync)
    }

    c.JSON(200, response)
}