package digraph

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


type Digraph struct {
	v int       // 顶点数量
	e int       // 边的数量
	adj []*Bag  // [v] = 与之相邻的顶点序号
}

/*
    从文件中读取图
 */
func NewDigraphFromFile(filename string) *Digraph {
	inputFile, inputError := os.Open(filename)
	if inputError != nil {
		panic("Open file error: " + inputError.Error())
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

	dg := &Digraph{
		v,
		e,
		make([]*Bag, v),
	}

	for i := 0; i < v; i++ {
		dg.adj[i] = NewBag(10)
	}

	for {
		inputString, readError = inputReader.ReadString('\n')
		// 去掉 \n
		inputString = strings.TrimSpace(inputString)

		if len(inputString) > 0 {
			eList := strings.Fields(inputString)
			// 第一个元素是 当前顶点，以后是与之相连的
			if len(eList) >= 2 {
				currentV, _ := strconv.Atoi(eList[0])
				for i := 1; i < len(eList); i++ {
					w, _ := strconv.Atoi(eList[i])
					dg.AddEdge(currentV, w)
				}
			}
		}

		if readError == io.EOF { //
			break
		}
	}
	return dg
}

/*
   构造一个空 Graph
 */
func NewDigraphWithSize(v int) (g *Digraph) {
	return &Digraph{
		v: v,
		e: 0,
		adj: make([]*Bag, v),
	}
}

func (g *Digraph) V() int {
	return g.v
}

func (g *Digraph) E() int {
	return g.e
}

// 增加一条边 bool 代表是否增加了一条边
func (g *Digraph) AddEdge(v int, w int) (ifAdd bool) {
	if g.adj[v] == nil {
		g.adj[v] = NewBag(10)
	}
	ifAdd = g.adj[v].Add(NewKey(w))
	return
}

// 一条边的所有顶点
func (g *Digraph) Adj(v int) chan int {
	return g.adj[v].IteratorChan()
}

// 字符串表示
func (g *Digraph) ToString() (s string) {

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
