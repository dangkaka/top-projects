package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type Repositories struct {
	Total             *int   `json:"total_count,omitempty"`
	IncompleteResults *bool  `json:"incomplete_results,omitempty"`
	Items             []Repo `json:"items,omitempty"`
}

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
	now := time.Now()
	backup := "backup/backup_" + now.Format("20060102") + ".md"
	exec.Command("mv", "README.md", backup).Run()
	save(result.Items, now)
}

func save(result []Repo, now time.Time) {
	readme, err := os.OpenFile("README.md", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	readme.WriteString(fmt.Sprintf("*Updated automatically at: %v* \n", now.Format(time.RFC3339)))
	readme.WriteString(`# Top Go Projects
A list of most popular github projects in Go (by stars)

|    | Project Name | Stars | Forks | Open Issues | Description |
| -- | ------------ | ----- | ----- | ----------- | ----------- |
`)
	for i, repo := range result {
		readme.WriteString(fmt.Sprintf("| %d | [%s](%s) | %d | %d | %d | %s |\n", i+1, repo.Name, repo.URL, repo.Stars, repo.Forks, repo.Issues, repo.Description))
	}
}
