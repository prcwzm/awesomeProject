package main

import (
	"../iota"
	"fmt"
)

func main() {
	fmt.Println(
		"KB", iota.KB,
		"MB", iota.MB,
		"GB", iota.GB,
		"TB", iota.TB,
		"PB", iota.PB,
		"EB", iota.EB,
	)
}
