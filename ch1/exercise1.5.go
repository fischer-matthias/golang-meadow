// lissajous - generates GIF animations of random Lissajous figures
package main

import (
	"image"
	"io"
	"image/gif"
	"image/color"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5 		// number of complete x scillator revolutions
		res = 0.001 	// angular resolution
		size = 100		// images canvas covers [-size...+size]
		nframes = 64	// number of frames
		deplay = 8		// deplay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative freq of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			
			img.SetColorIndex(size+int(x * size + 0.5), size+int(y * size + 0.5), greenIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, deplay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}