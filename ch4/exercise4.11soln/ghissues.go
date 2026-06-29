package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/linehk/gopl/ch4/github"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: issues [subcommand] [options]")
		os.Exit(1)
	}

	switch args[1] {
	case "create":
		if issue, err := createIssue(args[2:]); err != nil {
			fmt.Printf("Error creating issue: %s", err)
		} else {
			fmt.Printf("Issue created: %v", *issue)
		}
	case "close":
	case "list":
	case "read":
	case "update":
	}

}

func createIssue(args []string) (*github.Issue, error) {

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: ghissues create GITHUB URL [options]")
		os.Exit(1)
	}

	createFlags := flag.NewFlagSet("create", flag.ContinueOnError)
	title := createFlags.String("title", "issue 1", "Name of github issue")
	if err := createFlags.Parse(args[1:]); err != nil {
		return nil, err
	}

	payload := github.IssueCreateRequest{
		Title: *title,
		Body:  "Everything is still broken, nothing is still working",
	}

	client := github.IssueClient{}
	return client.Create(args[0], &payload)
}

func captureFromEditor(template string) (title, body string, err error) {
	tmp, err := os.CreateTemp("", "issue-*.md") //create temp file
	if err != nil {
		return "", "", err
	}
	defer os.Remove(tmp.Name()) //clean up

	//write to the temp file
	if _, err := tmp.WriteString(template); err != nil {
		return "", "", err
	}
	tmp.Close() //close it

	editor := pickEditor() // VISUAL → EDITOR → default
	parts := strings.Fields(editor)
	cmd := exec.Command(parts[0], append(parts[1:], tmp.Name())...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return "", "", err
	}

	raw, err := os.ReadFile(tmp.Name())
	if err != nil {
		return "", "", err
	}

	return parseTitleAndBody(string(raw)) // strip # lines, split title/body
}
