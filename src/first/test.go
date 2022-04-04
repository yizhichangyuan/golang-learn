package main

import "fmt"

func modifySlice(s []int) {
	s[0] = 1
}

var s2 []int = []int{0, 0, 0}

//s2 := make([]int, 3)

func main() {
	fmt.Printf("%#v\n", s2) //[]int{0, 0, 0}
	modifySlice(s2)
	fmt.Printf("%#v\n", s2) //[]int{1, 0, 0}
}
