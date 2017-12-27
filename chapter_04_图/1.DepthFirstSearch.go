package main

type DepthFirstSearch struct {
	source int          // 初始点
	graph *Graph        // 传入的图
	markedList []bool   // cap=graph.V() 状态表示是否连通
	count int           // 所有与之连通的点
}

func NewDepthFirstSearch(graph *Graph, s int) (se *DepthFirstSearch) {

	se = &DepthFirstSearch{
		source:s,
		graph:graph,
		markedList:make([]bool, graph.V()),
		count:0,
	}

	se.dfs(se.source)

	return
}

// 标记可标记的点
func (se *DepthFirstSearch) dfs(s int) {
	se.count++
	se.markedList[s] = true
	for v := range se.graph.Adj(s) {
		if ! se.markedList[v] {
			se.dfs(v)
		}
	}
}

func (se *DepthFirstSearch) Marked(v int) bool {
	return bool(se.markedList[v])
}

func (se *DepthFirstSearch) Count() int {
	return se.count
}
