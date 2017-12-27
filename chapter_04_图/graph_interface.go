package main

// BagInterface
// GraphInterface

/*
    NewBag()
 */
type BagInterface interface {
	Add(item *Key)     // 添加元素
	IsEmpty() bool     // 是否为空
	Size() int         // 元素数量
	Iterator() func() (int, bool)       // 元素列表. 每次读入一个元素. 还可以参考channel实现
	IteratorChan() chan int // chan 实现的iterator
}

/*
   广度优先搜索使用的FIFO队列. 用链表实现
 */
type FIFOQueueInterface interface {
	Push(item int)
	Pop() (int, bool)  // 弹出一个元素
}

/*
     NewGraph() 从文件中读取图
 */
type GraphInterface interface {
	V() int               // 顶点数
	E() int               // 边数
	AddEdge(v int, w int) // 增加边 v-w
	Adj(v int) chan int           // 和v相邻的所有顶点
	ToString() string           // 图的字符串表示
}

/*
    Search API NewSearch(graph *Graph, s int) 返回一个这样的类
    算法
 */
type SearchInterface interface {
	Marked(v int) bool // v, s 是否连通
	Count() int        // 与s连通的顶点的个数
}

/*
   Paths 搜索路径
 */
type PathsInterface interface {
	HasPathTo(v int) bool // 是否有到v的路径
	PathTo(v int) []int   // 到v的路径
}


