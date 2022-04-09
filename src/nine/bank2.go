package main

import (
	"fmt"
	"time"
)

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan draw)

type draw struct {
	amount  int
	success chan bool
}

func BalanceCon() int {
	return <-balances
}

func DepositCon(amount int) {
	deposits <- amount
}

func Withdraw(amount int) bool {
	succeed := make(chan bool)
	withdraws <- draw{amount, succeed}
	return <-succeed
}

// monitor goroutine
func bank() {
	var balance1 int
	for {
		select {
		case balances <- balance1:
		case amount := <-deposits:
			balance1 += amount
		case draws := <-withdraws:
			if draws.amount <= balance1 {
				balance1 -= draws.amount
				draws.success <- true
			} else {
				draws.success <- false
			}
		}
	}
}

func main() {
	go func() {
		bank()
	}()

	go func() {
		DepositCon(200)
		fmt.Println("=", BalanceCon())
	}()

	go DepositCon(100)

	go fmt.Println(Withdraw(500))

	time.Sleep(1 * time.Second)
}
