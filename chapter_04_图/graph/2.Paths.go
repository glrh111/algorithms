package graph

type Paths struct {
	source int          // 初始点
	graph *Graph        // 传入的图
	markedList []bool   // cap=graph.V() 状态表示是否连通
	edgeTo []int    // 里边存入到index 的前一个 顶点
}

func NewPaths(graph *Graph, s int) (p *Paths) {

	p = &Paths{
		source:s,
		graph:graph,
		markedList: make([]bool, graph.V()),
		edgeTo: make([]int, graph.V()),
	}

	p.dfs(p.source)

	return
}

// 参考了书中的实现
func (p *Paths) dfs(s int) {
	p.markedList[s] = true
	// 存入路径
	for v := range p.graph.Adj(s) {
		if ! p.markedList[v] {
			p.edgeTo[v] = s
			p.dfs(v)
		}
	}

}

func (p *Paths) HasPathTo(v int) bool {
	return bool(p.markedList[v])
}

// 路径
func (p *Paths) PathTo(v int) []int {

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

