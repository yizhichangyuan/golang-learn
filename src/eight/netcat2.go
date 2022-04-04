package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	tcpConn := conn.(*net.TCPConn)
	// 无缓冲的chan可以用来作为同步chan
	done := make(chan struct{})
	go func() {
		mustCopys(os.Stdout, tcpConn)
		tcpConn.CloseRead()
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopys(tcpConn, os.Stdin)
	tcpConn.CloseWrite()
	// 用户输入终止，main goroutine需要等待匿名函数的goroutine运行完毕才结束main
	// 因此main goroutine从chan等待接收，在匿名goroutine没有发送数据之前会一直阻塞直到接收到数据
	<-done
}

func mustCopys(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}
