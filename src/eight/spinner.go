package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	fmt.Println("end")
}

func spinner(delay time.Duration) {
	go func() {
		for {
			for _, r := range `1\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(delay)
			}
		}
	}()
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
