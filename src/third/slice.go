package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	var a = []int{1, 2, 4}
	b := a[0:1]
	c := a[1:2]
	a[1] = 3
	fmt.Println(b)
	fmt.Println(c)

	s := "abcd"
	s1 := s[0:2]
	s = "efg" + s
	fmt.Println(s1)

	d := []byte{97, 98, 99}
	e := d[0:1]
	d[0] = 100
	fmt.Println(e)

	// deliver slice to func attribute, slice contains ptr
	reverse(a)
	fmt.Println(a)
	fmt.Println(b)

	g := [...]int{1, 2, 3}
	reverse(g[:])
	fmt.Println(g)

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	runes = []rune("Hello, 世界")
	fmt.Printf("%q\n", runes)
}
