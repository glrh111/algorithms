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


type SymbolGraph struct {
	graph *Graph
	nameToIndex *LinearProbingHashST  // 还是使用符号表吧
	indexToName []string              // [index]name
}

/*
    从文件中读取图
    扫描第一遍： 
 */
func NewSymbolGraphFromFile(filename string) *Graph {
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

func (sg *SymbolGraph) Contains(name string) bool {
	return true
}

func (sg *SymbolGraph) Index(name string) int {
	retV := sg.nameToIndex.Get(NewComparable(name))
	if retV != nil {
		return retV.(int)
	}
	return -1 // 不允许返回 nil 蛋疼
}

func (sg *SymbolGraph) Name(v int) string {
	return sg.indexToName[v]
}

func (sg *SymbolGraph) G() *Graph {
	return sg.graph
}
