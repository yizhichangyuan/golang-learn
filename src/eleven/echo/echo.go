package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

var out io.Writer = os.Stdout // modified during testing

func main() {
	flag.Parse()
	if err := echo(!*n, flag.Args(), *sep); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

// 使用该函数增加参数目的：减少对全局变量的依赖；测试可以直接测试该函数即可
func echo(newline bool, args []string, sep string) (err error) {
	_, err = fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(out)
	}
	return
}
