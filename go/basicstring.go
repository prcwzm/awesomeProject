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
	fmt.Printf("%s\n", basename.Basenameeasy(input))
}
