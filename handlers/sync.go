package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func TriggerSync(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Trigger sync"})

    // Trigger from webhook
    // Run ansible playobook
    // Save sync to history in Redis
}

func GetSyncByDate(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Get sync by date","date": c.Param("date")})
}

func GetAllSyncs(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Get all syncs"})
}
