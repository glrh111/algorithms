package main

import "fmt"

// 链表实现符号表


// 链表的节点
type linkedListNode struct {
	key Key
	value interface{}
	next *linkedListNode
}

// 构造新节点的方法
func NewLinkedListNode(key interface{}, value interface{}, next *linkedListNode) (node *linkedListNode) {

	// 判断key的类型
	// 怎么样，能使一个ST里边全是一个类型的键？

	node = &linkedListNode{
		key: &Comparable{key},
		value: value,
		next: next,
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
func (this *LinkedListST) Get(key Key) {
	// 从 first 顺着node查找。。
	
	return
}

// 设置新值
func (this *LinkedListST) Put(key Key) {

}

// 删除
func (this *LinkedListST) Delete(key Key) {

}

