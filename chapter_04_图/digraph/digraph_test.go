// 测试内容
// 每个文件一个测试内容

package digraph

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)


func TestGragh(t *testing.T) {

	// 简单测试
	Convey("功能测试", t, func() {

		fmt.Println("Digraph 功能测试：")

		var dg *Digraph = NewDigraphFromFile("../exam/tinyDG.txt")

		fmt.Println(dg.ToString())

	})

}