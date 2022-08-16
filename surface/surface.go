package surface

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func Surface() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width = '%d' height = '%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			fmt.Printf("<polygon points=' %g,%g,%g,%g,%g,%g,%g,%g,'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Printf("</svg>")
}

func SurfaceWeb(out io.Writer) {
	_, err := fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width = '%d' height = '%d'>", width, height)
	if err != nil {
		return
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			bx, by, ok := corner(i, j)
			cx, cy, ok := corner(i, j+1)
			dx, dy, ok := corner(i+1, j+1)
			if ok > 0 {
				continue
			}
			fmt.Fprintf(out, "<polygon points=' %g,%g,%g,%g,%g,%g,%g,%g,'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (sx float64, sy float64, ok float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	if math.IsInf(sx, 0) || math.IsInf(sy, 0) {
		ok += 1
	}
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}
