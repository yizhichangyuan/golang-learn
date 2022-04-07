package main

import (
	"fmt"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"os"
)

var ctx, cancel = context.WithCancel(context.Background())

func get(url string) string {
	req, err := http.NewRequestWithContext(ctx, "get", url, http.NoBody)
	if err != nil {
		return ""
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s", b)
}

func main() {
	responses := make(chan string, len(os.Args[1:]))
	for _, url := range os.Args[1:] {
		responses <- get(url)
	}

	select {
	case x := <-responses:
		cancel()
		fmt.Println(x)
	}
}
