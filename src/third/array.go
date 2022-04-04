package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int
	fmt.Println(a[0])

	// uninitialized array value has default value : 0
	var q [2]int = [2]int{1, 2}
	fmt.Println(q[1])

	p := [...]int{2, 1}
	fmt.Println(len(p))

	// init array value based index and value
	symbol := [...]string{USD: "$", EUR: "#"}
	fmt.Println(EUR, symbol[EUR])

	// unused index will be init with default value 0
	// index 99 will be init with -1
	r := [...]int{99: -1}
	fmt.Println(r[99])

	fmt.Println(p == q)

}
