package main

import "fmt"

type QuickSort struct {
	Sortable
}

func (this *QuickSort) Sort() {
	this.sort(0, len(this.valueList)-1)
}

// 需要在额外的列表实线。
func (this *QuickSort) sort(lo int, hi int) {
	if (lo <= hi) {
		return
	}
	j := this.partition(lo, hi)
	this.sort(lo, j-1)
	this.sort(j+1, hi)
}

// 分区用的
// 将 valueList[j] 放到一个合适的位置。左边的都小于它，右边都大于它。放到一个合适的位置。
func (this *QuickSort) partition(lo int, hi int) (j int) {
	i := lo
	j = hi
	for {
		for ; ! this.Less(this.valueList[lo], this.valueList[i]); i++ {

		}
		for ; this.Less(this.valueList[j], this.valueList[lo]); j-- {

		}
		if (i > j) {
			break
		}
		this.Exchange(i, j)
	}
	this.Exchange(lo, j)
	return
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
	sortInstance.Show()
	fmt.Println(sortInstance.IsSorted(), "\n")
}