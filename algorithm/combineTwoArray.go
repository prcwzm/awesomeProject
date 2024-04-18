package algorithm

func CombineTwoArrays(a []int, b []int, m int, n int) []int {
	i, j, k := n-1, m-1, m+n-1
	for i >= 0 && j >= 0 {
		if a[i] > b[j] {
			a[k] = a[i]
			i--
		} else {
			a[k] = b[j]
			j--
		}
		k--
	}

	if j >= 0 {
		a[k] = a[j]
		j--
		k--
	}
	return a
}
