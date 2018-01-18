package main

import "fmt"

/*
   还得考虑效率 怎么进行幂运算 ?
 */
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	var (
		re = x
		negflag = false
		remainder = []int{} // 余数列表
		//remain int // 每次的余数
		from = n

	)

	if n < 0 {
		n = -n
		negflag = true
		from = n
	}

	for {
		if from == 1 {
			break
		}
		remainder = append(remainder, from % 2)
		from = from / 2
	}

	lr := len(remainder)
	for i:=0; i<lr; i++ { // 从小到大
		re *= re
		if remainder[lr-1-i] == 1 { // 余数只可能是 1,或者 0
			re *= x
		}
	}

	if negflag {
		re = 1 / re
	}

	return re
}

func main() {
	fmt.Println(myPow(0.00001, 2147483647))
}
