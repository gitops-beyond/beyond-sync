package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gitops-beyond/beyond-sync/api/routes"
)

// StartServer initializes and runs the HTTP server on port 8080
func StartServer() {
	// Create new Gin server with default middleware
	r := gin.Default()
	// Configure API routes
	routes.LoadRoutes(r)
	// Start server
	r.Run(":8080")
}