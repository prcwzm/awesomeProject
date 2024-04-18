package algorithm

import "math"

func Jump(jumpList []float64) bool {
	res := float64(0)
	for i, _ := range jumpList {
		if float64(i) <= res {
			res = math.Max(res, float64(i)+jumpList[i])
			if int(res) > len(jumpList) {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

func JumpZero(jumpList []float64) bool {
	if len(jumpList) == 1 {
		return true
	}
	for k, v := range jumpList {
		if float64(v) != 0 {
			continue
		}
		isHasValid := false
		for i := k - 1; i >= 0; i-- {
			if jumpList[i]+float64(i) > float64(k) || jumpList[i]+float64(i) >= float64(len(jumpList)-1) {
				isHasValid = true
				break
			}
		}
		if !isHasValid {
			return false
		}
	}
	return true
}
