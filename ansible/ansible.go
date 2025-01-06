package ansible

import (
	"fmt"
	"os"
	"os/exec"
)

func lookForAnsiblePrerequisites() error {
	// Ansible folder in repo
	_, err := os.ReadDir("./clonedRepo/ansible")
	if err != nil {
		return fmt.Errorf("ansible directory is not present")
	}
	// Ansible CLI installed locally
	_, err = exec.LookPath("ansible")
	if err != nil {
		return fmt.Errorf("ansible cli is not present")
	}
	return nil
}