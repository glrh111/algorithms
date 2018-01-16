package main

import "fmt"

/*
   选择排序
   每次找到剩余数组的最小元素，与头部交换。每次排定一个最小的元素.
 */
func SelectionSort(data Interface) {
	selectionSort(data, 0, data.Len()-1)
}

func selectionSort(data Interface, lo int, hi int) {
	if lo >= hi {
		return
	}
	minIndex := lo
	for i:=lo+1; i<=hi; i++ {
		if data.Less(i, minIndex) {
			minIndex = i
		}
	}
	data.Swap(lo, minIndex)
	selectionSort(data, lo+1, hi)
}

/*
   冒泡排序
   每次找到最大的元素，放到末尾
 */
func BubbleSort(data Interface) {
	bubbleSort(data, 0, data.Len()-1)
}

func bubbleSort(data Interface, lo int, hi int) {
	if lo >= hi { return }

	for i:=lo; i<=hi-1; i++ {
		if data.Less(i+1, i) {
			data.Swap(i+1, i)
		}
	}
	bubbleSort(data, lo, hi-1)
}

func main() {
	p := generateIntSlice(100, 10)
	fmt.Println("选择排序前", p)
	SelectionSort(p)
	fmt.Println(p, p.IsSorted())

	p.Shuffle()
	fmt.Println("冒泡排序前：", p)
	BubbleSort(p)
	fmt.Println(p, p.IsSorted())
}

