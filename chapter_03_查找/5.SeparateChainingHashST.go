package main

import (
	"hash/fnv"
	"strconv"
)

// 拉链法ST，使用BST作为数组的值
// 怎么计算key的hash值？

/*
			   性能测试如下(M=97) 单位 ms
	 元素使用数据结构  | 读取tale(1) | 读取leipzig1M(1)
		---         | ---       |  ---
	 BST            | 146       | 43308
	 BinarySearchST | 168       | 93932

 */


type SeparateChainingHashST struct {
	stList []SymbolTableInterface
	m int // stList 的容量
	size int
}

// m 最好是一个素数
func NewSeparateChainingHashST(m int) *SeparateChainingHashST {
	stList := make([]SymbolTableInterface, m)
	for i := 0; i < len(stList); i++ {
		//stList[i] = NewBST()
		stList[i] = NewBinarySearchST()
	}
	return &SeparateChainingHashST{
		stList: stList,
		m: m,
		size: 0,
	}
}

// 这里的特殊处理下。
// 从下面ST返回的数据，PUT，DELETE操作等，需要返回是否删除或者增加了数据. 这里先不做修改
func (this *SeparateChainingHashST) Size() (size int) {
	return this.size
}

// 返回某个键的 hashCode 对应到 [0, m-1] 上。
func (this *SeparateChainingHashST) hashIndex(key *Comparable) (index int) {
	repString := ""
	switch key.Value().(type) {
	case int:
		repString = strconv.Itoa(key.Value().(int))
	case string:
		repString = key.Value().(string)
	default:
		panic("type not support")
	}
	h := fnv.New32()
	h.Write([]byte(repString))
	index = int(h.Sum32() % uint32(this.m))
	//fmt.Println("repString: ", repString, " ", code)
	return
}

// GET 找出value
func (this *SeparateChainingHashST) Get(key *Comparable) (value interface{}) {
	// 1. 计算key对应的hashCode
	// 2. 从hashCode对应的索引里边找
	return this.stList[this.hashIndex(key)].Get(key)
}

// PUT 更新值
func (this *SeparateChainingHashST) Put(key *Comparable, value interface{}) {
	// FIXME size 实现的优雅方法 这段代码影响性能
	if !this.Contains(key) {
		this.size++
	}
	this.stList[this.hashIndex(key)].Put(key, value)
}

// DELETE 删除某个值
func (this *SeparateChainingHashST) Delete(key *Comparable) {
	// FIXME 找一个size实现的优雅方法 这段代码影响性能
	if this.Contains(key) {
		this.size--
	}
	this.stList[this.hashIndex(key)].Delete(key)
}

// Contains 是否包含某个值
func (this *SeparateChainingHashST) Contains(key *Comparable) bool {
	return this.stList[this.hashIndex(key)].Contains(key)
}

// IsEmpty bool
func (this *SeparateChainingHashST) IsEmpty() bool {
	return this.size == 0
}

// Keys 所有键的集合
func (this *SeparateChainingHashST) Keys() (keys []*Comparable) {
	keys = []*Comparable{}
	for i := 0; i < this.m; i++ {
		keys = append(keys, this.stList[i].Keys()...)
	}
	return
}

// 读取用例测试
func readAndCountBySeparateChainingHashST(filename string, lengthThreshold int) (totalWordCount int, differendWordCount int) {
	return ReadAndCount(
		NewSeparateChainingHashST(97),
		filename,
		lengthThreshold,
	)
}

