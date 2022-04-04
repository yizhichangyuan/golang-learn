package main

import (
	"html/template"
	"learn/git"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}} issues
{{range .Items}}-------------------
Number: {{.Number}}
User: {{.User}}
Title: {{.Title}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`

const issueHtml = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func onMust() {
	report := template.Must(template.New("reportHtml").Parse(issueHtml))
	// report := template.Must(
	// template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).
	//	Parse(templ))
	result, err := git.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func main() {
	onMust()
}
