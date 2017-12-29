package string_sort

import (
	"testing"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
)

/*
    低位优先排序
 */

// w 为字符串宽度
func TestLSDSort(t *testing.T) {

	// 简单测试
	Convey("功能测试", t, func() {

		fmt.Println("Paths 功能测试：")

		rawA := []string{
			"WOCAO",
			"NIDAY",
			"CACAC",
			"NIMAD",
			"WOCB2",
		}

		Timeit(rawA, LSDSort)

		// 生成 一系列日期
		Timeit(GenerateDateList(1000000), LSDSort)

	})

}
