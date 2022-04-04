package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type LineCounter int

type WordsCounter int

func (l *LineCounter) Write(p []byte) (n int, err error) {
	scan := bufio.NewScanner(bytes.NewReader(p))
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		*l++
	}
	return len(p), nil
}

func (w *WordsCounter) Write(p []byte) (n int, err error) {
	scan := bufio.NewScanner(bytes.NewReader(p))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		*w++
	}
	return len(p), nil
}

func main() {
	b := "hello, world\nls"
	var l LineCounter
	fmt.Fprintf(&l, b)
	fmt.Println(l)

	b = "a b c d"
	var w WordsCounter
	fmt.Fprintf(&w, b)
	fmt.Println(w)
}
