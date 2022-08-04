package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	echoEasy()
	println("%2f elapsed\n", time.Since(start).Seconds())
}

func echoRange() {

	fmt.Println("echoRange!")
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echoJoin() {
	fmt.Println("echoJoin!")
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func testOneDotTwo() {
	for step, value := range os.Args {
		fmt.Println(strconv.Itoa(step) + ":" + value)
	}
}

func testOneDotOne() {
	fmt.Println("test1.1")
	fmt.Println(os.Args[0])
}

func echoEasy() {
	fmt.Println("echoEasy!")
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func BinarySearchMain() {

	var list []int
	fmt.Println("1")
	list = ScanLine()
	fmt.Println(list)
	find := BinarySearch(list, -100)
	fmt.Println(find)
}

func BinarySearch(list []int, iterm int) (find int) {
	low := 0
	hight := len(list) - 1
	for low < hight {
		middle := (hight - low) / 2
		if list[middle] == iterm {
			return middle
		} else if list[middle] > iterm {
			hight = middle - 1
		} else if list[middle] < iterm {
			low = middle + 1
		}
	}
	return -1
}

func ScanLine() (bytes []int) {
	fmt.Println(2)
	var c int
	var save = 0
	var err error
	var pos = false
	for err == nil {
		_, err = fmt.Scanf("%c", &c)
		if c != '\n' {
			if c == ' ' {
				if pos {
					save = save * -1
				}
				bytes = append(bytes, save)
				save = 0
				pos = false
				continue
			} else if c == '-' {
				pos = true
			} else {
				save = save*10 + (c - '0')
			}
		} else {
			bytes = append(bytes, save)
			break
		}
	}
	return bytes
}
