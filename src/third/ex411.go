package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IssueParams struct {
	Owner       string
	Repo        string
	IssueNumber string
	Token       string
	Issue       *IssueContent
}

type IssueContent struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var BaseUrl string = "https://api.github.com/repos/"

func (p IssueParams) GetIssue(params *IssueParams) (*IssueContent, error) {
	url := BaseUrl + params.Owner + "/" + params.Repo + "/issues" + "/" + params.IssueNumber
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Get issue failed: %d", resp.StatusCode)
	}

	var content *IssueContent
	if err := json.NewDecoder(resp.Body).Decode(&content); err != nil {
		return nil, err
	}
	return content, nil
}

func (p IssueParams) CrateIssue() bool {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(p.Issue); err != nil {
		return false
	}
	u := BaseUrl + p.Owner + "/" + p.Repo + "/issues" + "/" + p.IssueNumber + "?access_token=" + p.Token
	request, err := http.NewRequest(http.MethodPatch, u, &buf)
	if err != nil {
		return false
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return false
	}
	request.Body.Close()
	return true
}
