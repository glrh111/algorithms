package compress

/*
   实现优先队列. 双向链表实现 PriorityQueue()

   node(1) <-> node(2) <-> node(3) <-> node(4)
   first                               last

   ↑ DequeueMin                        ↑ DequeueMax

 */

type PriorityQueueNode struct {
	freq int                      // 根据这个排序
	prev, next *PriorityQueueNode // 前一个，后一个 结点
	value interface{}             // 里边保存的内容。对于实际应用，可以和一个数组配合使用，以支持复杂的数据结构
}

func NewPriorityQueueNode(freq int, value interface{}) (node *PriorityQueueNode) {
	return &PriorityQueueNode{
		freq: freq,
		value: value,
		prev: nil,
		next: nil,
	}
}

func (node *PriorityQueueNode) Value() interface{} {
	return node.value
}

type PriorityQueue struct {
	first, last *PriorityQueueNode
	size int
}

func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{
		first: nil,
		last: nil,
		size: 0,
	}
}

func (pq *PriorityQueue) Enqueue(freq int, value interface{}) {
	pq.size++
	node := NewPriorityQueueNode(freq, value)
	// 从first方向插入
	if pq.first == nil {
		// 直接将之当作first，last
		pq.first, pq.last = node, node
	} else {
		for i:=pq.first; ; i=i.next {

			if i == nil { // 遍历到 last 了
				node.prev = pq.last
				pq.last.next = node
				pq.last = node
				break
			}

			// 做比较
			if i.freq >= node.freq { // 插入到这个结点所在的位置.

				node.prev = i.prev
				node.next = i

				if pq.first == i {
					pq.first = node
				} else {
					node.prev.next = node
				}

				// i == nil 的情况也需要特殊处理下
				i.prev = node
				break
			}
		}
	}
}

func (pq *PriorityQueue) DequeueMax() (freq int, value interface{}, ok bool) {
	if pq.Size() > 0 {
		pq.size--
		ok = true
		node := pq.last
		pq.last = pq.last.prev
		if pq.last == nil {
			pq.first = nil
		}
		freq, value = node.freq, node.value
	} else {
		ok = false
	}
	return
}

func (pq *PriorityQueue) DequeueMin() (freq int, value interface{}, ok bool) {
	if pq.Size() > 0 {
		pq.size--
		ok = true
		node := pq.first
		pq.first = pq.first.next
		if pq.first == nil {
			pq.last = nil
		}
		freq, value = node.freq, node.value
	} else {
		ok = false
	}
	return
}

func (pq *PriorityQueue) Size() int {
	return pq.size
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Size() == 0
}