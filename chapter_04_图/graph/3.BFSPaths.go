package graph

/*
   广度优先 paths
 */

type BFSPaths struct {
	source int          // 初始点
	graph *Graph        // 传入的图
	markedList []bool   // cap=graph.V() 状态表示是否连通
	edgeTo []int    // 里边存入到index 的前一个 顶点
}

func NewBFSPaths(graph *Graph, s int) (p *BFSPaths) {

	p = &BFSPaths{
		source:s,
		graph:graph,
		markedList: make([]bool, graph.V()),
		edgeTo: make([]int, graph.V()),
	}

	p.bfs()

	return
}

// 广度优先
func (p *BFSPaths) bfs() {

	// 维护一个 FIFOQueue
	var queue FIFOQueueInterface = NewFIFOQueue()

	queue.Enqueue(p.source)
	p.markedList[p.source] = true

	for v, ok := queue.Dequeue(); ok; v, ok = queue.Dequeue(){
		for w := range p.graph.Adj(v) {
			if !p.markedList[w] {
				p.markedList[w] = true
				queue.Enqueue(w)
				p.edgeTo[w] = v
			}
		}
	}
}

func (p *BFSPaths) HasPathTo(v int) bool {
	return bool(p.markedList[v])
}

// 路径
func (p *BFSPaths) PathTo(v int) []int {

	if ! p.HasPathTo(v) {
		return nil
	}

	path := []int{}
	for i:=v; i!=p.source; i=p.edgeTo[i] {
		path = append([]int{i}, path...)
	}
	path = append([]int{p.source}, path...)

	return path
}

