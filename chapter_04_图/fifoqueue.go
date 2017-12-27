package main

// first -> (Node-value-next) -> (Node-value-next) -> nil
// Pop()                                              Push(item)

type FIFONode struct {
	value int
	next *FIFONode
}

func NewFIFONode(value int, next *FIFONode) *FIFONode {
	return &FIFONode{
		value,
		next,
	}
}

func (node *FIFONode) Value() int {
	return node.value
}

type FIFOQueue struct {
	first *FIFONode
}

func NewFIFOQueue() *FIFOQueue {
	return &FIFOQueue{nil}
}

func (q *FIFOQueue) Push(item int) {
	newNode := NewFIFONode(item, nil)
	if q.first == nil {
		q.first = newNode
	} else {
		for i:=q.first; i!=nil; i=i.next {
			if i.next == nil {
				i.next = newNode
			}
		}
	}
}

func (q *FIFOQueue) Pop() (item int, ok bool) {
	if nil == q.first {
		ok = true
	} else {
		item = q.first.Value()
		q.first = q.first.next
	}
	return
}