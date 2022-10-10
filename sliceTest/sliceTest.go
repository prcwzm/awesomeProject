package sliceTest

func ReverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func MoveN(s []int, n int) {
	ReverseSlice(s[:n])
	ReverseSlice(s[n:])
	ReverseSlice(s)
}
