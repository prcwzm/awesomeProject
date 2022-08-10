package main

import (
	"../popcount"
	"fmt"
	"time"
)

func main() {
	num := uint64(12312323453)
	start := time.Now()
	result := popcount.PopCount(num)
	fmt.Printf("popcount.time =%v,result = %g \n", time.Since(start).Nanoseconds(), result)

	start = time.Now()
	result = popcount.PopCount(num)
	fmt.Printf("popcount.time =%v,result = %g \n", time.Since(start).Nanoseconds(), result)

	start = time.Now()
	result = popcount.PopCountSingle(num)
	fmt.Printf("popcountSingle.time =%v ,result = %g \n", time.Since(start).Nanoseconds(), result)

	start = time.Now()
	result = popcount.PopCountRemoveZero(num)
	fmt.Printf("PopCountRemoveZero.time =%v ,result = %g \n", time.Since(start).Nanoseconds(), result)

}
