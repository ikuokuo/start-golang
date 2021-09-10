package main

import (
	"image"
	"image/color"
	"math/cmplx"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.w, m.h)
}

func (m Image) At(x, y int) color.Color {
	return m.mandelbrot(
		complex(4*float64(x)/float64(m.w)-2, 4*float64(y)/float64(m.h)-2), 20)
}

// https://en.wikipedia.org/wiki/Mandelbrot_set
func (m Image) mandelbrot(c complex128, iterMax int) color.Color {
	z := c
	for i := 0; i < iterMax; i++ {
		if cmplx.Abs(z) > 2 {
			return color.Gray{uint8(255 * i / iterMax)}
		}
		z = cmplx.Pow(z, 2) + c
	}
	return color.Black
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
