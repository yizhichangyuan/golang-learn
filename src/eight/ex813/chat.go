package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)
// 利用定时器定时发送消息的特点，用户向centerMessages发送对话内容的同时向talking发送一个消息
// 如果talking收到消息就重置定时器，如果没有收到消息就向free发送该客户端准备清理程序
// broadcast中多路复用检查到free中有消息，就清理该客户端，主动断开连接
// 好处：不需要记录该用户发送消息的实时时间，然后再对比
type client struct{
	Ch chan string
	Conn net.Conn
}

const timeout = 1 * time.Minute

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
				select{
				case c.Ch <- msg:
				case <-time.After(10 * time.Second):
				}
			}
		case c := <-entering:
			clients[c] = time.Now()
		case c := <-leaving:
			delete(clients, c)
			close(c.Ch)
		case c := <-free:
			fmt.Fprintf(c.Conn, "connection has closed becauase of timeout 5 minutes\n")
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
		talking <- struct{}{}
	}

	leaving <- userClient
	centerMessages <- who + " has left"
}

func checkFree(userClient client) {
	tick := time.NewTicker(timeout)
	for {
		select{
		// 添加free而不直接选择关闭原因在于，关闭后还需要删除clients中对应key，由于checkFree和broadCast是不同的goroutines，所以操作同一个map会有线程不安全问题
		case <-tick.C:
			free <- userClient
			break
		case <-talking:
			tick.Reset(timeout)
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