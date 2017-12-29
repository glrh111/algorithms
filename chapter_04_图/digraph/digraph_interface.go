package digraph

// BagInterface
// GraphInterface

/*
    NewBag()
 */
type BagInterface interface {
	Add(item *Key) bool  // 添加元素
	IsEmpty() bool       // 是否为空
	Size() int           // 元素数量
	Iterator() func() (int, bool)       // 元素列表. 每次读入一个元素. 还可以参考channel实现
	IteratorChan() chan int // chan 实现的iterator
}

/*
     广度优先搜索使用的FIFO队列. 用链表实现
 */
type FIFOQueueInterface interface {
	Enqueue(item int)
	Dequeue() (int, bool)  // 弹出一个元素
	Size() int
}

/*
     NewDigraph() 从文件中读取有向图
 */
type GraphInterface interface {
	V() int               // 顶点数
	E() int               // 边数
	AddEdge(v int, w int) bool // 增加边 v-w
	Adj(v int) chan int           // 和v相邻的所有顶点
	ToString() string           // 图的字符串表示
	Reverse() *Digraph    // 反转
}





