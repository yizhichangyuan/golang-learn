package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Connecting countdown. Press return to abort")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case x := <-tick:
			fmt.Println(x)
		case <-abort:
			fmt.Println("launch aborted!")
			return
		}
	}
	launch1()
}

func launch1() {
	fmt.Println("launch.")
}
