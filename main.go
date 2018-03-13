package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Repositories struct {
	Total             *int         `json:"total_count,omitempty"`
	IncompleteResults *bool        `json:"incomplete_results,omitempty"`
	Items      []Repo `json:"items,omitempty"`
}

// Repo describes a Github repository
type Repo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stars       int       `json:"stargazers_count"`
	Forks       int       `json:"forks_count"`
	Issues      int       `json:"open_issues_count"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
	URL         string    `json:"html_url"`
}

var result Repositories

func main() {
	apiURL := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
	resp, err := http.Get(apiURL)

	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp.Status)
	}
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&result); err != nil {
		log.Fatal(err)
	}

	save(result.Items)
}

func save(result []Repo) {
	readme, err := os.OpenFile("README.md", os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	readme.WriteString(`# Top Go Projects
A list of most popular github projects related to Go (ranked by stars)

|    | Project Name | Stars | Forks | Open Issues | Description |
| -- | ------------ | ----- | ----- | ----------- | ----------- |
`)
	for i, repo := range result {
		readme.WriteString(fmt.Sprintf("| %d | [%s](%s) | %d | %d | %d | %s |\n", i+1, repo.Name, repo.URL, repo.Stars, repo.Forks, repo.Issues, repo.Description))
	}
	readme.WriteString(fmt.Sprintf("\n*Last Automatic Update: %v*", time.Now().Format("2006-01-02 15:04:05")))
}
