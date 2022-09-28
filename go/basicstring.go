package main

import (
	"../basename"
	"fmt"
	"os"
)

func main() {
	var input string
	if len(os.Args) < 2 {
		input = ""
	} else {
		input = os.Args[1]
	}
	fmt.Printf("%s\n", basename.NonRecursionComma(input))
}

func comma() {
	var input string
	if len(os.Args) < 2 {
		input = ""
	} else {
		input = os.Args[1]
	}
	fmt.Printf("%s\n", basename.Comma(input))
}

func intToString() {
	fmt.Println(basename.IntToString([]int{1, 2, 3}))
}
