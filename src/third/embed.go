package main

//type Circle struct {
//	X, Y, Radius int
//}
//
//type Wheel struct {
//	X, Y, Radius,Spokes int
//}
//
//var w Wheel

type Points struct {
	X, Y int
}

type Circle struct {
	Points
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	//w.X = 8
	//w.Y = 8
	//w.Radius = 5
	//w.Spokes = 20

	var w Wheel
	w.X = 8
	w.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
}
