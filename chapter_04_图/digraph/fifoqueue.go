package digraph

/*
   双向链表，实现queue
   first -> (Node-value-next) -> (Node-value-next) <- last
   Pop()                                              Push(item)
 */

type FIFONode struct {
	value int
	next *FIFONode
	prev *FIFONode
}

func NewFIFONode(value int, next *FIFONode, prev *FIFONode) *FIFONode {
	return &FIFONode{
		value,
		next,
		prev,
	}
}

func (node *FIFONode) Value() int {
	return node.value
}

type FIFOQueue struct {
	first *FIFONode
	last *FIFONode
	size int
}

func NewFIFOQueue() *FIFOQueue {
	return &FIFOQueue{nil, nil, 0}
}

func (q *FIFOQueue) Enqueue(item int) {
	newNode := NewFIFONode(item, nil, nil)
	if q.first == nil {          // 空表
		q.first = newNode
		q.last = newNode
	} else {
		// 还得修改上级元素
		q.last.next = newNode
		newNode.prev = q.last
		q.last = newNode
	}
	q.size++
}

func (q *FIFOQueue) Dequeue() (item int, ok bool) {
	if nil == q.first {
		ok = false
	} else {
		ok = true
		item = q.first.Value()
		q.first = q.first.next
		if q.first == nil { // 仅剩下一个元素, POP一个，啥都没了
			q.last = nil
		} else {
			q.first.prev = nil
		}
		q.size--
	}
	return
}

func (q *FIFOQueue) Size() int {
	return q.size
}