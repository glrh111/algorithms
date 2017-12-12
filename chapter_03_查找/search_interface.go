package main

import "strings"

type ComparableInterface interface {
	CompareTo(b Comparable) int
}

type ComparableInt struct {
	value int
}

/*
    比较类型的整数实现
 */

// 类型断言参考这里
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.6.md
func (this ComparableInt) compareTo(b ComparableInt) (re int) {
	if (this.value > b.value) {
		re = 1
	} else if (this.value == b.value) {
		re = 0
	} else {
		re = -1
	}
	return
}

/*
   比较类型的字符串实现
 */

type ComparableString struct {
	value string
}

func (this ComparableString) compareTo(b ComparableString) (re int) {
	re = strings.Compare(this.value, b.value)
	return
}

/*
   下面是ST的 interface
 */

type Key ComparableInterface

type SymbolTableInterface interface {
	// 更新操作
	put(key Key, value interface{})
	// 查找操作
	get(key Key) interface{}
	// 删除键值対
	delete(key Key)
	// 键值対个数
	size() int
	// 是否包含某个值
	contains(key Key) bool
	// 是否为空
	isEmpty() bool
	// 所有键的集合
	keys() []Key
}

type SortedSymbolTableInterface interface {
	SymbolTableInterface
	// 最小的键
	min() (Key, error)
	// 最大的键
	max() (Key, error)
	// 小于等于key的最大键
	floor() (Key, error)
	// 大于等于key的最小键
	ceiling() (Key, error)
	// 排名为k的键
	selectKey(index int) (Key, error)
	// 键k的排名
	rank(key Key) int
	// 删除最小值 表为空时返回错误
	deleteMin() error
	// 删除最大值
	deleteMax() error
	// lo, hi 之间键的数量
	sizeBetween(lo int, hi int) int
	// lo, hi 之间的所有键
	keysBetween(lo int, hi int) []Key
}

