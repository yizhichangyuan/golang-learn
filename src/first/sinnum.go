package main

func sinnum(x int) int {
	switch {
	case x > 0:
		return +1
	case x < 0:
		return -1
	default:
		return 0
	}
}

type Point struct {
	x, y int
}

var p Point
