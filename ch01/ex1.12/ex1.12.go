/*
	Exercise 1.12
*/

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	blackIndex    = 1
	defaultCycles = 5
	defaultDelay  = 8
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	cycles := defaultCycles
	if fv := r.FormValue("cycles"); fv != "" {
		if value, err := strconv.Atoi(fv); err == nil {
			cycles = value
		}
	}

	delay := defaultDelay
	if fv := r.FormValue("delay"); fv != "" {
		if value, err := strconv.Atoi(fv); err == nil {
			delay = value
		}
	}

	lissajous(w, cycles, delay)
}

func lissajous(out io.Writer, cycles, delay int) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
