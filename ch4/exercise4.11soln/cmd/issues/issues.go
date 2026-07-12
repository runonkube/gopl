package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/linehk/gopl/ch4/exercise4.11soln/internal/github"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: issues [subcommand] [options]\n")
		os.Exit(1)
	}

	switch args[1] {
	case "create":
		if issue, err := createIssue(args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating issue: %s", err)
			os.Exit(1)
		} else {
			fmt.Printf("Issue created.\nNumber: %d\nState: %s\nTitle: %s\n", issue.Number, issue.State, issue.Title)
		}
	case "update":
		if issue, err := updateIssue(args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error updating issue: %s", err)
			os.Exit(1)
		} else {
			fmt.Printf("Issue updated.\nNumber: %d\nState: %s\nTitle: %s\n", issue.Number, issue.State, issue.Title)
		}
	case "list":
		if issues, err := listIssues(args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		} else {
			fmt.Println("Issue Number\tState\tTitle")
			for _, issue := range issues {
				fmt.Printf("%-12d\t%-5s\t%s\n", issue.Number, issue.State, issue.Title)
			}
		}
	case "show":
		if issue, err := showIssue(args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		} else {
			fmt.Printf("Title: %s\nState: %v\n%s\n\n", issue.Title, issue.State, issue.Body)
		}
	case "close":
		if issue, err := closeIssue(args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		} else {
			if strings.ToLower(issue.State) == "closed" {
				fmt.Println("Issue closed")
			} else {
				fmt.Println("Issue not closed")
			}
		}
	}

}

func closeIssue(args []string) (*github.Issue, error) {
	usage := errors.New("Usage: issues close [options] <owner>/<repo>")
	if len(args) < 1 {
		return nil, usage
	}

	closeFlagSet := flag.NewFlagSet("close", flag.ExitOnError)
	issueNumber := closeFlagSet.Int("number", 0, "The issue number")

	if err := closeFlagSet.Parse(args); err != nil {
		return nil, err
	}

	remainingArgs := closeFlagSet.Args()
	if len(remainingArgs) == 0 {
		return nil, usage
	}

	ownerRepo := remainingArgs[0]

	return github.GetIssueClient().CloseIssue(ownerRepo, *issueNumber)
}

func showIssue(args []string) (*github.Issue, error) {
	usage := errors.New("Usage: issues show [options] <owner>/<repo>")
	if len(args) < 1 {
		return nil, usage
	}

	showFlagSet := flag.NewFlagSet("show", flag.ExitOnError)
	issueNumber := showFlagSet.Int("number", 0, "The issue number")

	if err := showFlagSet.Parse(args); err != nil {
		return nil, err
	}

	remainingArgs := showFlagSet.Args()
	if len(remainingArgs) == 0 {
		return nil, usage
	}

	ownerRepo := remainingArgs[0]

	return github.GetIssueClient().ShowIssue(ownerRepo, *issueNumber)
}

func listIssues(args []string) ([]github.Issue, error) {
	usage := errors.New("Usage: issues list [options] <owner>/<repo>")
	if len(args) < 1 {
		return nil, usage
	}

	listFlagSet := flag.NewFlagSet("list", flag.ExitOnError)
	state := listFlagSet.String("state", "", "The state of the issues to list i.e., open|closed|all")

	if err := listFlagSet.Parse(args); err != nil {
		return nil, err
	}

	remainingArgs := listFlagSet.Args()
	if len(remainingArgs) == 0 {
		return nil, usage
	}

	ownerRepo := remainingArgs[0]

	return github.GetIssueClient().ListIssues(ownerRepo, *state)
}

