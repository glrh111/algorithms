package main

import (
	"math"
	"fmt"
)

/*
   最长回文串。单个字母算不算回文？当然不算。
 */
func longestPalindrome(s string) string {
	var (
		ps = ""       // 最长的回文串
		pslength = 0  // 回文串的长度
		lo = 0
		hi = 0
		sindex = 0
		count = 0
	)

	for {

		if sindex >= len(s) {
			break
		}

		if hi >= len(s) {
			sindex++
			lo, hi = sindex, sindex + pslength - 1 // ?
			continue
		}

		isPs := scanString(s, lo, hi) // 这里需要增加一些判断
		count++
		fmt.Println("调用scanString：", count)

		if isPs { // 是回文
			if hi-lo+1 >= pslength {
				pslength = hi - lo + 1
				ps = s[lo:hi+1]
			}
		} else {  // 这里需要判断一下，是否还有价值继续往下扫描

		}

		hi++

	}


	return ps
}

/*
    是否是回文串
 */
func scanString(s string, lo int, hi int) (isPs bool) {
	loIndex := lo + int(math.Floor(float64(hi-lo+1)/2.0))
	hiIndex := lo + int(math.Ceil(float64(hi-lo+1)/2.0))

	reversedStr := ""
	rightStr := s[hiIndex:hi+1]
	for i:=0; i<len(rightStr); i++ {
		reversedStr += string(rightStr[len(rightStr)-1-i])
	}

	isPs = s[lo:loIndex] == reversedStr
	return
}

func main() {
	a := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	fmt.Println(longestPalindrome(a), len(a))
}

