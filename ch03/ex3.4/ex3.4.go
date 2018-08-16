/*
	Exercise 3.4
*/

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var (
	sin30   = math.Sin(angle)
	cos30   = math.Cos(angle)
	width   = 600.0
	height  = 320.0
	xyscale = width / 2 / xyrange
	zscale  = height * 0.4
	color   = "#ffffff"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "image/svg+xml")

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	if fv := r.FormValue("width"); fv != "" {
		if value, err := strconv.Atoi(fv); err == nil {
			width = float64(value)
		}
	}

	if fv := r.FormValue("height"); fv != "" {
		if value, err := strconv.Atoi(fv); err == nil {
			height = float64(value)
		}
	}

	if fv := r.FormValue("color"); fv != "" {
		color = fv
	}

	draw(w)
}

func draw(w io.Writer) {
	var svg string
	svg += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%g' height='%g'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			svg += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	svg += fmt.Sprintln("</svg>")
	w.Write([]byte(svg))
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	if r == 0 {
		return 1
	}
	return math.Sin(r) / r
}
