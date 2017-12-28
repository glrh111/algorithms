package graph

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
)

/*
    符号图 SymbolGraph

 */


type SymbolGraph struct {
	graph *Graph
	nameToIndex *LinearProbingHashST  // 还是使用符号表吧
	indexToName []string              // [index]name
}

/*
    从文件中读取图
    扫描第一遍：构建 nameToIndex , indexToName
    扫描第二遍：构建 graph

 */
func NewSymbolGraphFromFile(filename string, delimiter string) (sg *SymbolGraph) {

	sg = &SymbolGraph{
		graph: nil,
		nameToIndex: NewLinearProbingHashST(10),
		indexToName: []string{},
	}

	inputFile, inputError := os.Open(filename)
	if inputError != nil {
		fmt.Println("Open file error: ", inputError.Error())
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readError := inputReader.ReadString('\n')
		// 去掉 \n
		inputString = strings.Trim(inputString, "\n")

		if len(inputString) > 0 {

			for _, everyField := range strings.Split(inputString, delimiter) {
				if ! sg.nameToIndex.Contains(NewComparable(everyField)) {
					sg.nameToIndex.Put(NewComparable(everyField), len(sg.indexToName))
					sg.indexToName = append(sg.indexToName, everyField)
				}
			}
		}

		if readError == io.EOF { //
			break
		}
	}

	// 先不管 E
	inputReader = bufio.NewReader(inputFile)

	sg.graph = NewGraphWithSize(len(sg.indexToName))

	for {
		inputString, readError := inputReader.ReadString('\n')
		// 去掉 \n
		inputString = strings.Trim(inputString, "\n")

		if len(inputString) > 0 {

			// 所有里边的值，都得添加相连关系
			vList := strings.Split(inputString, delimiter)

			for i:=0; i<len(vList); i++ {
				for j:=i; j<len(vList); j++ {
					sg.graph.AddEdge(
						sg.Index(vList[i]),
						sg.Index(vList[j]),
					)
				}
			}
		}

		if readError == io.EOF { //
			break
		}
	}

	return
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
