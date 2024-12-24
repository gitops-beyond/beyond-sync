package main

import (
	"github.com/joho/godotenv"
	"github.com/gitops-beyond/beyond-sync/webhook"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()
    
	webhook.Sync()
}