package ansible

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"log"
)

func lookForAnsiblePrerequisites() error {
	// Ansible folder in repo
	_, err := os.ReadDir("./clonedRepo/ansible")
	if err != nil {
		return fmt.Errorf("Ansible directory is not present")
	}
	// Ansible CLI installed locally
	_, err = exec.LookPath("ansible")
	if err != nil {
		return fmt.Errorf("Ansible CLI is not present")
	}
	return nil
}

func pingAllHosts() error {
	os.Chdir("./clonedRepo/ansible")
	pingCmd := exec.Command("ansible", "all", "-i", "inventory", "-m", "ping")
	byteOutput, err := pingCmd.CombinedOutput()
	output := string(byteOutput)
	os.Chdir("../../")
	if err != nil {
		return fmt.Errorf(output)
	}
	if strings.Contains(output, "[WARNING]: No inventory was parsed"){
		return fmt.Errorf("Inventory is not present")
	}
	return nil
}

func runPlaybook() (string, error) {
	os.Chdir("./clonedRepo/ansible")
	pingCmd := exec.Command("ansible-playbook", "-i", "inventory", "playbook.yml")
	byteOutput, err := pingCmd.CombinedOutput()
	log.Println(string(byteOutput))
	os.Chdir("../../")
	if err != nil {
		return "", err
	}
	return string(byteOutput), nil
}