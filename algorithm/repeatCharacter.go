package algorithm

func RepeatCharacter(num []int) int {
	numMap := make(map[int]int)
	for _, k := range num {
		_, ok := numMap[k]
		if ok {
			return k
		} else {
			numMap[k] = 1
		}
	}
	return -1
}
