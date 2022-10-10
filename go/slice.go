package main

import (
	"../sliceTest"
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceTest.MoveN(a[:], 3)
	fmt.Println(a)
}

func reverseSlice() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceTest.ReverseSlice(a[:])
	fmt.Println(a)
}
