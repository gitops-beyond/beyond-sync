package ansible

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// cloneRepo clones a GitHub repo using env vars for auth
func cloneRepo() error {
	// Build repo URL from env vars
	repoUrl := fmt.Sprintf("https://github.com/%s/%s", os.Getenv("USERNAME"), os.Getenv("REPONAME"))
	log.Printf("Cloning repo %s", repoUrl)
	
	// Clone the repo
	_, err := git.PlainClone("./clonedRepo/", false, &git.CloneOptions{
		URL: repoUrl,
		Auth: &http.BasicAuth{
			Username:  os.Getenv("USERNAME"),
			Password: os.Getenv("TOKEN"),
		},
		Depth: 1, // Only get latest commit
	})

	if err != nil {
		return fmt.Errorf("Failed to clone repo %s: %s", repoUrl, err)
	}
	return nil
}

// removeRepo deletes the cloned repo directory
func removeRepo() error {
	err := os.RemoveAll("./clonedRepo")
	if err != nil {
		return fmt.Errorf("Failed to remove repo: %v", err)
	}
	return nil
}