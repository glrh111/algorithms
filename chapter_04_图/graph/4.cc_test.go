// 测试内容
// 每个文件一个测试内容

package graph

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	//"fmt"
	//"fmt"
	"fmt"
)


func TestCC(t *testing.T) {

	// 简单测试
	Convey("功能测试", t, func() {

		fmt.Println("CC连通分量 功能测试：")

		// tinyG , mediumG
		var graph *Graph = NewGraphFromFile("../exam/tinyG.txt")

		var cc CCInterface = NewCC(graph)

		fmt.Println(cc.Count())
		// Count
		So(cc.Count(), ShouldEqual, 3)

		// 0~6 7,8 9~12
		//fmt.Println("0 - 7: ", cc.Id(0), cc.Id(7), cc.Id(9), cc.Id(12))

		So(cc.Connected(0,6), ShouldEqual, true)
		So(cc.Connected(0,7), ShouldEqual, false)
		So(cc.Connected(0,12), ShouldEqual, false)

		// Id
		So(cc.Id(0), ShouldEqual, 0)
		So(cc.Id(7), ShouldEqual, 1)
		So(cc.Id(12), ShouldEqual, 2)

		// 打印出来所有的连通分量
		for i:=0; i<cc.Count(); i++ {
			fmt.Printf("Id [%d]: ", i)
			for ele := range cc.IteratorChan(i) {
				fmt.Printf("%d ", ele)
			}
			fmt.Println()
		}

	})

}