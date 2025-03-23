package webhook

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gitops-beyond/beyond-sync/internal/ansible"
	"github.com/gitops-beyond/beyond-sync/internal/redis"
)


func Sync() {
	// Create Webhook object
	w := Webhook{}
	// Define a variable to store sha value
	sha := ""
	// Define a variable to use a lock of routine
	var lock sync.Mutex

	// Define a collection of goroutines
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(2)

	// First goroutine
	// Automatic sync
	go func() {
		// End goroutine when function ends
		defer wg.Done()
		// Inifinite loop
		for {
			// Save last sha value
			newSha, err := w.GetLastCommit()
			if err != nil {
				log.Fatal(err)
			}
			
			// Compare sha value from current loop with previous one
			if sha != newSha {
				// Lock the process to not create a conflict between syncs
				lock.Lock()
				// Update sha value with current loop's one
				sha = newSha
				log.Printf("Sync is triggered with new commit hash value of %s", sha)
				// Run Ansible playbook
				ansible.RunAnsibleSync(sha)
				// Unlock the process to use concurrency
				lock.Unlock()
			// Make 30 second pause, if repo has not been updated
			} else {
				log.Println("Sleep")
				time.Sleep(30 * time.Second)
			}
		}
	}()
	
	// Second goroutine
	// Manul sync
	go func() {
		// End goroutine when function ends
		defer wg.Done()
		// Subscribe for Redis Pub/Sub channel
		sub, err := redis.Subscribe()
		if err != nil {
			log.Printf("Redis connection error: %v", err)
			return
		}
		// Close the connection after function ends
		defer sub.Close()

		// Infinite loop of reading messages from Redis channel
		for {
			// Get the message from channel
			msg, err := sub.ReceiveMessage(ctx)
			if err != nil {
				log.Printf("Redis subscription error: %v", err)
				return
			// If the message is there run the sync
			} else if len(msg.Payload) > 0 {
				// Lock the process to not create a conflict between syncs
				lock.Lock()
				log.Printf("Sync is triggered with manual trigger and commit hash value of %s", sha)
				// Run Ansible playbook
				ansible.RunAnsibleSync(sha)
				// Unlock the process to use concurrency
				lock.Unlock()
			}
		}
	}()

	// Block the main goroutine until all goroutines have completed their execution
	wg.Wait()
}