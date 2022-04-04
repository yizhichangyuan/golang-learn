package main

import (
	"math"
)

type Point struct{ X, Y float64 }
type Path []Point

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type P *int

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func setNil(ints []int) {
	ints = nil
}

//
//func main() {
//	p := Point{1, 2}
//	q := Point{4, 6}
//	fmt.Println(Distance(p, q))
//	fmt.Println(p.Distance(q))
//
//	perim := Path{{1,1}, {5,1}, {5,4}, {1, 1}}
//	fmt.Println(perim.Distance())
//	fmt.Println(Path.Distance(perim))
//
//	var b = 1
//	var x P = &b
//	fmt.Println(*x)
//
//	p.ScaleBy(2)
//	fmt.Println(p)
//
//	var buf bytes.Buffer
//	buf.Write([]byte("abc"))
//
//	var buf2 bytes.Buffer
//	io.Copy(&buf2,&buf)
//	tmp := []byte("d")
//	buf2.Write(tmp)
//
//	fmt.Println(buf2.String())
//	fmt.Println(buf.String())
//
//	s1 := [2]string{"a"}
//	s2 := s1
//	s2[0] = "b"
//	fmt.Println(s1)
//	fmt.Println(s2)
//
//	z := []int{1,2,3}
//	setNil(z)
//	fmt.Println(z)
//}
