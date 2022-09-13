package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
)

func main() {
	http.HandleFunc("/mandelbrot", handleMandelbrot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handleMandelbrot(writer http.ResponseWriter, r *http.Request) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 8192, 8192
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(writer, img)
}

func mandelbrot(z complex128) color.Color {
	const iterarions = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterarions; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast}
		}
	}
	return color.Black
}
