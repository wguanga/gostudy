package tool

import "fmt"

//var lock sync.Mutex

//两数之和
func TwoSum(nums []int, target int) []int {
	flagg := map[int]int{}
	for i, v := range nums {
		if p, ok := flagg[target-v]; ok {
			return []int{p, i}
		}
		flagg[v] = i
	}
	return nil
}

//无重复最长子串

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func LengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	res, endl, j := 0, 0, 0
	for i := 0; i < len(s) && endl < len(s); {
		for _, ok := m[s[endl]]; !ok && endl < len(s); endl++ {
			m[s[endl]] = endl
			if endl+1 < len(s) {
				_, ok = m[s[endl+1]]
			}
		}
		res = max(res, endl-i)
		if endl < len(s) {
			i = m[s[endl]] + 1
			for ; j < i; j++ {
				delete(m, s[j])
			}

			m[s[endl]] = endl
			endl++
		}

	}
	return res
}

//寻找两个正序数组的中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ll := len(nums1) + len(nums2)
	if ll%2 == 1 {
		return float64(findminkforarr(nums1, nums2, ll/2+1))
	} else {
		return (float64(findminkforarr(nums1, nums2, ll/2)) + float64(findminkforarr(nums1, nums2, ll/2+1))) / 2
	}

}

func findminkforarr(nums1, nums2 []int, k int) int {
	l1, l2 := 0, 0
	for len(nums1) != 0 || len(nums2) != 0 {
		if l1 == len(nums1) {
			return nums2[l2+k-1]
		}
		if l2 == len(nums2) {
			return nums1[l1+k-1]
		}
		if k == 1 {
			return min(nums1[l1], nums2[l2])
		}
		half := k/2 - 1
		newl1, newl2 := min(l1+half, len(nums1)-1), min(l2+half, len(nums2)-1)
		if nums1[newl1] < nums2[newl2] {
			k -= newl1 - l1 + 1
			l1 = newl1 + 1
		} else {
			k -= newl2 - l2 + 1
			l2 = newl2 + 1
		}
	}
	return 0
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Suanfa1() {
	/* fmt.Println("hello")
	tool.Method()
	tool.Constant() */

	/* //两数之和
	nums := []int{3, 2, 4}
	target := 6
	fmt.Printf("TwoSum(nums, target): %v\n", TwoSum(nums, target)) */

	/* //无重复最长子串
	s := ""
	fmt.Printf("lengthOfLongestSubstring(s): %v\n", LengthOfLongestSubstring(s)) */

	//寻找两个有序数组中位数
	nums1, nums2 := []int{}, []int{1}
	fmt.Printf("findMedianSortedArrays(nums1, nums2): %v\n", FindMedianSortedArrays(nums1, nums2))
}
