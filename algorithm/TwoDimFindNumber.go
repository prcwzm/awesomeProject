package algorithm

func TwoDimFIndNumber(matrix [][]int, target int) (x, y int) {
	horizontal := len(matrix)
	if horizontal == 0 {
		return -1, -1
	}
	vertical := len(matrix[0])
	if vertical == 0 {
		return -1, -1
	}
	i := 0
	j := vertical - 1
	for i < horizontal && j >= 0 {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return i, j
		}
	}
	return -1, -1
}
