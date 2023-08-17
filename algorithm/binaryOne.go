package algorithm

func BinaryOne(num int) (res int) {
	res = 0
	for num != 0 {
		res++
		num = num & (num - 1)
	}
	return
}
