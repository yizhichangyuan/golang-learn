package main

import "C"
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func handConn1(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	in := make(chan struct{})

	go func() {
		for input.Scan() {
			// when client answer, send message to channel
			in <- struct{}{}
		}
	}()

	tick := time.NewTimer(10 * time.Second)

	for {
		select {
		case <-tick.C:
			// when client take no action for timeout 10 seconds, server close connection automatically
			fmt.Fprintln(c, "connection closed by timeout.")
			c.Close()
		case <-in:
			go echo1(c, input.Text(), 1*time.Second, wg)
			// when client answer in 10 seconds, tick reset
			tick.Reset(10 * time.Second)
		}
	}
	wg.Wait()
	defer c.Close()
}

func echo1(c net.Conn, text string, delay time.Duration, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(text))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handConn1(conn)
	}
}
