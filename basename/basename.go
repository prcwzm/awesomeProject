package basename

import (
	"bytes"
	"strings"
)

func EasyBasename(s string) (sr string) {
	sr = ""
	if len(s) == 0 {
		return
	}
	var i int
	for i = len(s) - 1; i >= 1; i-- {
		if s[i] == '.' {
			sr = s[:i]
			break
		}
	}
	for ; i >= 0; i-- {
		if s[i] == '/' {
			sr = sr[i+1:]
			break
		}
	}
	return
}

func PromoteBasename(s string) (sr string) {
	slash := strings.LastIndex(s, "/")
	sr = s[slash+1:]
	if dot := strings.LastIndex(sr, "."); dot != -1 {
		sr = sr[:dot]
	}
	return
}

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

func NonRecursionComma(s string) string {
	var buffer bytes.Buffer
	read := len(s) % 3
	buffer.WriteString(s[:read-1])
	for i, v := range s[read:] {
		if i%3 == 0 {
			buffer.WriteString(",")
		}
		buffer.WriteRune(v)
	}
	return buffer.String()
}
