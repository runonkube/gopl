package main

import (
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
		fmt.Fprintf(os.Stderr, "Usage: issues [subcommand] [options]")
		os.Exit(1)
	}

	switch args[1] {
	case "create":
		if issue, err := createIssue(args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating issue: %s", err)
		} else {
			fmt.Printf("Issue created: %v", *issue)
		}
	case "close":
	case "list":
		if err := listIssues(args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "")
		}
	case "get":
	case "update":
	}

}

func listIssues(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: issues list [options] <owner>/<repo>")
	}

}

func createIssue(args []string) (*github.Issue, error) {

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: issues create [options] <owner>/<repo>")
		os.Exit(1)
	}

	createFlags := flag.NewFlagSet("create", flag.ExitOnError)
	title := createFlags.String("title", "", "Name of github issue")
	body := createFlags.String("body", "", "Body of github issue. Note: #lines starting with # will be ignored.")

	// 1. Parse flags first!
	if err := createFlags.Parse(args); err != nil {
		return nil, err
	}

	// 2. Grab the remaining non-flag arguments left over
	remainingArgs := createFlags.Args()
	if len(remainingArgs) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: issues create [options] <owner>/<repo>")
		os.Exit(1)
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
		return nil, fmt.Errorf("Title and Body cannot be empty")
	}

	payload := github.IssueCreateRequest{
		Title: titleToUse,
		Body:  bodyToUse,
	}

	return github.GetIssueClient().Create(repoTarget, &payload)
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
