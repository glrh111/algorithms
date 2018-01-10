package main

import (
	"fmt"
)

/*
   反转 32 bit 有符号整数
   如果溢出，返回 0
 */
func reverse(x int) int {
	var (
		xlength = 0 // 整数长度
		divisor = 1 // 被除的数
		revdivisor = 1 // 反向的
		retx = 0    // 返回值
		isnegtive = false // 是否是负数
		current = 0 // 当前位数字大小
		maxint32 = 2147483647
		minint32 = -2147483648
		overflow = false
	)

	if x < 0 {
		isnegtive = true
		x = -x
	}

	for {
		if x / divisor == 0 {
			break
		}
		divisor *= 10
		xlength++
	}

	divisor /= 10
	for i:=0; i<xlength; i++ {
		current = x / divisor % 10
		if current == 0 {
			divisor /= 10
			revdivisor *= 10
			continue
		}
		// 比较是否越界
		thisAddValue := current * revdivisor // FIXME 这里有问题。或许不能用 int32 装下
		if isnegtive {
			thisAddValue = -thisAddValue
			if retx < minint32 - thisAddValue {
				overflow = true
			}
		} else {
			if retx > maxint32 - thisAddValue {
				overflow = true
			}
		}

		if overflow {
			retx = 0
			break
		}
		retx += thisAddValue
		divisor /= 10
		revdivisor *= 10
	}
	return retx
}

func main() {
	fmt.Println(reverse(-2147483648))
}
