package newton

import (
	"image/color"
	"math/cmplx"
)

func NewtonColorLoop(z complex128, boundary float64) color.Color {
	var times uint8 = 0
	for cmplx.Abs(fz4(z)) > boundary && times < 255 {
		times += 1
		z = z - fz4(z)/oneFunc(z)
	}
	return color.Gray{255 - times}
}

func fz4(z complex128) complex128 {
	return z*z*z*z - 1
}

func oneFunc(z complex128) complex128 {
	return 4 * z * z * z
}
