package main

import (
	"learn/src/eight/crawl"
	"os"
)

// 限制爬虫数量，当worklist为空或者没有crawl的goroutines在运行时退出循环
func main() {
	worklist := make(chan []string)
	seen := make(map[string]bool)

	// n表示worklist中还有多少可爬取的链接，当n为0的时候表示没有链接了就终止
	var n int
	n++

	// 无缓存channel发送和接收必须不能在同一个goroutines中，否则会阻塞其中一方
	go func() {
		worklist <- os.Args[1:]
	}()

	for ; n > 0; n-- {
		// 取出一个使用完后，就n--，当n为0终止
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				// 每个url的爬取采用并发，将worklist当做一个队列使用
				// 表明有一个新链接可爬虫
				n++
				go func(url string) {
					worklist <- crawl.Crawl(link)
				}(link)
			}
		}
	}
}
