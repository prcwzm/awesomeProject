package basename

func Basenameeasy(s string) (sr string) {
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
