package algorithm

func Kuaipai(arry []int, left int, right int) {
	if left >= right {
		return
	} else {
		jug := (left + right) / 2
		i, j := left, right
		mid := arry[jug]
		for i < j {
			for ; i < jug && arry[i] <= mid; i++ {
			}
			if arry[i] > mid {
				arry[jug] = arry[i]
				jug = i
			}
			for ; j > jug && arry[j] > mid; j-- {
			}
			if arry[j] <= mid {
				arry[jug] = arry[j]
				jug = j
			}
			arry[jug] = mid
		}
		Kuaipai(arry, left, jug-1)
		Kuaipai(arry, jug+1, right)

	}
}

func Guibing(arry []int) func([]int) {
	return division(arry)
}

func division(arry []int) {

}
