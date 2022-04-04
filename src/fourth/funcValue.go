package main

import (
	"fmt"
	"strings"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

var fs func(int) int

func add1(r rune) rune {
	return r + 1
}

func main() {
	f := square
	fmt.Printf("%T", f)
	fmt.Println(f(3))

	f = negative
	fmt.Println(f(3))

	//fs(3) // compile error: can't assign func(int, int) int to func(int) int
	fmt.Println(strings.Map(add1, "abc"))
}
