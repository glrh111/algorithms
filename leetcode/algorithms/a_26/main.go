package main

import (
	"fmt"
)

/*
   FIXME 性能堪忧
 */
func removeDuplicates(nums []int) int {
	var (
		lastint = 0
		l = 0
		i = 0
		j = 0
	)
	for ; i<len(nums); i++ {
		if i != 0 {
			if lastint == nums[j] {
				removeEle(&nums, j)
			} else {
				lastint = nums[j]
				l++
				j++
			}
		} else {
			lastint = nums[i]
			l++
			j++
		}
	}
	//if l > 0 {
	//	nums = nums[0:l+1]
	//}
	return l
}

func removeEle(nums *[]int, index int) {
	if index >= len(*nums) - 1 {
		*nums = (*nums)[0:index]
	} else {
		*nums = append((*nums)[0:index], (*nums)[index+1:]...)
	}
}

func main() {
	nums := []int{1,1,1,2,2,4,5,5,6}
	fmt.Println(removeDuplicates(nums), nums)
}
