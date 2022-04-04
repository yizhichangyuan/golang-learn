package main

import "fmt"

func main() {
	abort := make(chan struct{})

	// 轮询channel，当channel没有值接收执行defaultcd
	for {
		select {
		case <-abort:
			fmt.Printf("Launch aborted!\n")
		default:
			fmt.Println("default action called")
		}
	}
}
