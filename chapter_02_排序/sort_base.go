package main

import "fmt"

/*
 *   Comparable
 */

type IComparable interface {
	CompareTo(a []IComparable )
}

type Comparable struct {
	value int
}

func (b *Comparable) CompareTo(a Comparable) int {
	if (b.value > a.value) {
		return 1
	} else if (b.value == a.value) {
		return 0
	} else {
		return -1
	}
}

func (b *Comparable) Value() int {
	return b.value
}

/*
 *   Sortable
 */

type ISortable interface {
	Sort()               // return
	Less(a Comparable, b Comparable) bool // a < b return true
	Exchange(a int, b int) // 交换位置
	Show() // 打印出来元素
	IsSorted() bool // 返回是否排序成功
}

type Sortable struct {
	valueList []Comparable
}

func (this *Sortable) Sort() {

}

func (this *Sortable) Less(a Comparable, b Comparable) bool {
	if (-1 == a.CompareTo(b)) {
		return true
	} else {
		return false
	}
}

func (this *Sortable) Exchange(a int, b int) {
	this.valueList[a], this.valueList[b] = this.valueList[b], this.valueList[a]
}

// 打乱元素
func (this *Sortable) Shuffle() {

}

func (this *Sortable) Show() {
	for _, value := range this.valueList {
		fmt.Printf("%d ", value.Value())
	}
	fmt.Print("\n")
}

func (this *Sortable) IsSorted() bool {
	if (len(this.valueList)>1) {
		for i:=0; i<len(this.valueList)-1; i++ {
			if this.Less(this.valueList[i+1], this.valueList[i]) {
				return false
			}
		}
	}
	return true
}





