package ansible

import "log"

func Run(){
	err := cloneRepo()
	if err != nil{
		log.Fatalf("ERROR %v", err)
	}

	err = lookForAnsiblePrerequisites()
	if err != nil{
		removeRepo()
		log.Fatalf("ERROR %v", err)
	}

	err = pingAllHosts()
	if err != nil{
		removeRepo()
		log.Printf("ERROR %v", err)
	}

	err = runPlaybook()
	if err != nil{
		removeRepo()
		log.Printf("ERROR %v", err)
	}

	removeRepo()
}