package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	a := StringSlice{"b", "c", "a"}
	sort.Sort(a)
	fmt.Println(a)
}
