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


func TestPathsBFS(t *testing.T) {

	// 简单测试
	Convey("功能测试", t, func() {

		fmt.Println("BFSPaths 功能测试：")

		// tinyG , mediumG
		var graph *Graph = NewGraphFromFile("exam/mediumG.txt")

		var paths PathsInterface = NewBFSPaths(graph, 0)

		fmt.Println(paths.PathTo(100))

	})

}