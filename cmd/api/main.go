package main

import (
	"github.com/joho/godotenv"
	"github.com/gitops-beyond/beyond-sync/api"
)

// @title           Beyond Sync API
// @version         1.0
// @description     API for managing sync operations
// @host            localhost:8080
// @BasePath        /
func main() {
	// Load environment variables from .env file
	godotenv.Load("../../.env")
    
	api.StartServer()
}