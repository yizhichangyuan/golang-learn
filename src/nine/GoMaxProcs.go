package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Int()
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
