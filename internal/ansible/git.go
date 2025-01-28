package ansible

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func cloneRepo() error{
	repoUrl := fmt.Sprintf("https://github.com/%s/%s", os.Getenv("USERNAME"), os.Getenv("REPONAME"))
	log.Printf("Cloning repo %s", repoUrl)
	
	_, err := git.PlainClone("./clonedRepo/", false, &git.CloneOptions{
		URL: repoUrl,
		Auth: &http.BasicAuth{
			Username:  os.Getenv("USERNAME"),
			Password: os.Getenv("TOKEN"),
		},
		Depth: 1,
	})

	if err != nil {
		return fmt.Errorf("Failed to clone repo %s: %s", repoUrl, err)
	}
	return nil
}

func removeRepo() error{
	err := os.RemoveAll("./clonedRepo")
	if err != nil {
		return fmt.Errorf("Failed to remove repo: %v", err)
	}
	return nil
}