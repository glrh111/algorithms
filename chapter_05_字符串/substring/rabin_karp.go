package substring

import "fmt"

/*
   Rabin Karp 方法计算子串
   不明觉历。看得头疼。先不看了。

   算法的高效性，来自対指纹的高效计算和比较
 */

type RabinKarp struct {
	q int64
	r int64
}

func NewRabinKarp() (rp *RabinKarp) {
	return &RabinKarp{
		q: 999999999777777777, // 很大的一个素数
		r: 256,
	}
}

func RabinKarpSearch(s string, patt string) (index int) {
	return NewRabinKarp().Search(s, patt)
}

func (rk *RabinKarp) hashCode(s string, M int) (hc int64) {
	hc = 0
	for i:=0; i<M; i++ {
		hc = (hc * rk.r + int64(s[i])) % rk.q
	}
	return
}

// 蒙特卡利检测
func (rk *RabinKarp) check(s string, index int, M int) bool {
	return true
}

func (rk *RabinKarp) Search(s string, patt string) (index int) {

	var (
		N = len(s)
		M = len(patt)
		//RM = int(math.Pow(float64(R), float64(M-1))) % Q //
		RM int64 = 1
	)

	// 1/ 计算RM
	for i:=0; i<M; i++ {
		RM = (rk.r * RM) % rk.q
	}

	// 2/ 计算初始hash
	hc := rk.hashCode(s, M)
	patthc := rk.hashCode(patt, M)

	fmt.Println(hc, patthc)

	// 初始就匹配上去了
	if hc == patthc && rk.check(s, 0, M) {
		index = 0
		return
	}

	// 继续往下匹配
	nexthc := hc
	for i:=1; i<=N-M; i++ {
		//nexthc = ((nexthc + int64(s[i-1]) * (rk.q - RM)) * rk.r + int64(s[i+M-1])) % rk.q
		nexthc = (nexthc + rk.q - RM * int64(s[i-1]) % rk.q) % rk.q
		nexthc = (nexthc * rk.r + int64(s[i+M-1])) % rk.q
		if nexthc == patthc && rk.check(s, i, M) {
			index = i
			return
		}
	}
	index = N
	return
}
