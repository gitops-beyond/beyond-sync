package ansible

import (
	"log"

	"github.com/gitops-beyond/beyond-sync/internal/redis"
)

// RunAnsibleSync orchestrates the ansible playbook execution process
func RunAnsibleSync(sha string) {
	// Clone the repository and ensure cleanup
	err := cloneRepo()
	defer removeRepo()

	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		redis.AddSyncRecord(sha, "Failed", err.Error())
		return
	}

	// Check if all required ansible files exist
	err = lookForAnsiblePrerequisites()
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		redis.AddSyncRecord(sha, "Failed", err.Error())
		return
	}

	// Verify connectivity to all hosts
	err = pingAllHosts()
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		redis.AddSyncRecord(sha, "Failed", err.Error())
		return
	}

	// Execute the ansible playbook
	output, err := runPlaybook()
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		redis.AddSyncRecord(sha, "Failed", err.Error())
		return
	}
	redis.AddSyncRecord(sha, "Synced", output)
}