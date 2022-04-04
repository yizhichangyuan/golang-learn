package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Printf("%q\n", y)
	fmt.Println(y, strconv.Itoa(x))

	// first int64 value, second attribute 进制位数
	fmt.Println(strconv.FormatInt(int64(x), 8))

	fmt.Printf("%o\n", x)

	x, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("error")
	}

	z, err := strconv.ParseInt("101", 2, 64)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(z)
}
