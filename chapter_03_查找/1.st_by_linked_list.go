package main

import "fmt"

// 链表实现符号表


// 链表的节点
type linkedListNode struct {
	key *Comparable
	value interface{}
	next, prev *linkedListNode
}

// 构造新节点的方法
func NewLinkedListNode(
	key *Comparable,
	value interface{},
	next *linkedListNode,
	prev *linkedListNode,
) (node *linkedListNode) {

	// 判断key的类型
	// 怎么样，能使一个ST里边全是一个类型的键？

	node = &linkedListNode{
		key: key,
		value: value,
		next: next,
		prev: prev,
	}
	return
}

// 链表
type LinkedListST struct {
	// 大小
	size int
	// 第一个元素
	first *linkedListNode
}

// 构造新的链表
func NewLinkedList() *LinkedListST {
	return &LinkedListST{
		size: 0,
		first: nil,
	}
}

// 查找
func (this *LinkedListST) Get(key *Comparable) (value interface{}) {
	// 从 first 顺着node查找。
	value = nil
	for i := this.first; i != nil; i = i.next {
		if (i.key.CompareTo(*key) == 0) {
			value = i.value
			break
		}
	}
	return
}

// 设置新值
func (this *LinkedListST) Put(key *Comparable, value interface{}) {
	// 设置新值。先查找一圈，如果没有，再设置新值. 设置新节点
	for i := this.first; i != nil; i = i.next {
		if (i.key.CompareTo(*key) == 0) {
			i.value = value
			return
		}
	}
	// 构造一个新节点, 加到最前面
	newNode := NewLinkedListNode(
		key,
		value,
		this.first,
		nil,
	)
	if (this.first != nil) {
		this.first.prev = newNode
	}
	this.first = newNode
	this.size += 1
}

// 删除 节点的上一个节点指向下一个节点，把它空出来
// 如果某个值不存在，不报错。
func (this *LinkedListST) Delete(key *Comparable) {
	// 先不实现。
	for i := this.first; i != nil; i = i.next {
		if (i.key.CompareTo(*key) == 0) {
			if (i.prev == nil) { // 这个元素是第一个元素
				this.first = i.next
			} else {
				i.prev.next = i.next // 该元素不是第一个
			}
			this.size -= 1
			break
		}
	}
}

// 元素数量
func (this *LinkedListST) Size() (int) {
	return this.size
}

// 是否包含某个值
func (this *LinkedListST) Contains(key *Comparable) (re bool) {
	re = false
	for i := this.first; i != nil; i = i.next {
		if (i.key.CompareTo(*key) == 0) {
			re = true
			break
		}
	}
	return
}

// 是否为空
func (this *LinkedListST) IsEmpty() (re bool) {
	re = this.size == 0
	return
}

// key 列表
func (this *LinkedListST) Keys() (keys []*Comparable) {
	for i := this.first; i != nil; i = i.next {
		fmt.Println("key ", i.key, " value ", i.value, i.key.Value())
		keys = append(keys, i.key)
	}
	return
}

