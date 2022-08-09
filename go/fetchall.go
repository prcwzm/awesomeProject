package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	messageChan := make(chan string)
	for i := 0; i < 2; i++ {
		go fetchall(messageChan)
	}
	for i := 0; i < 2; i++ {
		fmt.Println(<-messageChan)
	}
	fmt.Printf("%2fs elapsed\n", time.Since(start).Seconds())
}

func fetchall(messageChan chan string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchAllIntoFile(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	messageChan <- fmt.Sprintf("%2fs elapsed\n", time.Since(start).Seconds())

}

func fetchAllIntoFile(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	var b bytes.Buffer
	nbytes, err := io.Copy(&b, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	writeFile(b, url, ch)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}

func writeFile(b bytes.Buffer, url string, ch chan string) bool {
	filename := "fileoutput.txt"
	var f *os.File
	var err error
	if checkFileIsExist(filename) {
		f, err = os.OpenFile(filename, os.O_APPEND, 0666)
		if err != nil {
			ch <- fmt.Sprintf("文件存在:%s", err)
			return false
		}
	} else {
		f, err = os.Create(filename)
		if err != nil {
			ch <- fmt.Sprintf("文件不存在:%s", err)
			return false
		}
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	n, _ := w.WriteString(b.String())
	fmt.Printf("写入 %d 个字节n", n)
	w.Flush()
	return true
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
