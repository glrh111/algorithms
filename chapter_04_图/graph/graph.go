package graph

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
)

/* 邻接表实现的图

   adj : []Bag
       : 索引表示顶点序号v，从0开始
       : Bag 里的元素，是与v相邻的顶点的序号

 */


type Graph struct {
	v int       // 顶点数量
	e int       // 边的数量
	adj []*Bag  // [v] = 与之相邻的顶点序号
}

/*
    从文件中读取图
 */
func NewGraphFromFile(filename string) *Graph {
	inputFile, inputError := os.Open(filename)
	if inputError != nil {
		fmt.Println("Open file error: ", inputError.Error())
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	// 第一行是 v
	inputString, readError := inputReader.ReadString('\n')
	inputString = strings.Trim(inputString, "\n")
	v, _ := strconv.Atoi(inputString)

	// 第二行是 e
	inputString, readError = inputReader.ReadString('\n')
	inputString = strings.Trim(inputString, "\n")
	e, _ := strconv.Atoi(inputString)

	g := &Graph{
		v,
		e,
		make([]*Bag, v),
	}

	for i := 0; i < v; i++ {
		g.adj[i] = NewBag(10)
	}

	for {
		inputString, readError = inputReader.ReadString('\n')
		// 去掉 \n
		inputString = strings.Trim(inputString, "\n")

		if len(inputString) > 0 {
			eList := strings.Split(inputString, " ")
			// 第一个元素是 当前顶点，以后是与之相连的
			if len(eList) >= 2 {
				currentV, _ := strconv.Atoi(eList[0])
				for i := 1; i < len(eList); i++ {
					w, _ := strconv.Atoi(eList[i])
					g.AddEdge(currentV, w)
				}
			}
		}

		if readError == io.EOF { //
			break
		}
	}
	return g
}

/*
   构造一个空 Graph
 */
func NewGraphWithSize(v int) (g *Graph) {
	return &Graph{
		v: v,
		e: 0,
		adj: make([]*Bag, v),
	}
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

// 增加一条边
func (g *Graph) AddEdge(v int, w int) {
	for _, value := range []int{v,w} {
		if g.adj[value] == nil {
			g.adj[value] = NewBag(10)
		}
	}
	g.adj[v].Add(NewKey(w))
	g.adj[w].Add(NewKey(v))
}

// 一条边的所有顶点
func (g *Graph) Adj(v int) chan int {
	return g.adj[v].IteratorChan()
}

// 字符串表示
func (g *Graph) ToString() (s string) {

	// v , e
	s += "V: " + strconv.Itoa(g.v) + " E: " + strconv.Itoa(g.e) + "\n"

	// all edge
	for i := 0; i < g.v; i++ {
		s += fmt.Sprintf("%d: ", i)
		vChan := g.Adj(i)

		for v := range vChan {
			s += fmt.Sprintf("%d ", v)
		}
		s += "\n"
	}

	return
}

// v 的度数
func (g *Graph) Degree(v int) (degree int) {
	degree = 0

	for range g.Adj(v) {
		degree++
	}

	return
}

// 所有顶点的最大度数
func (g *Graph) MaxDegree(v int) (maxDegree int) {
	maxDegree = 0
	for i:=0; i<g.v; i++ {
		currentDegree := g.Degree(i)
		if currentDegree > maxDegree {
			maxDegree = currentDegree
		}
	}
	return
}

// 所有顶点的平均degree
func (g *Graph) AvgDegree(v int) int {
	return 2 * g.e / g.v
}

// 自环的个数
// TODO 实现与Page 334 不一样。为毛 n / 2 ?
func (g * Graph) NumberOfSelfLoops() (n int) {
	n = 0
	for v:=0; v<g.v; v++ {
		c := g.Adj(v)
		for w := range c {
			if w == v {
				n++
				//break // FIXME 这里 c 应该立马close, 但是规范不允许
			}
		}
	}
	return
}
