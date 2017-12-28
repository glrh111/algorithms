package graph

/*
    连通分量 将连通的点汇总到一起
 */

type CC struct {
	graph *Graph        // 传入的图

	markedList []bool   // cap=graph.V() 状态表示是否连通

	idToVList []*Bag      // 索引是连通分量的 id
	vToId []int          // v 在哪个连通分量里边
	count int           // 连通分量个数

}

func NewCC(graph *Graph) (cc *CC) {

	cc = &CC{
		graph:graph,
		markedList: make([]bool, graph.V()),

		idToVList: []*Bag{},
		vToId: make([]int, graph.V()),
		count: 0,
	}

	for v:=0; v<cc.graph.V(); v++ {
		cc.dfs(v, true)
	}

	return
}

// 深度优先. 为啥不用广度优先呢?
func (cc *CC) dfs(s int, ifNew bool) {

	if !cc.markedList[s] {
		if ifNew  {
			cc.idToVList = append(cc.idToVList, NewBag(10))
			cc.count++
		}
		cc.markedList[s] = true
		currentId := len(cc.idToVList) - 1
		cc.idToVList[currentId].Add(NewKey(s))
		cc.vToId[s] = currentId

		for v := range cc.graph.Adj(s) {
			if !cc.markedList[v] {
				cc.dfs(v, false)
			}
		}
	}
}

// connected (v, w) v, w 是否连通
func (cc *CC) Connected(v int, w int) bool {
	return cc.Id(v) == cc.Id(w)
}

// count 连通分量个数
func (cc *CC) Count() int {
	return cc.count
}

// id(v) v 在哪个连通分量
func (cc *CC) Id(v int) int {
	return cc.vToId[v]
}

// IdElemList
func (cc *CC) IteratorChan(id int) chan int {
	return cc.idToVList[id].IteratorChan()
}



