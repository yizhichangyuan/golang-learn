package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// 利用sync.Map，handle goroutines中用户一旦发送消息就更新sync.Map中用户发送消息时间
// checkFree goroutines专门检查该用户的时间是否超过5分钟，超过5分钟就向free发送消息，利用多路复用关闭该客户端

const timeout = 1 * time.Minute

type client struct {
	Ch   chan string
	Conn net.Conn
}

var centerMessages = make(chan string)
var clients = make(map[client]time.Time)
var entering = make(chan client)
var leaving = make(chan client)
var free = make(chan client)
var lastMessageTime sync.Map

func broadcast() {
	for {
		select {
		case msg := <-centerMessages:
			for c := range clients {
				c.Ch <- msg
			}
		case c := <-entering:
			clients[c] = time.Now()
		case c := <-leaving:
			delete(clients, c)
			close(c.Ch)
		case c := <-free:
			fmt.Fprintf(c.Conn, "connection has closed becauase of timeout 5 minutes\n")
			lastMessageTime.Delete(c.Conn)
			c.Conn.Close()
			// 一旦close完，for in.Scan会退出，然后进入leaving会关闭clients
			// 如果在这儿在进行delete(clients, c) close(c.Ch)会panic:  close of closed channel
		}
	}
}

func clientWrite(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func handleChat(conn net.Conn) {
	ch := make(chan string)
	userClient := client{ch, conn}
	go clientWrite(conn, ch)
	go checkFree(userClient)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	entering <- userClient
	centerMessages <- who + " has arrived"

	in := bufio.NewScanner(conn)
	for in.Scan() {
		centerMessages <- who + ": " + in.Text()
		lastMessageTime.Store(userClient, time.Now()) // update last send message time
	}

	leaving <- userClient
	centerMessages <- who + " has left"
	defer conn.Close()
}

func checkFree(userClient client) {
	for {
		v, ok := lastMessageTime.Load(userClient)
		if !ok {
			// 可能用户登录后5分钟内还没有发消息也就没有注册到lastMessageTime
			time.Sleep(timeout)
			continue
		}
		last := v.(time.Time)
		if time.Now().Sub(last) > timeout {
			free <- userClient
			break
		}
		time.Sleep(time.Now().Sub(last)) // 休息剩下的时间然后唤醒再检查够不够5分钟
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}
	go broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleChat(conn)
	}
}
