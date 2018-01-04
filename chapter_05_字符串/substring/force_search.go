package substring

func forceSearch(s string, patt string) (index int) {
	N := len(s)
	M := len(patt)
	index = N // 未匹配上
	for i:=0; i<=N-M; i++ {
		j := 0
		for j=0; j<M; j++ {
			if s[i+j] != patt[j] { // 不匹配
				break
			}
		}
		if j == M { // 说明匹配上了
			index = i
			break
		}
	}
	return
}

// 显式回退
func forceSearch2(s string, patt string) (index int) {
	var (
		i, j = 0, 0
		N, M = len(s), len(patt)
	)
	for ; i<N && j<M; i++ {
		if s[i] == patt[j] {
			j++
		} else {
			i -= j
			j = 0
		}
	}
	if j == M {
		index = i - M
	} else {
		index = N
	}
	return
}