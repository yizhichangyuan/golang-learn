package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// 防止阻塞下一个客户端请求处理，为每个客户端分配一个goroutine
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scan := bufio.NewScanner(conn)
	cwd := new(strings.Builder)
	for scan.Scan() {
		cmd := strings.Fields(scan.Text())
		oper := cmd[0]

		switch oper {
		case "cd":
			if len(cmd) != 2 {
				io.WriteString(conn, "dir choose not correct")
				continue
			} else {
				dir := cmd[1]
				if strings.IndexByte(dir, 0) == '/' {
					cwd.Reset()
				} else {
					cwd.WriteString("/" + dir)
				}
			}
		case "ls":
			files, err := ioutil.ReadDir(cwd.String())
			if err != nil {
				log.Printf("read dir %s error: %v\n", cwd.String(), err)
				continue
			}
			for _, file := range files {
				fmt.Printf("filename: %s\t size: %dbytes\n", file.Name(), file.Size())
			}
		case "close":
			fmt.Println("close connecting...")
			conn.Close()
		case "get":
			if len(cmd) != 2 {
				io.WriteString(conn, "cannot find file options")
				continue
			}
			getFile(conn, cmd[1])
		case "send":
			if len(cmd) < 2 {
				io.WriteString(conn, "filename and fileContent cannot be empty")
				continue
			}
			err := createFile(cmd[1], cmd[2])
			if err != nil {
				io.WriteString(conn, "write file error"+err.Error())
				continue
			}
		}
	}
}

func getFile(conn net.Conn, fileDir string) {
	f, err := os.Open(fileDir)
	defer f.Close()
	if os.IsNotExist(err) {
		io.WriteString(conn, "cannot find file: %s"+fileDir)
		return
	}
	buf := make([]byte, 1000)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("read file %s error: %v\n", fileDir, err)
			break
		}
		conn.Write(buf[:n])
	}
}

func createFile(fileName, fileContent string) error {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		log.Fatalf("create file %s error: %v", fileName, err)
		return err
	}
	_, err = io.Copy(f, strings.NewReader(fileContent))
	if err != nil {
		log.Fatalf("write content into file %s error, content: %s, error: %v\n", fileName, fileContent, err)
		return err
	}
	return nil
}
