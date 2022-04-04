package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

var found bool
var id string

func ElementByID(doc *html.Node) *html.Node {
	var depth int
	pre := func(n *html.Node, haveNoChild bool) bool {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s", depth*2, " ", n.Data)
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					found = true
				}
				fmt.Printf(" %s='%s'", attr.Key, attr.Val)
			}

			if haveNoChild {
				fmt.Print("/")
			}
			fmt.Println(">")
			depth++
		}
		return found
	}

	post := func(n *html.Node, haveNoChild bool) bool {
		if n.Type == html.ElementNode {
			depth--
			if !haveNoChild {
				fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
			}
		}
		return found
	}
	return forEachNode(doc, pre, post)
}

func forEachNode(n *html.Node, pre, post func(*html.Node, bool) bool) *html.Node {
	flag := n.FirstChild == nil
	if pre != nil {
		if pre(n, flag) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		doc := forEachNode(c, pre, post)
		if doc != nil {
			return doc
		}
	}
	if post != nil {
		post(n, flag)
	}
	return nil
}

func main() {
	resp, err := http.Get(os.Args[1])
	id = os.Args[2]
	if err != nil {
		fmt.Errorf("visit url: %s failed", err)
		os.Exit(-1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("visit url code: %d", resp.StatusCode)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Errorf("parse url %s as html failed: %v", os.Args[1], err)
		os.Exit(-1)
	}
	//forEachNode(doc, pre, post)
	n := ElementByID(doc)
	resp.Body.Close()
	fmt.Printf("<%s>", n.Data)
}
