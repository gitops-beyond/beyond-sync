package ansible

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func pingAllHosts() error {
	os.Chdir("./clonedRepo/ansible")
	pingCmd := exec.Command("ansible", "all", "-m", "ping")
	byteOutput, err := pingCmd.CombinedOutput()
	os.Chdir("../../")
	if err != nil {
		return err
	}
	output := string(byteOutput)
	if strings.Contains(output, "[WARNING]: No inventory was parsed"){
		return fmt.Errorf("Inventory is not present")
	}
	return nil
}