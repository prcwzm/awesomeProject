package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dupLineFile()
}

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		fmt.Printf("%s\n", input.Text())
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dupLineFile() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		if len(counts) != 0 {
			for data, n := range counts {
				if n > 1 {
					fmt.Printf("file:%s\t times:%d data:%s\n", os.Stdin, n, data)
				}
			}
		}
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %s", err)
			}
			countLines(f, counts)
			f.Close()
			if len(counts) != 0 {
				for data, n := range counts {
					if n > 1 {
						fmt.Printf("file:%s\t times:%d data:%s\n", file, n, data)
					}
				}
			}
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
