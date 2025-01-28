package main

import (
	"github.com/joho/godotenv"
	"github.com/gitops-beyond/beyond-sync/api"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load("../../.env")
    
	api.StartServer()
}