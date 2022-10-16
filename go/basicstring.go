package main

import (
	"../basename"
	"fmt"
	"os"
)

func main() {
	var left, right string
	if len(os.Args) < 3 {
		left = ""
		right = ""
	} else {
		left = os.Args[1]
		right = os.Args[2]
	}
	fmt.Printf("%s\n", basename.Heterogeneous(left, right))
}

func floatComma() {
	var input string
	if len(os.Args) < 2 {
		input = ""
	} else {
		input = os.Args[1]
	}
	fmt.Printf("%s\n", basename.FloatComma(input))
}

func nonRecursionComma() {
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
