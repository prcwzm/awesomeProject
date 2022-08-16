package surface

import (
	"fmt"
	"io"
	"math"
	"os"
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
			ax, ay, _, color := corner(i+1, j)
			bx, by, _, color := corner(i, j)
			cx, cy, _, color := corner(i, j+1)
			dx, dy, _, color := corner(i+1, j+1)
			fmt.Printf("<polygon points=' %g,%g,%g,%g,%g,%g,%g,%g,' %s/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
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
			ax, ay, ok, color := corner(i+1, j)
			bx, by, ok, color := corner(i, j)
			cx, cy, ok, color := corner(i, j+1)
			dx, dy, ok, color := corner(i+1, j+1)
			if ok > 0 {
				continue
			}
			fmt.Fprintf(out, "<polygon points=' %g,%g,%g,%g,%g,%g,%g,%g,' %s/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (sx float64, sy float64, ok float64, color string) {
	var colorStart, colorEnd string
	if len(os.Args) == 2 {
		colorStart = os.Args[1]
		colorEnd = "#0000ff"
	} else if len(os.Args) == 3 {
		colorStart = os.Args[1]
		colorEnd = os.Args[2]
	} else if len(os.Args) == 1 {
		colorStart = "#ff0000"
		colorEnd = "#0000ff"
	}

	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	if z > 200 {
		color = fmt.Sprintf("stroke='%s' fill='%s' stroke-width='1'", colorStart, colorEnd)
	} else if z < 50 {
		color = fmt.Sprintf("stroke='%s' fill='%s' stroke-width='1'", colorEnd, colorStart)
	}
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	if math.IsInf(sx, 0) || math.IsInf(sy, 0) {
		ok += 1
	}
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
