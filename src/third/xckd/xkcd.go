package xckd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	MinNum = 1
	MaxNum = 200
)

type Carton struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

type Index struct {
	Comics []*Carton
}

func New() Index {
	return Index{[]*Carton{}}
}

var baseUrl string = "https://xkcd.com/%d/info.0.json"

func Get(num int) (*Carton, error) {
	url := fmt.Sprintf(baseUrl, num)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Get an carton picture failed: %d\n", num)
	}

	var ct *Carton
	if err := json.NewDecoder(resp.Body).Decode(&ct); err != nil {
		return nil, err
	}
	resp.Body.Close()
	return ct, nil
}

func Filter(filters []string, cartons []*Carton) []*Carton {
	var result []*Carton
	for _, v := range cartons {
		if isMatch(filters, *v) {
			result = append(result, v)
		}
	}
	return result
}

func isMatch(filters []string, carton Carton) bool {
	for _, filter := range filters {
		if strings.Contains(carton.Title, filter) {
			return true
		}
	}
	return false
}
