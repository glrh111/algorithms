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


func TestDepthFirstSearch(t *testing.T) {

	// 简单测试
	Convey("功能测试", t, func() {

		fmt.Println("DepthFirstSearch 功能测试：")

		var graph *Graph = NewGraphFromFile("exam/tinyG.txt")

		// source = 0
		var search0 SearchInterface = NewDepthFirstSearch(graph, 0) // 7

		// Count
		So(search0.Count(), ShouldEqual, 7)

		// Marked
		for _, v := range []int{0,1,2,3,4,5,6} {
			So(search0.Marked(v), ShouldEqual, true)
		}
		for _, v := range []int{7,8,9,10,11,12} {
			So(search0.Marked(v), ShouldEqual, false)
		}

		// source = 9
		var search9 SearchInterface = NewDepthFirstSearch(graph, 9)

		// Count
		So(search9.Count(), ShouldEqual, 4)

		// Marked
		for _, v := range []int{9,10,11,12} {
			So(search9.Marked(v), ShouldEqual, true)
		}
		for _, v := range []int{0,1,2,3,4,5,6,7,8} {
			So(search9.Marked(v), ShouldEqual, false)
		}

	})

}