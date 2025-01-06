package main

import (
	"github.com/joho/godotenv"
	//"github.com/gitops-beyond/beyond-sync/webhook"
	"github.com/gitops-beyond/beyond-sync/ansible"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()
    
	//webhook.Sync()
	ansible.Run()
}