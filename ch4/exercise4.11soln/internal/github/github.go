package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	State     string `json:"state"`
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueRequest struct {
	IssueId int
	Title   string `json:"title"`
	Body    string `json:"body"`
}

type IssueClient struct {
	httpClient http.Client
}

func GetIssueClient() *IssueClient {
	return &IssueClient{}
}

func (client *IssueClient) CloseIssue(ownerRepo string, number int) (*Issue, error) {

	payload, err := json.Marshal(Issue{
		State: "closed",
	})

	if err != nil {
		return nil, fmt.Errorf("error marshalling body: %w", err)
	}

	request, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%v/repos/%v/issues/%d", issuesBaseUrl, ownerRepo, number), bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_TOKEN")))
	resp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error closing issue: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("github API %d, %s", resp.StatusCode, body)
	}

	issue := Issue{}
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func (client *IssueClient) ShowIssue(ownerRepo string, number int) (*Issue, error) {

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/repos/%v/issues/%d", issuesBaseUrl, ownerRepo, number), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_TOKEN")))

	resp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error fetching issue: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("github API %d, %s", resp.StatusCode, body)
	}

	issue := Issue{}
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func (client *IssueClient) ListIssues(ownerRepo, state string) ([]Issue, error) {

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/repos/%v/issues", issuesBaseUrl, ownerRepo), nil)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	qParams := request.URL.Query()
	qParams.Add("state", state)
	request.URL.RawQuery = qParams.Encode()

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_TOKEN")))

	resp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error fetch list of issues: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("github API %d, %s", resp.StatusCode, body)
	}

	issues := []Issue{}
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func (client *IssueClient) Create(ownerRepo string, payload *IssueRequest) (*Issue, error) {

	reqBody, err := json.Marshal(*payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/repos/%s/issues", issuesBaseUrl, ownerRepo), bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error when creating request:%w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_TOKEN")))

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error when creating issue:%w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("github API %d, %s", resp.StatusCode, body)
	}

	result := Issue{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *IssueClient) Update(ownerRepo string, payload *IssueRequest) (*Issue, error) {

	reqBody, err := json.Marshal(*payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/repos/%s/issues/%d", issuesBaseUrl, ownerRepo, payload.IssueId), bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error when creating request:%w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_TOKEN")))

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error when updating issue:%w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("github API %d, %s", resp.StatusCode, body)
	}

	result := Issue{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
