/*
	Exercise 3.7
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func newton(z complex128) color.Color {
	const iterations = 50
	const contrast = 15
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return chooseColor(z, contrast*i)
		}
	}
	return color.Black
}

func chooseColor(z complex128, contrast uint8) color.Color {
	root := complex(math.Round(real(z)), math.Round(imag(z)))

	switch root {
	case 1:
		return color.RGBA{255 - contrast, 0, 0, 255}
	case -1:
		return color.RGBA{255, 255 - contrast, 0, 255}
	case 1i:
		return color.RGBA{0, 255 - contrast, 0, 255}
	case -1i:
		return color.RGBA{0, 0, 255 - contrast, 255}
	}

	return color.Black
}
