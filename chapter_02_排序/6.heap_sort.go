package main

/*
   Heap 这个数据结构 用法如下：

	p := generateIntSlice(100, 10)
	fmt.Println(p)
	h := NewHeap()
	for _, i := range p {
		h.Insert(i)
		fmt.Println("Insert: ", i)
	}
	for value,ok := h.DelMax(); ok; value,ok = h.DelMax() {
		fmt.Print(value, " ")
	}
 */
type Heap struct {
	data IntSlice // 实际存储数据的 size = N + 1 第0个元素，不存储值
	size int       // 数组大小
}

func NewHeap() Heap { // 将第一个元素填入值
	return Heap{
		IntSlice([]int{0}), // 第一个元素用不到
		0,
	}
}

func (h *Heap) swim(k int) {
	for {
		if !(k > 1 && h.data.Less(k/2, k) ) { // 到顶
			break
		}
		h.data.Swap(k/2, k)
		k = k / 2
	}
}

func (h *Heap) sink(k int) {
	for {
		// l, 2 * k , 2 * k + 1 这三个一起比较
		//if k > h.Size()-1 {
		//	break
		//}

		var (
			child = 2*k
			nextk = child
		)

		if child > h.Size() {            // 两个都越界了
			break
		}

		if nextk+1 <= h.size && h.data.Less(nextk, nextk+1) {
			nextk++
		}

		// 查看需不需要交换
		if ! h.data.Less(k, nextk) {
			break
		}

		// 交换元素
		h.data.Swap(k, nextk)

		// 等于交换过去的元素
		k = nextk
	}
}

func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

func (h *Heap) Size() int {
	return h.size
}

// 删除最大的元素。将最后一个元素作为最大元素 sink
func (h *Heap) DelMax() (value int, ok bool) {
	if ! h.IsEmpty() {
		value = h.data[1]
		ok = true
		h.data[1] = h.data[h.size]
		h.data = h.data[:h.size]
		h.size--
		if !h.IsEmpty() { // sink
			h.sink(1)
		}
	}
	return
}

// 插入 插入到最后，然后 swim
func (h *Heap) Insert(value int) {
	h.data = append(h.data, value)
	h.size++

	if h.size > 1 {
		h.swim(h.size)
	}
}

/*
    堆排序的实现
 */
func HeapSort(data Interface) {
	var (
		l = data.Len() - 1 // 第一个元素不参与排序
	)
	// 1. 构造堆
	for i:=l/2-1; i>=0; i-- {
		sink(data, i, l)
	}

	// 2. 排序
	for {
		if l <= 0 {
			break
		}
		data.Swap(0, l)
		l--
		sink(data, 0, l)
	}

}
// 下沉操作
func sink(data Interface, k int, n int) {
	for {

		var (
			child = 2*(k+1) - 1
			nextk = child
		)

		if child > n {  // 越界
			break
		}

		if nextk+1 <= n && data.Less(nextk, nextk+1) {
			nextk++
		}

		// 查看需不需要交换
		if ! data.Less(k, nextk) {
			break
		}

		// 交换元素
		data.Swap(k, nextk)

		// 等于交换过去的元素
		k = nextk
	}
}

//func main() {
//	// heap
//	//p := generateIntSlice(100, 10)
//	//fmt.Println(p)
//	//h := NewHeap()
//	//for _, i := range p {
//	//	h.Insert(i)
//	//}
//	//for value,ok := h.DelMax(); ok; value,ok = h.DelMax() {
//	//	fmt.Print(value, " ")
//	//}
//	p := generateIntSlice(100, 10)
//	fmt.Println("排序前", p)
//	HeapSort(p)
//	fmt.Println(p, p.IsSorted())
//}
