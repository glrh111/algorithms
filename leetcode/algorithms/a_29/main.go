package main

import (
	"fmt"
)

/*
   用 Euclidean division 实现
   https://zh.wikipedia.org/wiki/%E5%B8%A6%E4%BD%99%E9%99%A4%E6%B3%95
   1. 二进制 + 二分法 。
   2. 长除法 得一个一个试，感觉比较麻烦。
 */
func divide(dividend int, divisor int) int {

	var (
		p, q, p2, q2, pvalue, qvalue uint64 // p 下界 q 上界
		result int
		udividend, udivisor uint64
		isnegtive = (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0)  // 正负数
		maxint32 = int(^uint32(0) >> 1) // 最大的 int32
		minint32 = int(-maxint32 - 1) // 最小的 int32
	)

	// 确定正负号
	if dividend < 0 {
		udividend = uint64(-dividend)
	} else {
		udividend = uint64(dividend)
	}
	if divisor < 0 {
		udivisor = uint64(-divisor)
	} else {
		udivisor = uint64(divisor)
	}

	// 1\ 先用递增的二进制法，找出大概区间
	q = 0
	for {
		if udividend < (udivisor << q) {
			p = q - 1
			break
		}
		q++
	}

	q2 = 1 << q // 2 ^ q
	p2 = 1 << p
	qvalue = udivisor << q
	pvalue = udivisor << p

	// 2\ 二分法 从 [p, q] 找出
	for {

		if qvalue - pvalue <= udivisor {
			if qvalue == udividend {
				result = int(q2)
			} else {
				result = int(p2)
			}
			break
		}

		mid := (p2 + q2) >> 1
		midvalue := (qvalue + pvalue) >> 1 // 除以 2
		if midvalue < udividend {
			p2 = mid
			pvalue = midvalue
		} else {
			q2 = mid
			qvalue = midvalue
		}
		//fmt.Println(p2, q2, pvalue, qvalue, mid, midvalue)
	}

	if isnegtive {
		result = -result
	}

	if result > maxint32  {
		result = maxint32
	} else if result < minint32 {
		result = minint32
	}

	return result

}

func main() {

	//for i:=0; i<1000; i++ {
	//	fmt.Println(i, divide(-i,11), -i/11)
	//}

	fmt.Println(divide(-2147483648, -1))
}
