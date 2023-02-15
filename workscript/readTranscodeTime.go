package workscript

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

// regexpStr totalTranscodeTime:(\d+)
func regexpTranscoder(str string, regStr string) (num int) {
	re := regexp.MustCompile(regStr)
	match := re.FindStringSubmatch(str)
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
