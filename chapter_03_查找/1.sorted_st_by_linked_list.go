package main

// 链表实现符号表



// 链表的节点
type linkedListNode struct {
	key Key
	value interface{}
	next *linkedListNode
}

// 构造新节点的方法
func NewLinkedListNode(key Key, value interface{}, next *linkedListNode) (node *linkedListNode) {
	node = &linkedListNode{
		key: key,
		value: value,
		next: next,
	}
	return
}

// 链表
type linkedListST struct {
	// 大小
	size int
	// 第一个元素
	first *linkedListNode
}

// 构造新的链表
func NewLinkedList() *linkedListST {
	return &linkedListST{
		size: 0,
		first: nil,
	}
}

// 查找方法
func (this *linkedListST) Get(key Key) {
	return
}

