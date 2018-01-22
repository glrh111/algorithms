package main

import (
	"fmt"
)

/*
    生成排列 不含重复元素。
 */
func permute(nums []int) [][]int {
	var (
		re = [][]int{}
		l = len(nums)
	)
	for x := range permutionGenerator(l) {
		newPer := make([]int, l)
		for i, index := range x {
			newPer[i] = nums[index]
		}
		re = append(re, newPer)
	}
	return re
}

/*
   排列生成器
 */
func permutionGenerator(n int) chan []int {

	var (
		c = make(chan []int)
		currentPer = make([]int, n)
		nextPer = make([]int, n)
	)

	for i:=0; i<n; i++ {
		currentPer[i] = i
	}

	go func() {
		newPer := make([]int, n)
		copy(newPer, currentPer)
		c <- newPer

		for {

			var (
				i, h int
			)

			// S1 求满足关系式 p_(j-1) < p_j 的最大值，设为 i，即 i = max{j | p_(j-1) < p_j}
			for j:=n-1; j>0; j-- {
				if currentPer[j-1] < currentPer[j] {
					i = j
					break
				}
			}
			if i == 0 {
				close(c)
				return
			}

			// S2 求满足关系式 p_(i-1) < p_k 的k的最大值，设为 h, 即 h = max{k | p_(i-1) < p_k}
			for k:=n-1; k>=0; k-- {
				if currentPer[i-1] < currentPer[k] {
					h = k
					break
				}
			}

			// S3 p_(i-1)与p_h 互换，得到 p1'p2'...pn'
			currentPer[i-1], currentPer[h] = currentPer[h], currentPer[i-1]

			// S4 令 p1'p2'...p_(i-1)'p_i'p_(i+1)'...p_n 中 i ~ n 的顺序逆转，即得到下一个排列
			copy(nextPer, currentPer)
			for j:=i; j<n; j++ {
				nextPer[j] = currentPer[n+i-j-1]
			}

			// copy
			newPer := make([]int, n)
			copy(newPer, nextPer)
			c <- newPer

			copy(currentPer, nextPer)
		}
	}()

	return c
}

func main() {
	fmt.Println(permute([]int{123, 456}))
}