package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Nanosecond)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return x + x
}

func main() {
	fmt.Println(triple(3))
	bigSlowOperation()
	//fmt.Println(triple(4))
	for _, val := range []int{1, 2, 3, 4} {
		//triple(val)
		//defer fmt.Println(val)
		defer func() { fmt.Println(val) }()
	}
}
