package sliceTest

import (
	"../sliceStack"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func ReverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ReverseBytes(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ReverseSlicePtr(s *[9]int, length int) {
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func MoveN(s []int, n int) {
	ReverseSlice(s[:n])
	ReverseSlice(s[n:])
	ReverseSlice(s)
}

func AppendInt(x []int, y int) []int {
	var z []int
	zLen := len(x) + 1
	if zLen < cap(x) {
		z = x[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func AppendIntList(x []int, y ...int) []int {
	var z []int
	zLen := len(x) + len(y)
	if zLen < cap(x) {
		z = x[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func Nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func NonemptyAppend(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func Rotate(x []int, r int) (y []int) {
	xLen := len(x)
	y = make([]int, xLen)
	for i := 0; i < xLen; i++ {
		ptr := i + r
		if ptr >= xLen {
			ptr = ptr - xLen
		}
		y[i] = x[ptr]
	}
	return
}

func RemoveNeighbors(strList []string) (rs []string) {
	stack := sliceStack.SliceStack[string]{}
	if len(strList) == 0 {
		return strList
	}
	stack.Push(strList[0])
	for _, str := range strList[1:] {
		if str != stack.GetTop() {
			stack.Push(str)
		}
	}
	return stack.GetList()
}

func ShortBlanks(bytesList []byte) (rs []byte) {
	stack := sliceStack.SliceStack[byte]{}
	if len(bytesList) == 0 {
		return bytesList
	}
	same := false
	for _, b := range bytesList {
		if unicode.IsSpace(rune(b)) {
			if same {
				continue
			} else {
				same = true
				stack.Push(b)
			}
		} else {
			stack.Push(b)
			same = false
		}
	}
	return stack.GetList()
}

func CharCount() {
	counts := make(map[rune]int)
	numberMap := make(map[int]int)
	letterMap := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) && n == 1 {
			letterMap[r]++
		}
		if unicode.IsNumber(r) && n == 1 {
			numberMap[int(r)]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Println("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	for num, n := range numberMap {
		fmt.Printf("number:%c\tcount:%d\n", num, n)
	}
	for letter, n := range letterMap {
		fmt.Printf("letter:%c\t%d\n", letter, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func WordFreq() {
	if len(os.Args) < 3 {
		fmt.Printf("请输入文件名 及 正则表达式")
		return
	}
	fileNumber, _ := strconv.Atoi(os.Args[1])
	charMap := make(map[string]int)

	for i := 0; i < fileNumber; i++ {
		filepath := os.Args[i+2]
		file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("Open file error!", err)
			return
		}

		_, err = file.Stat()
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			charMap[scanner.Text()]++
		}

		for char, n := range charMap {
			fmt.Printf("%s\t%d\n", char, n)
		}
	}

}
