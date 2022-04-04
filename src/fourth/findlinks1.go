package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

//var count map[string]int

func visit(links []string, n *html.Node) []string {
	if n.Type == html.TextNode && n.Data != "script" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(links, c)
	}
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(nil, doc)
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
