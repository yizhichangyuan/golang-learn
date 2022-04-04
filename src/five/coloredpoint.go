package main

import (
	"fmt"
	"image/color"
	"math"
)

type Points struct{ X, Y float64 }
type Paths []Points

type ColoredPoint struct {
	Points
	Color color.RGBA
}

func (p Points) Distance(q Points) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Points) Add(q Points) Points { return Points{p.X + q.X, p.Y + q.Y} }
func (p Points) Sub(q Points) Points { return Points{p.X - q.X, p.Y - q.Y} }

func (paths Paths) TranslateBy(offset Points, add bool) {
	var op func(p, q Points) Points
	if add {
		// 直接调用类型的方法，会使用参数列表第一个参数作为接收器，所以函数值都是func(p, q Points) Points
		op = Points.Add
	} else {
		op = Points.Sub
	}
	for i := range paths {
		paths[i] = op(paths[i], offset)
	}
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Points.X)
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Points{1, 1}, red}
	var q = ColoredPoint{Points{5, 4}, blue}
	fmt.Println(p.Distance(q.Points))
}
