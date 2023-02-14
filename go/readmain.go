package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

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

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)
	var sum int = 0
	var count int = 0
	var max = 0
	var sumwithoutzero = 0
	var countzero = 0
	var sumtew = 0
	var counttew = 0

	var sumhum = 0
	var counthum = 0

	var sumthree = 0
	var countthree = 0

	var sumfive = 0
	var countfive = 0

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		list := strings.Split(line, ":")
		number, _ := strconv.Atoi(list[len(list)-1])
		if number > max {
			max = number
		}
		sum += number
		count++
		if number != 0 {
			sumwithoutzero += number
			countzero++
		}

		if 100 >= number && number > 20 {
			sumtew += number
			counttew++
		}

		if 300 >= number && number > 100 {
			sumhum += number
			counthum++
		}

		if 500 >= number && number > 300 {
			sumthree += number
			countthree++
		}

		if number > 500 {
			sumfive += number
			countfive++
		}
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}

	fmt.Printf("平均回调时间： %d , 帧数 %d \n", sum/count, count)

	fmt.Printf("平均回调时间（去除0）： %d , 帧数 %d \n", sumwithoutzero/countzero, countzero)

	fmt.Printf("平均回调时间（x > 20 && x <= 100）： %d , 帧数 %d \n", sumtew/counttew, counttew)

	fmt.Printf("平均回调时间（x > 100 && x <= 300）： %d , 帧数 %d \n", sumhum/counthum, counthum)

	fmt.Printf("平均回调时间（x > 300 && x <=500）： %d , 帧数 %d \n", sumthree/countthree, countthree)

	fmt.Printf("平均回调时间（大于500）： %d , 帧数 %d \n", sumfive/countfive, countfive)

	println("最大耗时 ", max)
}
