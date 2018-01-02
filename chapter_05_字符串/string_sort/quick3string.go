package string_sort

/*
   三向字符串排序 看的我头疼
 */

func Quick3string(a []string) {
	quick3string(a, 0, len(a)-1, 0)
}

func quick3string(a []string, lo int, hi int, d int) {
	if lo >= hi {
		return
	}
	lt, gt := lo, hi
	v := charAt(a[lo], d)
	i := lo + 1

	for {
		if i > gt {
			break
		}
		t := charAt(a[i], d)
		if t < v {
			exch(a, lt, i)
			lt++
			i++
		} else if t > v {
			exch(a, i, gt)
			gt--
		} else {
			i++
		}
	}
	quick3string(a, lo, lt-1, d)
	if v >=0 {
		quick3string(a, lt, gt, d+1)
	}
	quick3string(a, gt+1, hi, d)
}

// 交换a b位置的元素
func exch(a []string, b int, c int) {
	a[b], a[c] = a[c], a[b]
}
