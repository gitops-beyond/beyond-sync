package main

import (
	"github.com/joho/godotenv"
	"github.com/gitops-beyond/beyond-sync/internal/webhook"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load("../../.env")
    
	webhook.Sync()
}