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
func TestMSDSort(t *testing.T) {

	// 简单测试
	Convey("功能测试", t, func() {

		fmt.Println("MSD 功能测试：")

		rawA := []string{
			"WOCB2",
			"WOCAO123",
			"NIDAY%^&*",
			"Y776",
			"CACACAAAAAAAAAA",
			"A",
		}

		PrintArray(rawA)

		Timeit(rawA, MSDSort)

		PrintArray(rawA)

	})

}
