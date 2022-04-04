package main

import (
	"fmt"
)

type Weekday int
type Flags uint
type Data uint

func main() {
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Saturday)

	const (
		FlagUp Flags = 1 << iota
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast
	)
	fmt.Println(FlagLoopback)

	const (
		_ = (1 << (10 * iota)) * 8
		KB
		MB
		GB
		YB
	)

	fmt.Println(KB)

	var f float64
	fmt.Printf("\t%d\n", (f-32)*5/9)
	fmt.Printf("\t%d\n", 5/9*(f-32))
	fmt.Printf("\t%d\n", 5.0/9.0*(f-32))

	var v float64 = 3 + 0i
	v = 2
	fmt.Printf("%T\t%q\n", v, v)
	v = 'a'
	fmt.Printf("%T\t%q\n", v, v)
	fmt.Println(10 / 4.0)
	var x = 10
	fmt.Println(x / 4.0)

	const a = 1
	const b = 1i
	fmt.Printf("%T", a/b)
}
