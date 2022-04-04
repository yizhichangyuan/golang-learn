package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

var scan *bufio.Scanner = bufio.NewScanner(os.Stdin)
var images int
var words int

func init() {
	scan.Split(bufio.ScanWords)
}

func CountWordsAndImages(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	if resp.StatusCode != 200 {
		return err
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	countWordsAndImage(doc)
	return nil
}

func countWordsAndImage(n *html.Node) {
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countWordsAndImage(c)
	}
}

func main() {
	err := CountWordsAndImages("https://mail.163.com/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("words:%d\timages:%d", words, images)
}
