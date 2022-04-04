package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// 目的将对应服务器时间和ip绑定在一起
type server struct {
	name    string
	ip      string
	message string
}

func main() {
	servers := parse([]string{"NewYork=localhost:8010", "Tokyo=localhost:8020", "London=localhost:8030"})
	// 1: 依次和每个服务器建立，然后读取时间直接打印，不符合题目一次显示所有时间
	// 2：和每个服务器建立，然后读取时间存放（go routines否则阻塞下一个服务器读取），另一个for打印信息
	// 3：for range修改每个内部变量的成员，受限于range局部变量拷贝，无法直接修改
	// 4: 所以要么索引要么声明为指针数组
	for i := 0; i < len(servers); i++ {
		conn := dial(servers[i].ip)
		defer conn.Close()
		// 传入指针，目的修改该结构体成员，否则值拷贝
		go func(s *server) {
			sc := bufio.NewScanner(conn)
			for sc.Scan() {
				s.message = sc.Text()
			}
		}(&servers[i])
	}

	for {
		fmt.Printf("\n")
		for _, server := range servers {
			fmt.Printf("%s: %s\n", server.name, server.message)
		}
		fmt.Print("------")
		time.Sleep(time.Second)
	}
}

func dial(ip string) net.Conn {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Fatalf("wrong dial with %s: %v", ip, err)
	}
	return conn
}

func parse(in []string) []server {
	var servers []server
	for _, str := range in {
		s := strings.SplitN(str, "=", 2)
		servers = append(servers, server{s[0], s[1], ""})
	}
	return servers
}