func createIssue(args []string) (*github.Issue, error) {
	usage := errors.New("Usage: issues create [options] <owner>/<repo>")

	if len(args) < 1 {
		return nil, usage
	}

	createFlagSet := flag.NewFlagSet("create", flag.ExitOnError)
	title := createFlagSet.String("title", "", "Name of github issue")
	body := createFlagSet.String("body", "", "Body of github issue. Note: #lines starting with # will be ignored.")

	// 1. Parse flags first!
	if err := createFlagSet.Parse(args); err != nil {
		return nil, err
	}

	// 2. Grab the remaining non-flag arguments left over
	remainingArgs := createFlagSet.Args()
	if len(remainingArgs) < 1 {
		return nil, usage
	}

	repoTarget := remainingArgs[0] // This will be "owner/repo"
	bodyToUse := *body
	titleToUse := *title

	if strings.TrimSpace(bodyToUse) == "" {
		template := getTemplate(*title)
		titleFromEditor, bodyFromEditor, err := captureFromEditor(template)

		if err != nil {
			return nil, err
		}
		bodyToUse = bodyFromEditor
		titleToUse = titleFromEditor
	}

	if strings.TrimSpace(bodyToUse) == "" || strings.TrimSpace(titleToUse) == "" {
		return nil, errors.New("Title and Body cannot be empty")
	}

	payload := github.IssueRequest{
		Title: titleToUse,
		Body:  bodyToUse,
	}

	return github.GetIssueClient().Create(repoTarget, &payload)
}

func updateIssue(args []string) (*github.Issue, error) {
	usage := errors.New("Usage: issues update [options] <owner>/<repo>")

	if len(args) < 1 {
		return nil, usage
	}

	updateFlagSet := flag.NewFlagSet("update", flag.ExitOnError)
	issueNumber := updateFlagSet.Int("number", 0, "The issue number")
	title := updateFlagSet.String("title", "", "Name of github issue")
	body := updateFlagSet.String("body", "", "Body of github issue. Note: #lines starting with # will be ignored.")

	// 1. Parse flags first!
	if err := updateFlagSet.Parse(args); err != nil {
		return nil, err
	}

	// 2. Grab the remaining non-flag arguments left over
	remainingArgs := updateFlagSet.Args()
	if len(remainingArgs) < 1 {
		return nil, usage
	}

	repoTarget := remainingArgs[0] // This will be "owner/repo"
	bodyToUse := *body
	titleToUse := *title

	if strings.TrimSpace(bodyToUse) == "" {
		template := getTemplate(*title)
		titleFromEditor, bodyFromEditor, err := captureFromEditor(template)

		if err != nil {
			return nil, err
		}
		bodyToUse = bodyFromEditor
		titleToUse = titleFromEditor
	}

	if strings.TrimSpace(bodyToUse) == "" || strings.TrimSpace(titleToUse) == "" {
		return nil, errors.New("Title and Body cannot be empty")
	}

	payload := github.IssueRequest{
		IssueId: *issueNumber,
		Title:   titleToUse,
		Body:    bodyToUse,
	}

	return github.GetIssueClient().Update(repoTarget, &payload)
}

func getTemplate(title string) string {
	return title + "\n" +
		"#Please enter the issue title on the first line." + "\n" +
		"# Lines starting with '#' will be ignored." + "\n" +
		"# An empty file aborts the issue."
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
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr //re-assign the std input,err and output ot the command we're running

	//run command to open editor and wait for it to complete
	if err := cmd.Run(); err != nil {
		return "", "", err
	}

	raw, err := os.ReadFile(tmp.Name())
	if err != nil {
		return "", "", err
	}

	return parseTitleAndBody(string(raw)) // strip # lines, split title/body
}

func pickEditor() string {
	//Check $VISUAL first — historically reserved for full-screen editors (vim, emacs, nano)
	// If $VISUAL is empty then check $EDITOR next - historically for line editors but in practice what most people set
	//If both $VISUAL and $EDITOR are empty, then a sensible default should be vi on Unix (POSIX guarantees it exists), notepad on Windows
	if e := os.Getenv("VISUAL"); e != "" {
		return e
	}
	if e := os.Getenv("EDITOR"); e != "" {
		return e
	}
	if runtime.GOOS == "windows" {
		return "notepad"
	}
	return "vi"
}

func parseTitleAndBody(raw string) (title, body string, err error) {
	var content []string
	for _, line := range strings.Split(raw, "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		content = append(content, line)
	}

	joined := strings.TrimSpace(strings.Join(content, "\n"))
	if joined == "" {
		return "", "", fmt.Errorf("empty issue: aborting")
	}

	parts := strings.SplitN(joined, "\n", 2)
	title = parts[0]
	if len(parts) > 1 {
		body = strings.TrimSpace(parts[1])
	}
	return title, body, nil
}
