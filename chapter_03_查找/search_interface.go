package main

import (
	"strings"
)

type ComparableInterface interface {
	CompareTo(b Comparable) int
	Value() interface{}
}

type Comparable struct {
	value interface{}
}

func NewComparable(value interface{}) *Comparable {
	return &Comparable{value}
}

func (this *Comparable) CompareTo(b Comparable) (re int) {
	switch this.value.(type) {
	case int:
		thisValue := this.value.(int)
		thatValue := b.value.(int)
		if (thisValue > thatValue) {
			re = 1
		} else if (thisValue == thatValue) {
			re = 0
		} else {
			re = -1
		}
	case string:
		thisValue := this.value.(string)
		thatValue := b.value.(string)
		re = strings.Compare(thisValue, thatValue)
	default:
		panic("CompareTo only support int and string")
	}
	return
}

// 返回key的值
func (this *Comparable) Value() interface{} {
	return this.value
}

/*
   下面是ST的 interface
 */

 type SymbolTableInterface interface {
	// 更新操作
	Put(key *Comparable, value interface{})
	// 查找操作
	Get(key *Comparable) interface{}
	// 删除键值対
	Delete(key *Comparable)
	// 键值対个数
	Size() int
	// 是否包含某个值
	Contains(key *Comparable) bool
	// 是否为空
	IsEmpty() bool
	// 所有键的集合
	Keys() []*Comparable
}

type SortedSymbolTableInterface interface {
	SymbolTableInterface
	// 最小的键
	min() (*Comparable, error)
	// 最大的键
	max() (*Comparable, error)
	// 小于等于key的最大键
	floor() (*Comparable, error)
	// 大于等于key的最小键
	ceiling() (*Comparable, error)
	// 排名为k的键
	selectKey(index int) (*Comparable, error)
	// 键k的排名
	rank(key *Comparable) int
	// 删除最小值 表为空时返回错误
	deleteMin() error
	// 删除最大值
	deleteMax() error
	// lo, hi 之间键的数量
	sizeBetween(lo int, hi int) int
	// lo, hi 之间的所有键
	keysBetween(lo int, hi int) []*Comparable
}

