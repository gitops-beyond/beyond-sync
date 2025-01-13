package ansible

import "log"

func Run(){
	err := cloneRepo()
	if err != nil{
		log.Printf("ERROR %v", err)
		addNonAnsibleErrorRecord()
		return
	}

	err = lookForAnsiblePrerequisites()
	if err != nil{
		removeRepo()
		log.Printf("ERROR %v", err)
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

	removeRepo()
}