package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})       // the signal to cancel all goroutines
var tokens = make(chan struct{}, 20) // control the number of goroutines to be created

type info struct {
	root string
	size int64
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	select {
	case <-done:
		return nil
	case tokens <- struct{}{}: // acquire token
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	defer func() { <-tokens }() // release token
	return entries
}

func walkDir(dir string, fileSizes chan<- info, n *sync.WaitGroup, tokens chan struct{}) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir(subdir, fileSizes, n, tokens)
		} else {
			fileSizes <- info{dir, entry.Size()}
		}
	}
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var tick *time.Ticker
	if *verbose {
		tick = time.NewTicker(500 * time.Millisecond)
	}

	var n sync.WaitGroup

	// Traverse the file tree
	fileSizes := make(chan info)
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, fileSizes, &n, tokens)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		// 利用close的特点，channel会返回零值，实现广播多处同时退出
		close(done)
	}()

	infoMap := make(map[string]int64)

	var nfiles, nbytes int64
	for {
		select {
		case detailInfo, ok := <-fileSizes:
			if !ok {
				break
			}
			nfiles++
			nbytes += detailInfo.size
			infoMap[detailInfo.root] += detailInfo.size
		case <-tick.C:
			printDiskUsage(nfiles, nbytes, infoMap)
		// when done channel close, value received from channel will be zero value
		case <-done:
			// need to drain fileSizes otherwise some goroutines sleep forever
			for range fileSizes {
				// Drain fileSizes to allow existing goroutines to finish
			}
			return
		}
	}
	tick.Stop()
	printDiskUsage(nfiles, nbytes, infoMap)
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage(nfiles int64, nbytes int64, infoMap map[string]int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
	for root, size := range infoMap {
		fmt.Printf("%s files %.1f GB\n", root, float64(size)/1e9)
	}
}
