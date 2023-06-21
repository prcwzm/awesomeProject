package algorithm

/* 0 1 2 3 4 5 */

func FindLessElement(list []int) int {
	var indexBegin = 0
	var indexEnd = len(list) - 1
	for indexBegin < indexEnd {
		mid := indexBegin + (indexEnd-indexBegin)/2
		// 值在右边
		if list[indexEnd] < list[mid] {
			indexBegin = mid + 1
		} else if list[indexBegin] > list[mid] {
			indexEnd = mid
		} else {
			indexEnd--
		}
	}
	return indexBegin
}
