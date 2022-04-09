package main

import (
	"fmt"
	"sync"
)

var (
	balance4 int
	mu       sync.RWMutex
)

func Balance4() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance4
}

func Deposit4(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit4(amount)
}

func deposit4(amount int) {
	balance4 -= amount
}

func Withdraw4(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	if balance4-amount >= 0 {
		deposit4(amount)
		return true
	}
	return false
}

func main() {
	go func() {
		Deposit4(200)
		fmt.Println(Balance4())
	}()

	go Deposit4(100)
}
