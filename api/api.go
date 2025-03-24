package api

import (
    "github.com/gin-gonic/gin"
    "github.com/gitops-beyond/beyond-sync/api/routes"
    _ "github.com/gitops-beyond/beyond-sync/docs"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Beyond Sync API
// @version         1.0
// @description     API for managing sync operations
// @host            localhost:8080
// @BasePath        /

// StartServer initializes and runs the HTTP server on port 8080
func StartServer() {
    // Create new Gin server with default middleware
    r := gin.Default()

    // Swagger documentation endpoint
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Static file server for swagger.json
    r.Static("/docs", "./docs")

    // Configure API routes
    routes.LoadRoutes(r)

    // Start server
    r.Run(":8080")
}