package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// 1. 用户发送一个消息给服务器，服务器需要给此时所有其他在线用户发送消息
// 因此需要一个中央结构存放所有用户发送过来的消息，一旦该结构中有用户的消息，就应该向对应所有用户发送该消息
// 每个用户是一个goroutines，中央结构一旦有消息就向所有的channel发送消息，但是避免遍历集中在一个goroutines中进行，就应该向每个用户的channel发送消息，因此中央结构应该利用的是channel
// 2. 每个在线用户的消息都是一个channel，专门用于存放该用户发送的消息
// 3. 此外要有一个结构存放所有活跃用户的channel，便于从中央结构拿到消息后就遍历发送这些用户的channel
type client struct{
	Ch chan string
	Conn net.Conn
}

var centerMessages = make(chan string)
var clients = make(map[client]time.Time)
var entering = make(chan client)
var leaving = make(chan client)
var free = make(chan client)
var talking = make(chan struct{})

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
			c.Conn.Close()
			close(c.Ch)
			delete(clients, c)
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
		talking <- struct{}{}
	}

	leaving <- userClient
	centerMessages <- who + " has left"
	defer conn.Close()
}

func checkFree(userClient client) {
	tick := time.NewTicker(5 * time.Minute)
	for {
		select{
		// 添加free而不直接选择关闭原因在于，关闭后还需要删除clients中对应key，由于checkFree和broadCast是不同的goroutines，所以操作同一个map会有线程不安全问题
		case <-tick.C:
			free <- userClient
			break
		case <-talking:
			tick.Reset(5 * time.Minute)
		}
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