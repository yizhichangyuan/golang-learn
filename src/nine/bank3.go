package main

import "sync"

var (
	balance3 int
	m        sync.Mutex
)

func Deposit3(amount int) {
	m.Lock()
	defer m.Unlock()
	deposit3(amount)
}

func deposit3(amount int) {
	balance3 += amount
}

func Balance3() int {
	m.Lock()
	defer m.Unlock()
	return balance3
}

func Withdraw3(amount int) bool {
	m.Lock()
	defer m.Unlock()
	if balance3-amount >= 0 {
		deposit3(balance3)
		return true
	}
	return false
}
