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

	// select会等待case中有能够执行的case去执行，select才会去通信并执行case之后的语句
	// 这时候其他通信不会执行，执行对应case之后就会直接退出
	// select多路复用保证可以同时监听多个channel，对应channel发送执行相关逻辑
	// 可以防止顺序监听有一个channel没有接收造成阻塞
	select {
	case <-abort:
		// 中途如果检测用户输入，则终止发射
		abortLaunch()
	case <-time.After(10 * time.Second):
		// 十秒后发射事件
		launch()
	}
	fmt.Println("end...")
}

func launch() {
	fmt.Println("launch")
}

func abortLaunch() {
	fmt.Println("abort")
}
