package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type work struct {
	link  string
	depth int
}

func main() {

	var conMap sync.Map
	worklist := make(chan []work)

	go func() {
		var workItems []work
		for _, url := range os.Args[1:] {
			workItems = append(workItems, work{url, 1})
		}
		worklist <- workItems
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for works := range worklist {
				for _, workItem := range works {
					if _, ok := conMap.Load(workItem.link); !ok {
						conMap.Store(workItem.link, true)
						go func() {
							download(workItem.link)
						}()

						if workItem.depth <= 2 {
							worklist <- crawl2(workItem)
						}
					}
				}
			}
		}()
	}
}

func crawl2(w work) []work {
	fmt.Printf("depth: %d, url: %s\n", w.depth, w.link)

	urls, err := links.Extract(w.link)
	if err != nil {
		log.Print(err)
	}

	var works []work
	for _, url := range urls {
		works = append(works, work{url, w.depth + 1})
	}
	return works
}

func download(url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Printf("get url %s error: %v\n", url, err)
		return
	}

	url = strings.TrimPrefix(strings.TrimPrefix(url, "http"), "https")
	f, err := os.Create(url + ".html")
	if err != nil {
		log.Printf("create file %s failed: %v", f.Name(), err)
		return
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Printf("write to file %s error: %v", f.Name(), err)
		return
	}
}
