package ansible

import (
	"log"
)

func Run(sha string){
	err := cloneRepo()
	defer removeRepo()

	if err != nil{
		log.Printf("ERROR: %s", err.Error())
		addSyncRecord(sha, "Failed", err.Error())
		return
	}

	err = lookForAnsiblePrerequisites()
	if err != nil{
		log.Printf("ERROR: %s", err.Error())
		addSyncRecord(sha, "Failed", err.Error())
		return
	}

	err = pingAllHosts()
	if err != nil{
		log.Printf("ERROR: %s", err.Error())
		addSyncRecord(sha, "Failed", err.Error())
		return
	}

	output, err := runPlaybook()
	if err != nil{
		log.Printf("ERROR: %s", err.Error())
		addSyncRecord(sha, "Failed", err.Error())
		return
	}
	addSyncRecord(sha, "Synced", output)
}