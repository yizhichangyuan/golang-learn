package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

type StringReader struct {
	s string
}

func (s *StringReader) Read(p []byte) (int, error) {
	copy(p, s.s)
	return len(s.s), io.EOF
}

func NewReader(s string) io.Reader {
	r := StringReader{s}
	return &r
}

func main() {
	_, err := html.Parse(NewReader("<h1>Hello</h1>"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse err: %v", err)
		os.Exit(-1)
	}
}
