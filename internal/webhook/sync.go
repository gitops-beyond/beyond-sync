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
	w := Webhook{}
	sha := ""
	var lock sync.Mutex

	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			newSha, err := w.GetLastCommit()
			if err != nil {
				log.Fatal(err)
			}

			if sha != newSha {
				lock.Lock()
				sha = newSha
				log.Printf("Sync is triggered with new commit hash value of %s", sha)
				ansible.RunAnsibleSync(sha)
				lock.Unlock()
			} else {
				log.Println("Sleep")
				time.Sleep(30 * time.Second)
			}
		}
	}()

	go func() {
		defer wg.Done()
		sub, err := redis.Subscribe()
		if err != nil {
			log.Printf("Redis connection error: %v", err)
			return
		}
		defer sub.Close()

		for {
			msg, err := sub.ReceiveMessage(ctx)
			if err != nil {
				log.Printf("Redis subscription error: %v", err)
				return
			} else if len(msg.Payload) > 0 {
				lock.Lock()
				log.Printf("Sync is triggered with manual trigger and commit hash value of %s", sha)
				ansible.RunAnsibleSync(sha)
				lock.Unlock()
			}
		}
	}()

	wg.Wait()
}