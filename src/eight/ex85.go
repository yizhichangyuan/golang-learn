package main

import (
	"image/png"
	"learn/src/eight/ex85"
	"os"
	"runtime"
)

func main() {
	workers := runtime.GOMAXPROCS(-1)
	img := mandelbrot.Render(workers)
	png.Encode(os.Stdout, img)
}
