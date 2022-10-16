package basename

import (
	"bytes"
	"strings"
	"unicode"
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
	buffer.WriteString(s[:read])
	if read > 0 {
		buffer.WriteByte(',')
	}
	for i, v := range s[read:] {
		if i%3 == 0 && i != 0 {
			buffer.WriteString(",")
		}
		buffer.WriteRune(v)
	}
	return buffer.String()
}

func NonRecursionBackDotComma(s string) string {
	if s == "" {
		return s
	}
	var buffer bytes.Buffer
	for i, v := range s {
		if i%3 == 0 && i != 0 {
			buffer.WriteString(",")
		}
		buffer.WriteRune(v)
	}
	return buffer.String()
}

func FloatComma(s string) string {
	if s == "" {
		return s
	}
	var buffer bytes.Buffer
	if !unicode.IsDigit(rune(s[0])) {
		buffer.WriteByte(s[0])
		s = s[1:]
	}
	spliterator := strings.Split(s, ".")
	if len(spliterator) != 2 {
		return s
	}
	integerPart := spliterator[0]
	decimalPart := spliterator[1]
	//integerPart
	buffer.WriteString(NonRecursionComma(integerPart))
	buffer.WriteByte('.')
	buffer.WriteString(NonRecursionBackDotComma(decimalPart))
	return buffer.String()
}

func Heterogeneous(l string, r string) bool {
	for len(l) > 0 || len(r) > 0 {
		v := l[0]
		if strings.Count(l, string(v)) == strings.Count(r, string(v)) {
			strings.ReplaceAll(l, string(v), "")
			strings.ReplaceAll(r, string(v), "")
		} else {
			return false
		}
	}
	if len(l) == 0 && len(r) == 0 {
		return true
	} else {
		return false
	}
}
