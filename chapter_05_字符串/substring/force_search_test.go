package substring

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

/*
    低位优先排序
 */

// w 为字符串宽度

type SearchTestCase struct {
	S string
	Patt string
	Index int
}

var examMapList = []*SearchTestCase {
	&SearchTestCase{"Wocao", "W", 0},
	&SearchTestCase{"Wocao", "c", 2},
	&SearchTestCase{"wocaoni", "f", 7},
	&SearchTestCase{"Wocao", "ca", 2},
}

func TestLSDSort(t *testing.T) {

	// 简单测试
	Convey("ForceSearch 功能测试", t, func() {

		for _, exam := range examMapList {
			So(forceSearch(exam.S, exam.Patt), ShouldEqual, exam.Index)
		}

		for _, exam := range examMapList {
			So(forceSearch2(exam.S, exam.Patt), ShouldEqual, exam.Index)
		}

		for _, exam := range examMapList {
			fmt.Println(exam)
			So(RabinKarpSearch(exam.S, exam.Patt), ShouldEqual, exam.Index)
		}

	})

}
