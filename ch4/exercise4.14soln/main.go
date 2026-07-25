package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/runonkube/issues-cli/pkg/github"
)

var (
	client          = github.NewIssueClient(nil)
	displayTemplate = template.Must(template.New("issuelist").Parse(`
<h1>{{. | len}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>`))
)

func main() {
	http.HandleFunc("/issues", GetIssues)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func GetIssues(responseWriter http.ResponseWriter, request *http.Request) {

	queryParams := request.URL.Query()
	repo := queryParams.Get("repo")
	state := queryParams.Get("state")

	if repo == "" || state == "" {
		fmt.Fprintln(responseWriter, "Must specify repo and state e.g. 'http://localhost:8080/issues?repo=runonkube/issues-cli&state=open'")
		return
	}

	issues, err := client.ListIssues(repo, state)

	if err != nil {
		fmt.Fprintf(responseWriter, "Error: %s", err)
		return
	}

	responseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := displayTemplate.Execute(responseWriter, issues); err != nil {
		fmt.Fprintf(responseWriter, "Error: %s", err)
	}

}
