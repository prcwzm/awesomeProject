package algorithm

import (
	"fmt"
	"strconv"
)

func FindUniqueItem(list []int) {
	mid := findHigh(list)
	i := mid
	j := mid + 1
	count := 0
	for i > -1 || j < len(list) {
		if list[i] == list[j] {
			count++
			i--
			j++
		} else if list[i] > list[j] {
			i--
		} else {
			j++
		}
	}
	fmt.Printf(strconv.Itoa(count))
}

func findHigh(list []int) int {
	var i = 0
	var j = len(list) - 1
	for i < j {
		mid := i + (j-i)/2
		if list[mid] > list[mid+1] {
			//递减序列商
			j = mid
		} else {
			//递增
			i = mid + 1
		}
	}
	return max(list, i, j)
}

func max(list []int, a, b int) int {
	if list[a] > list[b] {
		return a
	} else {
		return b
	}
}
