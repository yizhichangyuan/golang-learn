package main

import (
	"fmt"
	"time"
)

var balance int

func Deposit(amount int) {
	balance += amount
}

func Balance() int {
	return balance
}

func main() {
	go Deposit(100)

	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
	}()
	time.Sleep(1 * time.Second)
}
