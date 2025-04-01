package ansible

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"log"
)

// Check if target repo is cloned and Ansible installed on host
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

// Check if all hosts are reachable
func pingAllHosts() error {
	// Go to cloned Ansible code
	os.Chdir("./clonedRepo/ansible")
	// Execute ansible ping command
	pingCmd := exec.Command("ansible", "all", "-i", "inventory", "-m", "ping")
	// Get the command output
	byteOutput, err := pingCmd.CombinedOutput()
	output := string(byteOutput)
	// Go back to home dir
	os.Chdir("../../")
	if err != nil {
		return fmt.Errorf(output)
	}
	if strings.Contains(output, "[WARNING]: No inventory was parsed"){
		return fmt.Errorf("Inventory is not present")
	}
	return nil
}

// Run Ansible playbook
func runPlaybook() (string, error) {
	// Go to cloned Ansible code
	os.Chdir("./clonedRepo/ansible")
	// Execute ansible playbook
	pingCmd := exec.Command("ansible-playbook", "-i", "inventory", "playbook.yml")
	// Get the command output
	byteOutput, err := pingCmd.CombinedOutput()
	log.Println(string(byteOutput))
	// Go back to home dir
	os.Chdir("../../")
	if err != nil {
		return "", err
	}
	return string(byteOutput), nil
}