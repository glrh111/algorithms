package main

import (
	"math"
	"fmt"
)

/*
   是否是回文数字
   Palindrome Number

   将数字存入数组里边吗？ good，这个可以啊

   TODO 性能有待改进 只超过了 12 % 的 golang 用户
 */

func isPalindrome(x int) bool {
	var (
		raw = []int{}
		rev = []int{}
		is = true
		i = 1
	)
	if x < 0 {
		return false
	}

	for {
		base10i := int(math.Pow10(i))
		yushu := (x % base10i)
		yushu2 := yushu / int(math.Pow10(i-1))
		rev = append(rev, yushu2)
		raw = append([]int{yushu2}, raw...)
		i++
		if yushu == x {
			break
		}
	}

	for j:=0; j<len(raw); j++ {
		if raw[j] != rev[j] {
			is = false
			break
		}
	}

	return is
}

func main() {
	fmt.Println(isPalindrome(-1))
}
