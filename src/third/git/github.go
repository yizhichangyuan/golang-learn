package git

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	Body     string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

const IssuesURL = "https://api.github.com/search/issues"

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func HandleResult(result IssueSearchResult) map[string][]Issue {
	ClassifyResult := map[string][]Issue{}
	for _, v := range result.Items {
		y, m, _ := time.Now().Date()
		iy, im, _ := v.CreateAt.Date()
		switch {
		case m-im <= time.Month(1):
			ClassifyResult["oneMonth"] = append(ClassifyResult["oneMonth"], *v)
		case m-im > time.Month(1):
			ClassifyResult["oneYear"] = append(ClassifyResult["oneYear"], *v)
		case y-iy <= 1:
			ClassifyResult["oneYear"] = append(ClassifyResult["oneYear"], *v)
		case y-iy > 1:
			ClassifyResult["moreThanYear"] = append(ClassifyResult["moreThanYear"], *v)
		}
	}
	return ClassifyResult
}

func main() {
	result, err := SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	classifyMap := HandleResult(*result)
	for _, item := range classifyMap {
		fmt.Printf("%v\n",
			item)
	}
}
