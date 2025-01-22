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
		removeRepo()
		log.Printf("ERROR %v", err)
		return
	}

	err = runPlaybook()
	if err != nil{
		removeRepo()
		log.Printf("ERROR %v", err)
		return
	}
}