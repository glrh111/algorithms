package main

import "fmt"

/*
   快速排序，原地排序
 */
func quickSort(data Interface) {
	_quickSort(data, 0, data.Len()-1)
}

func _quickSort(data Interface, lo int, hi int) {
	if hi <= lo {
		return
	}
	j := partition(data, lo, hi)
	_quickSort(data, lo, j-1)
	_quickSort(data, j+1, hi)
}

/*
   切分总是能排定一个元素
 */
func partition(data Interface, lo int, hi int) int {

	var i, j = lo, hi + 1

	for {

		// i 往后扫描
		for {
			i++
			if data.Less(lo, i) { break }
			if i == hi {  break }
		}

		// j 往前扫描
		for {
			j--
			if data.Less(j, lo) { break }
			if j == lo { break }
		}

		if i >= j { break }

		data.Swap(i, j)
	}

	data.Swap(lo, j)

	return j
}

/*
    三向切分快速排序
 */
func quickSort3Way(data Interface) {
	_quickSort3Way(data, 0, data.Len()-1)
}

func _quickSort3Way(data Interface, lo int, hi int) {
	if lo >= hi { return }
	// 切分元素就是 lo
	var (
		lt, gt = lo+1, hi // [lt, gt] 之间的元素相等
		i = lo + 1 // i 为索引指针 lo 为切分元素。
	)

	for {
		if i > gt { break }
		cmp := data.Compare(i, lo)
		if cmp > 0 {  // i > v 跟gt交换。gt-- i 不变
			data.Swap(i, gt)
			gt--
		} else if cmp == 0 { // 元素留在本地
			i++
		} else {             //
			data.Swap(i, lt)
			lt++
			i++
		}
	}
	data.Swap(lo, lt-1)
	_quickSort3Way(data, lo, lt-1)
	_quickSort3Way(data, gt+1, hi)
}

func main() {
	p := generateIntSlice(1000000, 10)
	fmt.Println(p.IsSorted())

	quickSort3Way(p)
	fmt.Println(p.IsSorted())

	p.Shuffle()
	fmt.Println("After shuffle!")
	quickSort(p)
}

