package workscript

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
)

// regexpStr totalTranscodeTime:(\d+)
func regexpTranscoder(str string, regStr string) (num int) {
	re := regexp.MustCompile(regStr)
	match := re.FindStringSubmatch(str)
	if match == nil || len(match) == 0 {
		return
	}
	num, _ = strconv.Atoi(match[1])
	return
}

func ReadTranscodeTime() {
	if len(os.Args) < 2 {
		fmt.Printf("请输入文件名")
		return
	}

	filepath := os.Args[1]
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	_, err = file.Stat()
	if err != nil {
		panic(err)
	}

	var totalTime = 0
	var lineCount = 0

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read finished")
			break
		} else if err != nil {
			fmt.Println("read error: ", err)
			return
		}
		if line == "" {
			continue
		}
		totalTime = totalTime + regexpTranscoder(line, "totalTranscodeTime:(\\d+)")
		lineCount++
	}
	fmt.Println("average: ", totalTime/lineCount)
}

func ReadRegexp() {
	if len(os.Args) < 3 {
		fmt.Printf("请输入文件名 及 正则表达式")
		return
	}

	filepath := os.Args[1]
	regexpStr := os.Args[2]
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	_, err = file.Stat()
	if err != nil {
		panic(err)
	}

	var totalTime = 0
	var lineCount = 0

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read finished")
			break
		} else if err != nil {
			fmt.Println("read error: ", err)
			return
		}
		if line == "" {
			continue
		}
		totalTime = totalTime + regexpTranscoder(line, regexpStr)
		lineCount++
	}
	fmt.Println("average: ", totalTime/lineCount)
}

func ReadRegexpReadFiles() {
	if len(os.Args) < 3 {
		fmt.Printf("请输入文件名 及 正则表达式")
		return
	}
	regexpStr := os.Args[1]
	fileNumber, _ := strconv.Atoi(os.Args[2])
	var totalTime = 0
	var lineCount = 0

	for i := 0; i < fileNumber; i++ {
		filepath := os.Args[i+3]
		file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("Open file error!", err)
			return
		}

		_, err = file.Stat()
		if err != nil {
			panic(err)
		}

		buf := bufio.NewReader(file)
		for {
			line, err := buf.ReadString('\n')
			if err == io.EOF {
				fmt.Println("read finished")
				break
			} else if err != nil {
				fmt.Println("read error: ", err)
				return
			}
			if line == "" {
				continue
			}
			totalTime = totalTime + regexpTranscoder(line, regexpStr)
			lineCount++
		}
		err = file.Close()
		if err != nil {
			fmt.Println("close error: ", err)
			return
		}
	}
	fmt.Println("average: ", totalTime/lineCount)
}

func ReadRegexpReadCntFiles() {
	if len(os.Args) < 3 {
		fmt.Printf("请输入文件名 及 正则表达式")
		return
	}
	regexpStr := os.Args[1]
	fileNumber, _ := strconv.Atoi(os.Args[2])

	numSet := make(map[int]int)

	for i := 0; i < fileNumber; i++ {
		filepath := os.Args[i+3]
		file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("Open file error!", err)
			return
		}

		_, err = file.Stat()
		if err != nil {
			panic(err)
		}

		buf := bufio.NewReader(file)
		for {
			line, err := buf.ReadString('\n')
			if err == io.EOF {
				fmt.Println("read finished")
				break
			} else if err != nil {
				fmt.Println("read error: ", err)
				return
			}
			if line == "" {
				continue
			}
			numSet[regexpTranscoder(line, regexpStr)]++
		}
		err = file.Close()
		if err != nil {
			fmt.Println("close error: ", err)
			return
		}

		keys := make([]int, 0, len(numSet))
		for k := range numSet {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, k := range keys {
			fmt.Printf("time:%v cnt:%v \n", k, numSet[k])
		}
	}

}
