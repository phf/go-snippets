// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// This is the concurrent version with an insane number of goroutines:
// We start a goroutine for each pixel we want to render.
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

func render(xx, yy int, img *image.RGBA) {
	defer wg.Done()
	y := float64(yy)/float64(height)*(ymax-ymin) + ymin
	x := float64(xx)/float64(width)*(xmax-xmin) + xmin
	z := complex(x, y)
	// Image point (px, py) represents complex value z.
	img.Set(xx, yy, mandelbrot(z))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			wg.Add(1)
			go render(x, y, img)
		}
	}
	wg.Wait()

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}
