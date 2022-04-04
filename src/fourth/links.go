package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var base = flag.String("base", "https://www.sulinehk.com", "base url to crawl")

// Extract parse url to return url slice which included in html
func Extract(url string) ([]string, error) {
	resp, err := parse(url)

	// download file use copy and html.Parse both will read util EOF of file so need to get url twice
	download(*base, url)
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(attr.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	// depthFirst to parse html
	forEachNodes(doc, visitNode, nil)
	//fmt.Println(links)
	return links, nil
}

func forEachNodes(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNodes(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// f extract url html to return url included in html such as func crawl
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func parse(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s : %s", url, resp.Status)
	}
	return resp, nil
}

func download(base, url string) {
	resp, err := parse(url)

	// html file parse will parse until to EOF
	if !strings.HasPrefix(url, base) {
		return
	}

	dir := strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://")
	if !Exists(dir) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatalln(err)
		}
	}
	filename := dir + "index.html"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	defer resp.Body.Close()
	n, err := io.Copy(f, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d bytes write to file\n", n)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//func main() {
//	flag.Parse()
//	breadthFirst(crawl, []string{*base})
//}
