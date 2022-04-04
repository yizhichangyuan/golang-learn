package main

import (
	"errors"
	"fmt"
	"syscall"
)

func main() {
	var err error = syscall.Errno(2)
	fmt.Println(err.Error())

	err1 := errors.New("abc")
	fmt.Println(err1)
}
