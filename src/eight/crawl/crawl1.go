package crawl

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
)

// tokens限制爬虫并发数量
var tokens = make(chan struct{}, 20)

func Crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
