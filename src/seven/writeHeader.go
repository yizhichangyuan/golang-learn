package main

import (
	"fmt"
	"io"
)

func writeHeader(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(s string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s)
	}
	return w.Write([]byte(s))
}

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return x
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func main() {
	var x interface{} = 1
	fmt.Println(sqlQuote(x))
}
