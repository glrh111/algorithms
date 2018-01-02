package string_sort

import (
	"fmt"
	"math/rand"
	"time"
)

/*
    低位优先排序

    性能：
       + 100W 日期 2449 ms
       + 1W   日期    1 ms

 */

// w 为字符串宽度
func LSDSort(a []string) {

	R := 256 // 字母表大小
	N := len(a)
	w := len(a[0])

	aux := make([]string, N) // 临时数组

	// w 次排序
	for i:=0; i<w; i++ {
		// 0/ 初始化
		count := make([]int, R+1)

		// 1/ 频率统计
		for j:=0; j<N; j++ {  // a[j][w-i-1] 为从右数，第w个字符，以此为key
			//fmt.Println("key", a[j], w-i-1, a[j][w-i-1]+1)
			count[a[j][w-i-1]+1]++ // count[key+1]++
		}

		// 2/ 频率转化为索引
		for j:=0; j<R; j++ {
			count[j+1] += count[j]
		}

		// 3/ 数据分类 数据移动到aux里边
		for j:=0; j<N; j++ {
			aux[ count[a[j][w-i-1]] ] = a[j]
			count[a[j][w-i-1]]++
		}

		// 4/ 将aux中的元素，复制进去 a[]
		copy(a, aux)

	}

}

func PrintArray(a []string) {
	for index, value := range a {
		fmt.Println(index, value)
	}
	fmt.Println()
}

// 生成一系列日期 生成n个
func GenerateDateList(n int) (dateList []string) {
	dateList = make([]string, n)
	for i:=0; i<n; i++ {

		// 随机生成年月日
		year := rand.Intn(1019) + 1000  // [1000, 1018]
		month := rand.Intn(12) + 1      // [1, 12]
		day := rand.Intn(28) + 1        // [1, 28]

		dateList[i] = fmt.Sprintf("%d-%02d-%02d", year, month, day)
	}
	return
}

// Timeit
func Timeit(a []string, sortFunc func (b []string)) {

	// 计时

	startNano := time.Now()

	sortFunc(a)

	escapeMilliseconds := time.Now().Sub(startNano).Nanoseconds() / 1000000 // ms

	//PrintArray(a)

	fmt.Printf("消耗时间: %d ms\n", escapeMilliseconds)


}