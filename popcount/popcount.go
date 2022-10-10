package popcount

var pc [255]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x)] +
		pc[byte(x>>8)] +
		pc[byte(x>>16)] +
		pc[byte(x>>24)] +
		pc[byte(x>>32)] +
		pc[byte(x>>40)] +
		pc[byte(x>>48)] +
		pc[byte(x>>56)])
}

func PopCountSingle(x uint64) (count int) {
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x = x >> 1
	}
	return
}

func PopCountRemoveZero(x uint64) (count int) {
	for x > 0 {
		count++
		x = x & (x - 1)
	}
	return
}

func H264PopCount(x [32]uint8, y [32]uint8) (count int) {
	for i := 0; i < 32; i++ {
		if !(x[i] == y[i]) {
			count += popCount(x[i], y[i])
		}
	}
	return
}

func popCount(x uint8, y uint8) (count int) {
	for i := 0; i < 8; i++ {
		xl := x & 0x01
		yl := y & 0x01
		if xl^yl == 1 {
			count++
		}
		x = x >> 1
		y = y >> 1
	}
	return
}
