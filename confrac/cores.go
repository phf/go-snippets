// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// This is the concurrent version with a sane number of goroutines:
// For n CPUs we split rendering into n slices.
package main

import (
	"image"
	"image/png"
	"os"
	"runtime"
	"sync"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

var wg sync.WaitGroup

func render(xx, yy, w, h int, img *image.RGBA) {
	defer wg.Done()
	for py := yy; py < h; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := xx; px < w; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
}

func main() {
	ncpu := runtime.NumCPU()
	runtime.GOMAXPROCS(ncpu)

	slice := height / ncpu
	leftover := height % ncpu

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < ncpu; i++ {
		ys := i*slice
		wg.Add(1)
		go render(0, ys, width, ys+slice, img)
	}
	if leftover != 0 {
		ys := ncpu*slice
		wg.Add(1)
		go render(0, ys, width, ys+leftover, img)
	}
	wg.Wait()

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}
