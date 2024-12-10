package main

import (
    "os"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "github.com/gitops-beyond/beyond-sync/handlers"
    "github.com/gitops-beyond/beyond-sync/webhook"
)

func main() {
    // Load environment variables from .env file
	godotenv.Load()
    if os.Getenv("PORT") == "" {
        log.Fatal("PORT env variable missing")
    }

    w := &webhook.Webhook{}
    w.TestAuth()

    r := gin.Default()

    r.POST("/api/sync", handlers.TriggerSync)
    r.GET("/api/sync", handlers.GetAllSyncs)
    r.GET("/api/sync/:date", handlers.GetSyncByDate)

    r.Run(":" + os.Getenv("PORT"))
}