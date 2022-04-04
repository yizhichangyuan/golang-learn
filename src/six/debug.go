package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func f(out io.Writer) {
	if out != nil {
		fmt.Println("access")
		out.Write([]byte("done\n"))
	}
}

func main() {
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
	}
}
