package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gitops-beyond/beyond-sync/handlers"
    "github.com/joho/godotenv"
	"os"
    "log"
)

func main() {
    // Load environment variables from .env file
	godotenv.Load()
    if os.Getenv("PORT") == "" {
        log.Fatal("PORT env variable missing")
    }

    r := gin.Default()

    r.POST("/api/sync", handlers.TriggerSync)
    r.GET("/api/sync", handlers.GetAllSyncs)
    r.GET("/api/sync/:date", handlers.GetSyncByDate)

    r.Run(":" + os.Getenv("PORT"))
}