package main

import "fmt"

func isShuffer(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[rune]int, len(a))
	for _, v := range a {
		m[v]++
	}

	for _, v := range b {
		m[v]--
		if v < 0 {
			return false
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isShuffer("abcb", "bcad"))
}
