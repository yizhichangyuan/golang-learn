package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	return f(s)
}

func replace(s string) string {
	return strings.ReplaceAll(s, "foo", "p")
}

func main() {
	fmt.Println(expand("foofoo", replace))
}
