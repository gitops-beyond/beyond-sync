package ansible

import "log"

func Run(){
	err := cloneRepo()
	if err != nil{
		log.Fatalf("ERROR %v", err)
		return
	}

	err = lookForAnsiblePrerequisites()
	if err != nil{
		removeRepo()
		log.Fatalf("ERROR %v", err)
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