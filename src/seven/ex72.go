package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type NewWriter struct {
	writer  io.Writer
	written int64
}

func (c *NewWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.written = int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := NewWriter{w, 0}
	return &c, &c.written
}

func main() {
	w, n := CountingWriter(ioutil.Discard)
	w.Write([]byte("abc"))
	fmt.Println(*n)
}
