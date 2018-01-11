package main

import "fmt"

/*
   找到两个数组的中位数 median. 不用合并数组，找到中位数对应的索引，即可计算
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		leftindex, rightindex, leftvalue, rightvalue, index int
		leftflag, rightflag bool
		alllen = len(nums1) + len(nums2)

		j, k int    // 记录两个数组位置的指针
		minvalue int
	)

	if alllen == 0 {
		return 0.0
	}

	// 计算两个索引
	rightindex = alllen / 2
	if alllen % 2 == 0 { // 偶数
		leftindex = rightindex - 1
	} else {
		leftindex = rightindex
	}

	for {

		a, b := lstValue(nums1, j), lstValue(nums2, k)
		if a >= b { // choose b
			minvalue = b
			k++
		} else {
			minvalue = a
			j++
		}

		if leftindex == index {
			leftvalue = minvalue
			leftflag = true
		}
		if rightindex == index {
			rightvalue = minvalue
			rightflag = true
		}

		if leftflag && rightflag {
			break
		}

		index++
	}

	// 计算中位数
	return (float64(leftvalue) + float64(rightvalue)) / 2

}

func lstValue(lst []int, index int) int {
	if index >= len(lst) {
		return int(^uint(0) >> 1) // 最大的整数
	}
	return lst[index]
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,2,3,4}, []int{4,5,6}))
}
