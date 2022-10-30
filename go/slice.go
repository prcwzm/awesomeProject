package main

import (
	"../sliceTest"
	"fmt"
)

func main() {
	a := [...]byte{'a', ' ', ' ', ' ', 'c', 'd', 'e', 'f', ' ', ' ', ' ', ' ', ' ', 'j'}
	c := sliceTest.ShortBlanks(a[:])
	fmt.Println(string(c))
}

func removeNeighbors() {
	a := [...]string{"a", "a", "b", "b", "b", "b", "c", "d", "e", "f", "g"}
	b := sliceTest.RemoveNeighbors(a[:])
	for _, r := range b {
		fmt.Printf(r)
	}
	fmt.Printf("\n")
}

func rotate() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := sliceTest.Rotate(a[:], 3)
	fmt.Println(b)
}

func reverseSlicePtr() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ptr := &a
	sliceTest.ReverseSlicePtr(ptr, 9)
}

func nonemptyAppend() {
	strings := []string{"foo", "bar", "", "baz"}
	strings = sliceTest.NonemptyAppend(strings)
	fmt.Println(strings)
}

func reverseSlice() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceTest.ReverseSlice(a[:])
	fmt.Println(a)
}

func moveN() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceTest.MoveN(a[:], 3)
	fmt.Println(a)
}

func appendInt() {
	a := make([]int, 8, 9)
	a = sliceTest.AppendInt(a, 2)
	fmt.Println(a)
}

func appendIntList() {
	a := make([]int, 8, 9)
	a = sliceTest.AppendIntList(a, 2, 3, 4, 5)
	fmt.Println(a)
}

func nonempty() {
	strings := []string{"foo", "bar", "", "baz"}
	strings = sliceTest.Nonempty(strings)
	fmt.Println(strings)
}
