package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gitops-beyond/beyond-sync/handlers"
)

func main() {
    r := gin.Default()

    r.POST("/api/sync", handlers.TriggerSync)
    r.GET("/api/sync", handlers.GetAllSyncs)
    r.GET("/api/sync/:date", handlers.GetSyncByDate)

    r.Run(":8080")
}