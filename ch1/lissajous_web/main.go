// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
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

// web

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// The sequence of image is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		lissajous(w, r)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

const (
	DEFAULT_CYCLES = 5   // number of complete x oscillator revolutions
	DEFAULT_SIZE   = 100 // image canvas covers [-size..+size]
)

func lissajous(out io.Writer, r *http.Request) {
	const (
		res     = 0.001 // angular resolution
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	var theCycle = getCycles(r)
	var theSize = getSize(r)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*theSize+1, 2*theSize+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(theCycle)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(theSize+int(x*float64(theSize)+0.5), theSize+int(y*float64(theSize)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func getCycles(r *http.Request) int {
	if v, ok := r.Form["cycles"]; ok {
		if cycleParam, err := strconv.Atoi(v[0]); err != nil {
			fmt.Printf("Cannot convert cycles to a number defaulting to %v\n", DEFAULT_CYCLES)
			return DEFAULT_CYCLES
		} else {
			return cycleParam
		}
	}
	return DEFAULT_CYCLES
}

func getSize(r *http.Request) int {
	if v, ok := r.Form["size"]; ok {
		if sizeParam, err := strconv.Atoi(v[0]); err != nil {
			fmt.Printf("Cannot convert size to a number defaulting to %v\n", DEFAULT_SIZE)
			return DEFAULT_SIZE
		} else {
			return sizeParam
		}
	}
	return DEFAULT_SIZE
}
