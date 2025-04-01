package main

import (
	"github.com/gitops-beyond/beyond-sync/api"
)

// @title           Beyond Sync API
// @version         1.0
// @description     API for managing sync operations
// @host            localhost:8080
// @BasePath        /
func main() {
	api.StartServer()
}