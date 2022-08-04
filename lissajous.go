package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.RGBA{R: 0x54, G: 0x8c, B: 0x45, A: 0xFF},
	color.White, color.Black,
	color.RGBA{R: 0x00, G: 0x8e, B: 0xc1, A: 0xFF},
	color.RGBA{R: 0xdb, G: 0x00, B: 0x5f},
	color.RGBA{R: 0xfb, G: 0xff, B: 0x90}}

const (
	whiteIndex = 0
	blackIndex = 1
)

//func main() {
//	rand.Seed(time.Now().UTC().UnixNano())
//	if len(os.Args) > 1 && os.Args[1] == "web" {
//		handler := func(w http.ResponseWriter, r *http.Request) {
//			lissajousGreenBlack(w)
//		}
//		http.HandleFunc("/", handler)
//		log.Fatal(http.ListenAndServe("localhost:8080", nil))
//		return
//	}
//	lissajousGreenBlack(os.Stdout)
//}

func lissajous(out io.Writer) {
	const (
		cycles  = 20
		res     = 0.001
		size    = 100
		nframes = 128
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func lissajousGreenBlack(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 128
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.5
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		t := 0.0
		for ; t < cycles*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(len(palette))))
		}
		for ; t < 2*cycles*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(len(palette))))
		}
		for ; t < 2*cycles*math.Pi; t += res {
			x := math.Sin(t*freq + 2*phase)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func lissajousGreenBlackCyc(out io.Writer, cyclesSetting int) {
	const (
		res     = 0.001
		size    = 100
		nframes = 128
		delay   = 8
	)
	cycles := float64(cyclesSetting)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.5
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		t := 0.0
		for ; t < cycles*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(len(palette))))
		}
		for ; t < 2*cycles*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(len(palette))))
		}
		for ; t < 2*cycles*math.Pi; t += res {
			x := math.Sin(t*freq + 2*phase)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
