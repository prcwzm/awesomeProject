package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"os"
	"strconv"
)

const (
	xmin, xmax, ymin, ymax = -2, +2, -2, +2
	width, height          = 1024, 1024
)
const iterations = 200
const constraints = 15

func Mandelbrot(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	err := png.Encode(w, img)
	if err != nil {
		return
	}

}

func mandelbrot(z complex128) color.Color {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - constraints*n}
		}
	}
	return color.Black
}

func MandelbrotColor(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotColor(z))
		}
	}
	err := png.Encode(w, img)
	if err != nil {
		return
	}

}

func mandelbrotColor(z complex128) color.Color {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{255 - constraints*n, 255 - constraints*n, 255 - constraints*n}
		}
	}
	return color.White
}

func SuperSampleMax(w io.Writer) {
	superWidth := width
	superHeight := height
	var err error
	if os.Args[1] != "" {
		superWidth, err = strconv.Atoi(os.Args[1])
	}
	if os.Args[2] != "" {
		superHeight, err = strconv.Atoi(os.Args[1])
	}

	img := image.NewRGBA(image.Rect(0, 0, superWidth, superHeight))
	for py := 0; py < superHeight; py++ {
		y := float64(py)/float64(superHeight)*(ymax-ymin) + ymin
		for px := 0; px < superWidth; px++ {
			x := float64(px)/float64(superWidth)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotColor(z))
		}
	}
	err = png.Encode(w, img)
	if err != nil {
		return
	}
}

func SuperSample(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	imgReal := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotColor(z))
		}
	}
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			colorSet := make([]color.Color, 5)
			colorsWeight(px, py, img, colorSet)
			colorsWeight(px+1, py, img, colorSet)
			colorsWeight(px-1, py, img, colorSet)
			colorsWeight(px, py+1, img, colorSet)
			colorsWeight(px, py-1, img, colorSet)

			imgReal.Set(px, py, mixColor(colorSet))
		}
	}
	png.Encode(w, imgReal)

}

func mixColor(colorSet []color.Color) color.Color {
	var r, g, b, a uint32 = 0, 0, 0, 0
	setLength := uint32(len(colorSet))
	for i := 0; i < int(setLength); i++ {
		tr, tg, tb, ta := colorSet[i].RGBA()
		r += tr / setLength
		g += tg / setLength
		b += tb / setLength
		a += ta / setLength
	}
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}

}

func colorsWeight(px int, py int, img *image.RGBA, set []color.Color) {
	if px < 0 || py < 0 || px > height || py > width {
		return
	}
	set = append(set, img.At(px, py))
	return
}

func superSample(x, y float64) color.Color {
	deltaX := float64(0.5) / width * (xmax - xmin)
	deltaY := float64(0.5) / height * (ymax - ymin)
	colorA := mandelbrotYCbCr(complex(x+deltaX, y))
	colorB := mandelbrotYCbCr(complex(x-deltaX, y))
	colorC := mandelbrotYCbCr(complex(x, y+deltaY))
	colorD := mandelbrotYCbCr(complex(x, y-deltaY))
	colorE := mandelbrotYCbCr(complex(x, y))
	return color.YCbCr{Y: (colorA.Y + colorB.Y + colorC.Y + colorD.Y + colorE.Y) / 5,
		Cb: (colorA.Cb + colorB.Cb + colorC.Cb + colorD.Cb + colorE.Cb) / 5,
		Cr: (colorA.Cr + colorB.Cr + colorC.Cr + colorD.Cr + colorE.Cr) / 5}
}

func mandelbrotYCbCr(z complex128) color.YCbCr {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{Y: 255 - constraints*n, Cb: 255 - constraints*n, Cr: 255 - constraints*n}
		}
	}
	return color.YCbCr{}
}
