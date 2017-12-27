// 测试内容
// 每个文件一个测试内容

package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	//"fmt"
	//"fmt"
	"fmt"
)


func TestGragh(t *testing.T) {

	// 简单测试
	Convey("功能测试", t, func() {

		fmt.Println("Graph 功能测试：")

		var graph *Graph = NewGraphFromFile("exam/tinyG.txt")

		fmt.Println(graph.ToString())

		fmt.Println("自环数量：", graph.NumberOfSelfLoops())

	})

}