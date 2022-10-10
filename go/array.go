package main

import (
	"../array"
	"../popcount"
	"crypto/sha256"
	"fmt"
)

func main() {
	array.HashSHA([]byte("x"), 0)
}

func shaPopCount() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(popcount.H264PopCount(c1, c2))
}
