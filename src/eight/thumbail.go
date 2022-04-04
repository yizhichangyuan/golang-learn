package main

import (
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
	"sync"
)

func makeThumbnails(filenames []string) (thumfiles []string, err error) {
	type item struct {
		thumfile string
		err      error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		// 匿名函数循环变量问题
		go func(f string) {
			var it item
			it.thumfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumfiles = append(thumfiles, it.thumfile)
	}
	return thumfiles, nil
}

// makeThumbnails6 makes thumbnails for each file received from channel
// it returns the number of bytes occupied by the files it creates
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			// 每个goroutine将信号量-1
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	// for range sizes channel only stopped by the channel closed
	// so close the channel calls not in main goroutine, must in another goroutine
	for range sizes {
		total += <-sizes
	}

	return total
}
