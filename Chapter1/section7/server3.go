package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/lissajous", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(writer, "Hander[%q] = %q\n", k, v)
	}
	fmt.Fprintf(writer, "Host = %q\n", r.Host)
	fmt.Fprintf(writer, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(writer, "Form[%q] = %q\n", k, v)
	}
}

func lissajous(writer http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())
	const (
		res     = 0.001 // angular resolution
		size    = 250   // image canvas covers [-size..+size]
		nframes = 256   // number of animation frames
		delay   = 16    // delay between frames in 10ms units
	)
	var cycles float64 = 5 // number of complete x oscillator revolutions
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "cycles" {
			cycle, err := strconv.Atoi(v[0])
			if err != nil {
				cycles = 5
			}
			cycles = float64(cycle)
		}
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size, 2*size)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size), size+int(y*size),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(writer, &anim) // NOTE: ignoring encoding errors
}
