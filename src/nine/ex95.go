package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var ops uint64

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		for {
			// count atomically the number of communication between two goroutine
			atomic.AddUint64(&ops, 1)
			ch1 <- struct{}{}
			<-ch2
		}
	}()

	go func() {
		for {
			<-ch1
			ch2 <- struct{}{}
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(atomic.LoadUint64(&ops))
}
