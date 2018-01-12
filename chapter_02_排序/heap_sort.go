package main

type Heap struct {
	data Interface // 实际存储数据的 size = N + 1 第0个元素，不存储值
	size int       // 数组大小
}

func NewHeap() Heap { // 将第一个元素填入值
	return Heap{
		IntSlice([]int{0}), // 第一个元素用不到
		0,
	}
}

func (h Heap) swim(k int) {
	for {
		if !(k > 1 && h.data.Less(k/2, k) ) { // 到顶
			break
		}
		h.data.Swap(k/2, k)
		k = k / 2
	}
}

func (h Heap) sink(k int) {
	for {
		// l, 2 * k , 2 * k + 1 这三个一起比较
		//if k > h.Size()-1 {
		//	break
		//}

		var (
			leftk = 2*k
			rightk = leftk + 1
			nextk = 0
		)

		if leftk > h.Size() {            // 两个都越界了
			break
		} else if leftk == h.Size() {    // 右边的越界了
			nextk = leftk
		} else {                         // 都没越界
			// 找出两者的最大值, 还得防止越界
			if h.data.Less(leftk, rightk) {
				nextk = 2 * k + 1
			} else {
				nextk = 2 * k
			}

			if h.data.Less(k, nextk) {   // swap
				h.data.Swap(k, nextk)
			} else {
				break // 这个break 有问题 FIXME 回家了
			}
		}

		// 等于交换过去的元素
		k = nextk
	}
}

func (h Heap) IsEmpty() bool {
	return h.size == 0
}

func (h Heap) Size() int {
	return h.size
}
