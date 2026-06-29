// Package github provides a GO API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.

package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const issuesBaseUrl = "https://api.github.com"

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueCreateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type IssueClient struct {
	http http.Client
}

func (client *IssueClient) Create(ownerRepo string, payload *IssueCreateRequest) (*Issue, error) {

	reqBody, err := json.Marshal(*payload)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling request payload: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/repos/%s/issues", issuesBaseUrl, ownerRepo), bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("Error when creating request:%s", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_TOKEN")))

	resp, err := client.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error when creating issue:%s", err)
	}
	defer resp.Body.Close()

	result := Issue{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &result, nil
}
