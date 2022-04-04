package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type BufReader struct {
	reader      io.Reader
	remainBytes int64
}

func (b *BufReader) Read(p []byte) (int, error) {
	if b.remainBytes <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > b.remainBytes {
		p = p[:b.remainBytes]
	}
	n, err := b.reader.Read(p)
	b.remainBytes -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &BufReader{reader: r, remainBytes: n}
}

func main() {
	lr := LimitReader(strings.NewReader("12345"), 2)
	b, err := ioutil.ReadAll(lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}
	fmt.Printf("%s\n", b)
}
