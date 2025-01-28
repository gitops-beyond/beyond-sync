package ansible

import (
	"log"

	"github.com/gitops-beyond/beyond-sync/internal/redis"
)

func RunAnsibleSync(sha string){
	err := cloneRepo()
	defer removeRepo()

	if err != nil{
		log.Printf("ERROR: %s", err.Error())
		redis.AddSyncRecord(sha, "Failed", err.Error())
		return
	}

	err = lookForAnsiblePrerequisites()
	if err != nil{
		log.Printf("ERROR: %s", err.Error())
		redis.AddSyncRecord(sha, "Failed", err.Error())
		return
	}

	err = pingAllHosts()
	if err != nil{
		log.Printf("ERROR: %s", err.Error())
		redis.AddSyncRecord(sha, "Failed", err.Error())
		return
	}

	output, err := runPlaybook()
	if err != nil{
		log.Printf("ERROR: %s", err.Error())
		redis.AddSyncRecord(sha, "Failed", err.Error())
		return
	}
	redis.AddSyncRecord(sha, "Synced", output)
}