package main

import "fmt"

/*
   二分法查找，元素在数组中的位置
 */
func searchInsert(nums []int, target int) int {
	var (
		p, q = 0, len(nums)-1
		mid, targetIndex int
	)
	for {
		if nums[p] >= target {
			targetIndex = p
			break
		} else if nums[q] == target {
			targetIndex = q
			break
		} else if nums[q] < target {
			targetIndex = q+1
			break
		} else {
			mid = (p+q) / 2
			if nums[mid] > target {
				q = mid
				p++
			} else {
				p = mid
				q--
			}
		}
	}
	return targetIndex
}

func main() {
	nums := []int{1, 3, 4, 5, 7}
	fmt.Println(searchInsert(nums, 2))
}
