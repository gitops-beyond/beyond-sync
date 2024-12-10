package webhook

import (
	"log"
	"os"
	"reflect"
	"strings"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Webhook struct {
	Repo string
	Username string
	Token string
}

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

func(w *Webhook) TestAuth(){
	w.Init()

	// Create an HTTP client using resty
	client := resty.New()
	// Make the GET request to GitHub API with the token for authentication
	resp, err := client.R().
		SetHeader("Authorization", "token "+w.Token). // Authenticate using PAT
		Get(fmt.Sprintf("https://api.github.com/repos/%s/%s", w.Username, w.Repo))

	// Handle any errors that might have occurred
	if err != nil {
		log.Fatalf("Error while fetching repo info: %v", err)
	}

	// Print out the response body (repository details)
	fmt.Printf("Response: %s\n", resp.String())
}