package algorithm

func BlankReplace(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			s = s[:i] + "%20" + s[i+1:]
		}
	}
	return s
}
