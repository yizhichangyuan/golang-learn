package main

import "fmt"

// new return a pointer to int
var p = new(int)

func main() {
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	p = new(int)
	q := new(int)
	fmt.Println(p == q)
}
