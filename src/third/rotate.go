package main

import "fmt"

func rotate(s []int, n int) []int {
	n %= len(s)
	s = append(s, s[:n]...)
	return s[n:]
}

func main() {
	s := []int{0, 1, 2, 3, 4}
	s = rotate(s, 2)
	fmt.Println(s)
}
