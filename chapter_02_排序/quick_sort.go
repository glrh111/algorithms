package main

import "fmt"

type QuickSort struct {
	Sortable
}

func (this *QuickSort) Sort() {

}

// 需要在额外的列表实线。
func (this *QuickSort) sort(valueList []Comparable, ) {
	mid := this.Partition(0, len(this.valueList))

}

// 分区用的
func (this *QuickSort) Partition(lo int, hi int) int {

}

// 创建新对象用的
func NewQuickSortInstance(a []int) *QuickSort {
	var rawSortedList []Comparable
	for _, value := range a {
		rawSortedList = append(rawSortedList, Comparable{value})
	}
	return &QuickSort{Sortable{rawSortedList}}
}



func main() {
	sortInstance := NewQuickSortInstance([]int{5,4,3,1,1,8})
	sortInstance.Show()
	fmt.Print(sortInstance.IsSorted(), "\n")
	sortInstance.Sort()
	fmt.Println(sortInstance.IsSorted(), "\n")
}