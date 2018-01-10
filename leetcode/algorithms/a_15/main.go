package main

import (
	"fmt"
)

/*
   给一个列表，找出来所有的3个数加一起等于 0 的集合

   暴力法：尝试所有数组的组合 C^(3)_(len(nums))，对于每一个 排列，计算出和，如果满足，那么插入

   参考wiki：https://zh.wikipedia.org/wiki/%E7%B5%84%E5%90%88
 */
func threeSum(nums []int) [][]int {
	m := 3
	allSum3 := [][]int{}
	sli := make([]int, m)
	for i:=0; i<m; i++ {
		sli[i] = i
	}
	for com := range comGenerator(sli, len(nums), m) {
		// 计算和
		sum := 0
		thisSli := make([]int, m)
		for i:=0; i<m; i++ {
			sum += nums[com[i]]
			thisSli[i] = nums[com[i]]
		}
		if sum == 0 {
			allSum3 = append(allSum3, thisSli)
		}
	}
	return allSum3
}

/*
   一个排列算法
 */
func comGenerator(sli []int, n int, m int) (c chan []int) {
	// len(sli) = m
	c = make(chan []int)
	go func() {
		sliCopy := make([]int, m)
		copy(sliCopy, sli)
		c <- sliCopy

		for {

			i := m - 1
			e := n - m

			for {
				sli[i]++
				if sli[i] > e + i && i >= 1 {
					i--
				} else {
					break
				}
			}

			if sli[0] > e {
				close(c)     // 是这样写的呀
				break
			}

			for {
				i++
				if i >= m {
					break
				}
				sli[i] = sli[i-1] + 1
			}

			sliCopy = make([]int, m)
			copy(sliCopy, sli)
			c <- sliCopy
		}
		//close(c)

	}()

	return
}

func main() {
	count := 0
	for sli := range comGenerator([]int{0,1,2}, 9, 3) {
		fmt.Println(sli)
		count++
	}
	fmt.Println(count)
	//s := []int{1,2,3}
	//c := s
	//c[0] = 1000
	//fmt.Println(s, c)
}
