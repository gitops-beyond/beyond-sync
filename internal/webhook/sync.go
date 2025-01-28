package webhook

import (
	"log"
	"time"
	"github.com/gitops-beyond/beyond-sync/internal/ansible"
)

func Sync() {
	w := Webhook{}
	sha := ""
	for {
		newSha, err := w.GetLastCommit()
		if err != nil {
			log.Fatal(err)
		}

		if sha != newSha {
			sha = newSha
			log.Printf("Sync is triggered with new commit hash value of %s", sha)
			ansible.Run(sha)
		} else {
			log.Println("Sleep")
			time.Sleep(30 * time.Second)
		}
	}
}