package main

import (
	"fmt"
)

type mnode struct {
	name string
	id   int
}

func r1() (rr1 []mnode) {
	rr1 = append(rr1, mnode{"h1", 1})
	rr1 = append(rr1, mnode{"h2", 2})
	return
}

func main() {
	mr1 := r1()
	fmt.Println(mr1)
}
