package mapable

import (
	"bufio"
	"fmt"
	"os"
)

func Dedup() {
	screen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !screen[line] {
			screen[line] = true
			fmt.Println(line)
			if line == "out" {
				os.Exit(0)
			}
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup : %v\n", err)
		os.Exit(1)
	}
}
