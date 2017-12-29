// 测试内容
// 每个文件一个测试内容

package graph

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)


func TestSymbolGragh(t *testing.T) {

	// 简单测试
	Convey("功能测试", t, func() {

		fmt.Println("SymbolGraph 功能测试：")

		//var sg *SymbolGraph = NewSymbolGraphFromFile("../exam/movies.txt", "/")
		var sg *SymbolGraph = NewSymbolGraphFromFile("../exam/routes.txt", " ")

		fmt.Println(sg.graph.V())

		// Contains "Martinez, Patrice (I)" not "Wocaonidaye"
		//So(sg.Contains("Martinez, Patrice (I)"), ShouldEqual, true)
		//So(sg.Contains("Wocaonidaye"), ShouldEqual, false)

		for key, value := range sg.indexToName {
			fmt.Println(key, value, len(value))
		}

		fmt.Println(sg.graph.ToString())

	})

}