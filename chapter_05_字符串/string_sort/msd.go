package string_sort

/*
    高位优先排序 MSD
    从高位开始排序，非等宽字符串
 */

// 以 第d 个字符为键，対 a[lo] ~ a[hi] 之间的字符进行排序
func msdsort(a []string, lo int, hi int, d int) {

	if lo >= hi {
		return
	}

	R := 256 // 字母表大小
	N := len(a) // 数组长度

	aux := make([]string, N) // 临时数组

	// 0/ 初始化
	count := make([]int, R+2) // 加上了 -1 这个索引，所以大小为 R+2

	// 1/ 频率统计
	for j:=lo; j<=hi; j++ { // 第 w 个字符，来排序
		count[charAt(a[j], d)+2]++
	}

	// 2/ 频率转化为索引
	for j:=0; j<R+1; j++ {
		count[j+1] += count[j]
	}

	// 3/ 复制元素
	for j:=lo; j<=hi; j++ {
		aux[ count[charAt(a[j], d)+1] ] = a[j]
		count[charAt(a[j], d)+1]++
	}

	// 4/ 回写
	for j:=lo; j<=hi; j++ {
		a[j] = aux[j-lo]
	}

	// 5/ 继续递归
	for j:=0; j<R; j++ {
		msdsort(a, lo+count[j], lo+count[j+1]-1, d+1)
	}

}

func MSDSort(a []string) {
	msdsort(a, 0, len(a)-1, 0)
}

func charAt(s string, pos int) (index int) {
	if pos < len(s) {
		index = int(s[pos])
	} else {
		index = -1
	}
	return
}