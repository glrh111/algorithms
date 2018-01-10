package main

import "fmt"

/*
  规则如下：I(1) V(5) X(10) L(50) C(100) D(500) M(1000)
  https://zh.wikipedia.org/wiki/%E7%BD%97%E9%A9%AC%E6%95%B0%E5%AD%97
  + 重复次数：一个Roman数字重复几次，就表示这个数字的几倍
  + 右加左减：
    + 可以放在左边表示减的数字: I X C
    + 左右相对与较大的数字而言
    + 左减必须为 1 位
    + 右加数字不可连续超过3位
  + 加线乘千：用不到吧
  + 数码限制：同一个数码最多连续出现3次

  FIXME 这个实现慢的一笔，超过了 0% 的提交者。FUCK 3999 个用例，跑了 96 ms

  第二次提交相同的代码，用的时间是一半！ 先不管了，擦
 */
func intToRoman(num int) string {
	var (
		roman = ""
		alphabet = [][]string{
			[]string{"I", "V"},
			[]string{"X", "L"},
			[]string{"C", "D"},
			[]string{"M", "0"}, // 最后一个永不到，因为最大的数是 3999
		}
		divisor = 1000
		divident = 0
		appendStr = ""
	)

	// num 最高 4 位
	for i:=3; i>=0; i-- {

		divident = num / divisor % 10 // 最高位

		if 1 <= divident && divident <= 3 { // 重复 1 到 3 次
			for j:=0; j<divident; j++ {
				appendStr += alphabet[i][0]
			}
		} else if 4 == divident {           // 左边减去 1
			appendStr = alphabet[i][0] + alphabet[i][1] // 右边的数字，代表 5
		} else if 5 <= divident && divident <= 8 {           // 右加
			appendStr = alphabet[i][1]
			for j:=0; j<divident-5; j++ {
				appendStr += alphabet[i][0]
			}
		} else if 9 == divident {
			appendStr = alphabet[i][0] + alphabet[i+1][0]
		}

		roman += appendStr
		appendStr = ""
		divisor /= 10
	}


	return roman
}

/*
   将 Roman 转化为整数
   有什么规律？
   从左到右以此扫描，记忆上一个数字
 */
func romanToInt(s string) int {
	var (
		alphabet = map[string]int{
			"I": 1, "V": 5,
			"X": 10, "L": 50,
			"C": 100, "D": 500,
			"M": 1000, "": 0,
		}
		lastLetter = "" // 上一个字母
		lastint = 0
		currentLetter = ""
		romanint = 0
		currentint = 0
	)

	for i:=0; i<len(s); i++ {
		currentLetter = string(s[i])
		currentint = alphabet[currentLetter]
		lastint = alphabet[lastLetter]

		if currentint > lastint { // romanint - last + current
			romanint += -2 * lastint + currentint // 减去双倍的。本来就该减去，结果还多加了。
		} else {
			romanint += currentint
		}

		lastLetter = currentLetter
	}

	return romanint
}

func main() {
	fmt.Println(intToRoman(1234), intToRoman(600), romanToInt(intToRoman(1234)))
}
