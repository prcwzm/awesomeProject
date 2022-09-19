package basename

import "strings"

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
