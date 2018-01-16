package main

import (
	"fmt"
	"math"
)

func merge(data IntSlice, lo int, mid int, hi int) {
	var (
		i, j = lo, mid+1
		aux = IntSlice(make([]int, hi+1-lo))
	)
	copy(aux, data[lo:hi+1])

	// 将元素复制一份出来
	for k:=lo; k<=hi; k++ {
		if i > mid { // 左边用完了, 取右边
			data[k] = aux[j-lo]
			j++
		} else if j > hi { // 右边用完了，取左边的
			data[k] = aux[i-lo]
			i++
		} else if aux.Less(i-lo, j-lo) { // 取 i
			data[k] = aux[i-lo]
			i++
		} else {
			data[k] = aux[j-lo]
			j++
		}
	}
	fmt.Println(lo, mid, hi, aux, data)
}

// 将元素氛围两部分排序 [lo, mid] (mid, hi]
func mergeSort(data IntSlice, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := (hi+lo) / 2
	mergeSort(data, lo, mid)
	mergeSort(data, mid+1, hi)
	merge(data, lo, mid, hi)
}

func MergeSort(data IntSlice) {
	mergeSort(data, 0, data.Len()-1)
}

/*
   自第向下 排序
 */
func MergeSortFromBottomToTop(data IntSlice) {
	l := data.Len()
	for sz:=1; sz<l; sz+=sz { // 1, 2, 4, 8
		for i:=0; i<l-sz; i+=2*sz {
			merge(data, i, i+sz-1, int(math.Min(float64(i+2*sz-1), float64(l-1))))
		}
	}
}

func main() {
	p := generateIntSlice(999, 10)
	//fmt.Println("归并排序前", p)
	//MergeSort(p)
	//fmt.Println("归并排序后：", p, p.IsSorted())

	p.Shuffle()
	MergeSortFromBottomToTop(p)
	fmt.Println(p, p.IsSorted())
}
