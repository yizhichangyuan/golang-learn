package main

import (
	"fmt"
	"testing"
)

func Test(T *testing.T) {
	fmt.Println("hello")
	fmt.Errorf("%s", "hello")
}
