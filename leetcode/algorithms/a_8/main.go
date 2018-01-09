package main

import (
	"strconv"
	"fmt"
	"math"
	"strings"
)

/*
   可能碰到以下几种情况:
   + 越界：往int上累加 32 bit int 上下界都得判断
   + 非法字符: 直接panic
   + +/- 符号，只能出现在首位
   下面这些情况也需要处理：
   + "  -0012a42"

 */
func myAtoi(str string) int {
    str = strings.TrimSpace(str)
	var (
		l = len(str)
		value = 0
		negtiveflag = false
		scale = 1
		maxint32 = int(math.Pow(2, 31)) - 1
		minint32 = -maxint32-1
	)

	for i:=0; i<l; i++ {
		// 判断字符 [0, 9] [48, 57]
		// + 43 - 45
		ch := str[i]
		if (ch == 45 || ch == 43) && i == 0  {
			if ch == 45 {
				negtiveflag = true
			}
			continue
		}
		if ch < 48 || ch > 57 {
			scale = l - i + 1
			break
		}
		if ch != 48 {
			value += int(math.Pow10(l-1-i)) * int((ch - 48))
		}
	}
	if negtiveflag {
		value = -value
	}
	value = value / (int(math.Pow10(scale-1)))

	if !negtiveflag && value < 0 {
		value = maxint32
	}

	if value > maxint32 {
		value = maxint32
	}
	if value <= minint32 {
		value = minint32
	}

	return value
}

func main() {
	fmt.Println(strconv.Atoi("    010"))
	fmt.Println(myAtoi("-9223372036854775809"))
}
