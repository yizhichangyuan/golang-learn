package mandelbrot

import (
	"image"
	"image/color"
	"math/cmplx"
	"sync"
)

func Render(worker int) *image.RGBA {
	const (
		width, height          = 1024, 1024
		xmin, xmax, ymin, ymax = -2, 2, -2, 2
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	rows := make(chan int, height)
	for row := 0; row < height; row++ {
		rows <- row
	}
	close(rows)
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			for py := range rows {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
