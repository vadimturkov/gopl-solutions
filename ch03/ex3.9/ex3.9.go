/*
	Exercise 3.9
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	width, height = 1024, 1024
)

var (
	zoom       = 1.0
	xmin, xmax = -2.0, 2.0
	ymin, ymax = -2.0, 2.0
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	if fv := r.FormValue("xmin"); fv != "" {
		if value, err := strconv.ParseFloat(fv, 64); err == nil {
			xmin = value
		}
	}

	if fv := r.FormValue("xmax"); fv != "" {
		if value, err := strconv.ParseFloat(fv, 64); err == nil {
			xmax = value
		}
	}

	if fv := r.FormValue("ymin"); fv != "" {
		if value, err := strconv.ParseFloat(fv, 64); err == nil {
			ymin = value
		}
	}

	if fv := r.FormValue("ymax"); fv != "" {
		if value, err := strconv.ParseFloat(fv, 64); err == nil {
			ymax = value
		}
	}

	if fv := r.FormValue("zoom"); fv != "" {
		if value, err := strconv.ParseFloat(fv, 64); err == nil {
			zoom = value
		}
	}

	draw(w)
}

func draw(w io.Writer) {
	sxmin, sxmax := lineScaling(xmin, xmax, zoom)
	symin, symax := lineScaling(ymin, ymax, zoom)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(symax-symin) + symin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(sxmax-sxmin) + sxmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(w, img)
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

func lineScaling(min, max, zoom float64) (smin, smax float64) {
	len := max - min
	mid := min + len/2
	return mid - len/zoom/2, mid + len/zoom/2
}
