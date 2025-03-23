package webhook

import (
	"log"
	"os"
	"reflect"
	"strings"
	"fmt"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// Struct definition to set webook parameters
type Webhook struct {
	RepoName string
	Username string
	Token string
}

// Create Webhook object with field values from env vars
func(w *Webhook) Init(){
	// Use reflection to get the value of the struct
	val := reflect.ValueOf(w).Elem()

	// Iterate over the struct fields
	for i := 0; i < val.NumField(); i++ {
		// Get the field name in uppercase
		field := strings.ToUpper(val.Type().Field(i).Name)
		// Check if env var is set
		if os.Getenv(field) == "" {
			log.Fatal(field, " env var is missing")
		}
		// Set the obect field value
		val.Field(i).SetString(os.Getenv(field))
	}
}

// Create Webhook object with field values from env vars
func (w* Webhook) GetLastCommit() (string, error){
	// Create Webhook object
	w.Init()

	// Create an HTTP client using resty
	client := resty.New()
	// Make the GET request to GitHub API with the token for authentication
	resp, err := client.R().
		SetHeader("Authorization", "token "+w.Token). // Authenticate using PAT
		Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?path=ansible", w.Username, w.RepoName))

	// Handle any errors that might have occurred
	if err != nil {
		log.Fatalf("Error while fetching repo info: %v", err)
	}

	// Check the HTTP status code
	if resp.StatusCode() != 200 {
		log.Fatalf("Error: received status code %d\nResponse: %s", resp.StatusCode(), resp.String())
	}

	// Parse the JSON response
	var commits []map[string]interface{}
	err = json.Unmarshal(resp.Body(), &commits)
	if err != nil {
		log.Fatalf("Error parsing JSON response: %v", err)
	}

	// Find the first `sha` key value
	if len(commits) > 0 {
		if sha, ok := commits[0]["sha"].(string); ok {
			return sha, nil
		} else {
			return "", fmt.Errorf("SHA key not found in the last commit")
		}
	} else {
		return "", fmt.Errorf("No commits found in the response")
	}
}