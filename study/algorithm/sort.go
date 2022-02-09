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

func Guibing(arry []int) {
	division(arry, 0, len(arry)-1)
}

func division(arry []int, left int, right int) {
	if right <= left {
		return
	} else {
		division(arry, left, (right+left)/2)
		division(arry, (left+right)/2+1, right)
		merge(arry, left, right)
	}
}

func merge(arry []int, left int, right int) {
	mid := (left+right)/2 + 1
	cop := []int{}
	i, j := left, mid
	for i < mid && j <= right {
		if arry[i] <= arry[j] {
			cop = append(cop, arry[i])
			i++
		} else {
			cop = append(cop, arry[j])
			j++
		}
	}
	if i >= mid {
		for j <= right {
			cop = append(cop, arry[j])
			j++
		}
	} else if j > right {
		for i < mid {
			cop = append(cop, arry[i])
			i++
		}
	}
	k := left
	for _, v := range cop {
		arry[k] = v
		k++
	}
}
